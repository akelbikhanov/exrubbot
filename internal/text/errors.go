package text

// Сообщения об ошибках.
const (
	ErrEnvMissingVar = "не найдена переменная окружения"
	ErrEnvFileLoad   = "ошибка загрузки .env файла"

	ErrBotAlreadyRunning = "бот уже запущен"
	ErrRequestUpdates    = "ошибка получения getUpdates"

	ErrRequestBuild     = "не удалось сформировать HTTP-запрос"
	ErrRequestSend      = "ошибка выполнения HTTP-запроса"
	ErrRequestTimeout   = "истёк таймаут запроса"
	ErrRequestCancelled = "запрос отменён контекстом"
	ErrRequestClose     = "ошибка закрытия тела ответа"
	ErrRequestAPI       = "ошибка получения данных, код"
	ErrRequestNoQuotes  = "биржа вернула пустой ответ"

	ErrDataFeedName = "неизвестный поставщик данных"

	ErrDecodeJSON = "ошибка декодирования JSON"
)
