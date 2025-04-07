package common

import (
	"time"
)

// Название переменных окружения
const (
	EnvBotToken       = "RUBALERTBOT_TOKEN"
	EnvDatabaseDSN    = "RUBALERTBOT_DATABASE_DSN"
	EnvSupportGroupID = "RUBALERTBOT_SUPPORT_GROUP_ID"
	EnvDebugChannelID = "RUBALERTBOT_DEBUG_CHANNEL_ID"
	DefaultTimeout    = 5 * time.Second
)

// Внутренние ошибки и сообщения, которые логируются
const (
	ErrBotAlreadyRunning = "бот уже запущен"
	ErrLoadEnv           = "ошибка загрузки переменных окружения из .env файла: %v"
	ErrMissingEnvVar     = "отсутствует переменная окружения: %s"
	InfoUpdateSkip       = "update #%d был пропущен"
	InfoUpdateProcessed  = "update #%d обработан за %s"
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
	MessageStart = "Вас приветствует сервис RubAlert!\n\nЯ умею:\n" +
		"• получать актуальную цену USDT/RUB с биржи ABCEX (/price)\n" +
		"• периодически присылать её вам (/repeat)\n\n" +
		"Для подробностей используйте /help."
	MessageHelp = "Доступные команды:\n" +
		CommandStart + " – запуск/описание бота\n" +
		CommandHelp + " – список доступных команд\n" +
		CommandPrice + " – получить текущую цену USDT/RUB с ABCEX\n" +
		CommandRepeat + " – периодическое получение цены\n" +
		CommandStop + " – остановить рассылку"
	MessagePrice   = "USDT/RUB • ABCEX\nask: %.2f RUB (%0.f USDT)\nbid: %.2f RUB (%0.f USDT)\n\n%s MSK (UTC+3)"
	MessageRepeat  = "Выберите интервал получения цен:"
	MessageStopNo  = "У вас нет активной подписки." + MessageStopAdd
	MessageStopYes = "Периодическое получение остановлено." + MessageStopAdd
	MessageStopAdd = "\nВозобновить: " + CommandRepeat + ", разово: " + CommandPrice
	MessageError   = "Команда не распознана!"
)

// Сообщения Garantex
const (
	ErrABCEX         = "ошибка при получении цен: "
	ErrABCEXRequest  = "ошибка запроса к API ABCEX: %v"
	ErrABCEXAPI      = "ошибка API ABCEX: код %d"
	ErrABCEXJSON     = "ошибка декодирования ответа: %v"
	ErrABCEXNoQuotes = "недостаточно данных в ответе API"
)
