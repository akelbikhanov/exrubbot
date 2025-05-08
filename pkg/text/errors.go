// Package text содержит строковые константы.
package text

// Ошибки запроса поставщиков данных.
const (
	ErrRequestBuild     = "не удалось сформировать HTTP-запрос"
	ErrRequestSend      = "ошибка выполнения HTTP-запроса"
	ErrRequestTimeout   = "истёк таймаут запроса"
	ErrRequestCancelled = "запрос отменён контекстом"
	ErrRequestStatus    = "сервер вернул код ошибки:"
	ErrRequestBody      = "сервер вернул пустое тело ответа"
	ErrDecodeJSON       = "ошибка декодирования JSON"
	ErrRequestClose     = "ошибка закрытия тела ответа"
	ErrRequestNoQuotes  = "сервер вернул пустой стакан цен (нет котировок)"
	ErrFeedNotSupport   = "данный поставщик пока не поддерживается"
)
