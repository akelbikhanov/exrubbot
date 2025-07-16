package text

// Флаги командной строки
const (
	FlagVersion      = "version"
	FlagVersionDesc  = "show version and exit"
	FlagLogLevel     = "log-level"
	FlagLogLevelDesc = "set the logging level:\n" +
		"-4 debug, 0 info (default), 4 warn, 8 error"
)

const (
	DescLogLevel1 = "set the logging level, for example:\n" +
		"-log-level -4	debug\n" +
		"-log-level 0	info (default)\n" +
		"-log-level 4	warn\n" +
		"-log-level 8	error"
)

// Переменные окружения
const (
	EnvBotToken    = "EXRUBBOT_TOKEN"
	EnvStoragePath = "EXRUBBOT_SUBSCRIBERS_JSON_PATH"
)

// Форматы версий и файлов
const (
	StorageVersion     = "1.0"
	StorageBackupFmt   = "%s.backup-%d"
	StorageTempPattern = "subscriptions-*.tmp"
)
