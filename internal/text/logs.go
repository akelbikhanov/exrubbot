// Package text содержит строковые константы, используемые в различных частях приложения.
package text

// Поля структурированного логирования
const (
	LogFieldChatID      = "ChatID"
	LogFieldFeedID      = "FeedID"
	LogFieldError       = "Error"
	LogFieldMessageID   = "MessageID"
	LogFieldCallbackID  = "CallbackID"
	LogFieldMessageText = "Message.Text"
	LogFieldIntervalSec = "IntervalSec"
)

// Debug уровень
const (
	LogDebugUpdateProcessed   = "обновление #%d за %s"
	LogDebugUpdateUnknownType = "update #%d не поддерживаемый тип обновления"
	LogDebugSchedulerStart    = "запуск задачи/планировщика"
	LogDebugSchedulerStop     = "остановка задачи/планировщика"
	LogDebugSendMessage       = "bot.SendMessage()"
	LogDebugEditMessage       = "bot.EditMessage()"
	LogDebugAnswerCallback    = "bot.AnswerCallbackQuery()"
)

// Info уровень
const (
	LogInfoStarting      = "запуск бота"
	LogInfoStopped       = "бот остановлен"
	LogInfoSubsCreated   = "создана новая подписка"
	LogInfoUserBlocked   = "пользователь заблокировал бота"
	LogInfoSubsCancelled = "пользователь отменил подписку"
	LogInfoSubsRestored  = "восстановлено подписок: %d, пропущено: %d"
	LogInfoSubsSaved     = "сохранено подписок: %d"
)

// Warn уровень
const (
	LogWarnRateLimit      = "превышен лимит запросов, пауза %s"
	LogWarnAccessDenied   = "доступ запрещён (403)"
	LogWarnUnknownCommand = "update #%d: неизвестная команда"
)

// Error уровень
const (
	LogErrEnvFile         = "ошибка загрузки .env"
	LogErrEnvMissing      = "отсутствует переменная"
	LogErrConfig          = "ошибка конфигурации"
	LogErrTelegramInit    = "ошибка инициализации telegram"
	LogErrCritical        = "критическая ошибка"
	LogErrQuoteFetch      = "ошибка получения котировки"
	LogErrFeedNotFound    = "поставщик '%s' не найден"
	LogErrInvalidInterval = "некорректный интервал: %d"
	LogErrSubsSave        = "ошибка сохранения подписок"
	LogErrSubsLoad        = "ошибка загрузки подписок"
	LogErrSubsRestore     = "ошибка восстановления подписки"
	LogErrSubsCreate      = "ошибка создания подписки"
	LogErrSchedulerRun    = "перехвачена паника в планировщике"
)

// Ошибки хранилища
const (
	ErrStorageRead    = "чтение файла: %w"
	ErrStorageParse   = "парсинг json: %w"
	ErrStorageMarshal = "формирование json: %w"
	ErrStorageTemp    = "создание временного файла: %w"
	ErrStorageWrite   = "запись файла: %w"
	ErrStorageSync    = "синхронизация файла: %w"
	ErrStorageClose   = "закрытие файла: %w"
	ErrStorageRename  = "переименование файла: %w"
)

// Ошибки конфигурации
const (
	ErrConfigEnvLoad    = "%s: %w"
	ErrConfigTokenEmpty = "%s '%s'"
)
