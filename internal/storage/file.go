package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/entity"
	"github.com/akelbikhanov/exrubbot/internal/text"
)

// FileStorage файловое хранилище.
type FileStorage struct {
	filePath string
}

// fileSchema схема данных в файле (формат JSON)
type fileSchema struct {
	Version string                `json:"version"`
	Updated time.Time             `json:"updated"`
	Items   []entity.Subscription `json:"items"`
}

// NewFileStorage создаёт новое хранилище.
func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{filePath: filePath}
}

// Load загружает подписки из файла.
// Если файл не существует, возвращает пустой список.
func (f *FileStorage) Load() ([]entity.Subscription, error) {
	// Передан пустой путь, работаем только в оперативной памяти.
	if f.filePath == "" {
		return []entity.Subscription{}, nil
	}

	// Указанный файл не найден. Это нормально при первом запуске.
	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		return []entity.Subscription{}, nil
	}

	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, fmt.Errorf(text.ErrStorageRead, err)
	}

	var file fileSchema
	if err = json.Unmarshal(data, &file); err != nil {
		// Пытаемся сохранить повреждённый файл
		backup := fmt.Sprintf(text.StorageBackupFmt, f.filePath, time.Now().Unix())
		_ = os.Rename(f.filePath, backup)
		return nil, fmt.Errorf(text.ErrStorageParse, err)
	}

	return file.Items, nil
}

// Save сохраняет подписки в файл атомарно.
func (f *FileStorage) Save(items []entity.Subscription) error {
	// Передан пустой путь, работаем только в оперативной памяти.
	if f.filePath == "" {
		return nil
	}

	file := fileSchema{
		Version: text.StorageVersion,
		Updated: time.Now(),
		Items:   items,
	}

	data, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		return fmt.Errorf(text.ErrStorageMarshal, err)
	}

	// Создаём временный файл в той же директории
	dir := filepath.Dir(f.filePath)
	tmp, err1 := os.CreateTemp(dir, text.StorageTempPattern)
	if err1 != nil {
		return fmt.Errorf(text.ErrStorageTemp, err1)
	}
	tmpPath := tmp.Name()

	// Удаляем временный файл в случае ошибки
	defer func() {
		if err != nil {
			_ = os.Remove(tmpPath)
		}
	}()

	// Записываем и синхронизируем
	if _, err = tmp.Write(data); err != nil {
		_ = tmp.Close()
		return fmt.Errorf(text.ErrStorageWrite, err)
	}

	if err = tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf(text.ErrStorageSync, err)
	}

	if err = tmp.Close(); err != nil {
		return fmt.Errorf(text.ErrStorageClose, err)
	}

	// Атомарно перемещаем
	if err = os.Rename(tmpPath, f.filePath); err != nil {
		return fmt.Errorf(text.ErrStorageRename, err)
	}

	return nil
}
