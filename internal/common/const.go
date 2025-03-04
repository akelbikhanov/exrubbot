package common

import "time"

// Название переменных окружения
const (
	EnvBotToken       = "GARANTEX_BOT_TOKEN"
	EnvDatabaseDSN    = "GARANTEX_DATABASE_DSN"
	EnvSupportGroupID = "GARANTEX_SUPPORT_GROUP_ID"
	EnvDebugChannelID = "GARANTEX_DEBUG_CHANNEL_ID"
	DefaultTimeout    = 5 * time.Second
)

// Внутренние ошибки и сообщения, которые логируются
const (
	ErrBotAlreadyRunning = "бот уже запущен"
	ErrLoadEnv           = "ошибка загрузки переменных окружения из .env файла: %v"
	ErrMissingEnvVar     = "отсутствует переменная окружения: %s"
	InfoUpdateSkip       = "update #%d был пропущен"
	InfoUpdateProcessed  = "update #%d был обработан за %v"
)

// Команды бота
const (
	CommandStart  = "/start"
	CommandHelp   = "/help"
	CommandPrice  = "/price"
	CommandRepeat = "/repeat"
	CommandStop   = "/stop"
)

// Сообщения бота
const (
	MessageStart = "Вас приветствует сервис Garantex Notifier!\n\nЯ умею:\n" +
		"• получать актуальную цену USDT/RUB с Garantex (/price)\n" +
		"• периодически присылать её вам (/repeat)\n\n" +
		"Для подробностей используйте /help."
	MessageHelp = "Доступные команды:\n" +
		"• /price – получает актуальную цену USDT/RUB с Garantex.\n" +
		"• /repeat – периодическое получение цены.\n" +
		"• /stop – прекращает автоматическую рассылку."
	MessagePrice   = "USDT/RUB • Garantex\nask: %.2f RUB (%0.f USDT)\nbid: %.2f RUB (%0.f USDT)\n\n%s MSK (UTC+3)"
	MessageStopNo  = "У вас нет активной подписки"
	MessageStopYes = "Периодическая отправка остановлена"
	MessageError   = "произошла ошибка: %v"
)

// Сообщения Garantex
const (
	ErrGarantex         = "ошибка при получении цен: "
	ErrGarantexRequest  = "ошибка запроса к API Garantex: %v"
	ErrGarantexAPI      = "ошибка API Garantex: код %d"
	ErrGarantexJSON     = "ошибка декодирования ответа: %v"
	ErrGarantexNoQuotes = "недостаточно данных в ответе API"
)
