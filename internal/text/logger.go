package text

// logger.Error
const (
	ErrEnvMissingVar     = "не найдена переменная окружения"
	ErrEnvFileLoad       = "ошибка загрузки .env файла"
	ErrBotAlreadyRunning = "бот уже запущен"
	ErrBotTerminate      = "бот остановлен: не определена корректная реакция на ошибку"
)

// logger.Info
const (
	InfoEnvFileNotFound      = ".env файл не найден"
	InfoUpdateProcessingTime = "update #%d обработан за %s"
	InfoUpdateProcessingSkip = "пропуск обработки update.%s"
	InfoUpdateUnknownType    = "(не определён)"
	InfoChatMemberBanned     = "пользователь %d заблокировал бота, рассылка для него прекращена"
)

// logger.Warn
const (
	WarnRequestDelay         = "ошибка 429: превышен лимит запросов, введена техническая пауза на %s"
	WarnForbiddenSubscribe   = "ошибка 403: нарушение конфиденциальности, рассылка пользователю прекращена"
	WarnForbiddenUnsubscribe = "ошибка 403: нарушение конфиденциальности, пользователь не был подписан на рассылку"
)
