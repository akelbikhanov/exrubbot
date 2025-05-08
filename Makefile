# Makefile для проекта ExRubBot
# Упрощает запуск линтера, тестов и сборку проекта.

.PHONY: all lint test build run clean

# Бинарный файл
BINARY_NAME=exrubbot
# Путь к точке входа
ENTRYPOINT=cmd/bot/main.go
# Папка для бинарников
BIN_DIR=bin

# Цель по умолчанию
all: lint test build

# Линтер
lint:
	@golangci-lint run

# Тестирование
test:
	@go test ./... -v -count=1

# Сборка бинарника
build:
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(ENTRYPOINT)

# Запуск собранного бинарника
run: build
	@./$(BIN_DIR)/$(BINARY_NAME)

# Очистка артефактов сборки
clean:
	@if exist $(BIN_DIR) ( \
		echo Cleaning $(BIN_DIR)... && \
		rmdir /S /Q $(BIN_DIR) \
	) else ( \
		echo Nothing to clean. \
	)
