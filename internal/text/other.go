package text

// Название переменных окружения
const (
	EnvBotToken       = "EXRUBBOT_TOKEN"
	EnvDatabaseDSN    = "EXRUBBOT_DATABASE_DSN"
	EnvSupportGroupID = "EXRUBBOT_SUPPORT_GROUP_ID"
	EnvDebugChannelID = "EXRUBBOT_DEBUG_CHANNEL_ID"
)

const (
	InfoEnvFileNotFound      = ".env файл не найден"
	InfoUpdateProcessingTime = "update #%d обработан за %s"
	InfoUpdateProcessingSkip = "пропуск обработки update.%s"
	InfoUpdateUnknownType    = "(не определён)"
)
