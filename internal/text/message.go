// Package text содержит строковые константы, используемые в различных частях приложения.
package text

// Сообщения бота.
const (
	MessageStart = "Вас приветствует сервис ExRubBot!\n\n" +
		"Я умею получать текущий курс USDT/RUB с разных бирж и агрегаторов.\n\n" +
		"Для запроса котировки воспользуйтесь командой " + CommandQuote

	MessageNoFeed         = "Не удалось получить список поставщиков данных!"
	MessageFeed           = "Выберите поставщика котировки:"
	MessageTimer          = "Выберите интервал получения:"
	MessageStopNo         = "У вас нет активной подписки." + MessageStopAdd
	MessageStopYes        = "Периодическое получение остановлено." + MessageStopAdd
	MessageStopAdd        = "\nВозобновить: " + CommandQuote
	MessageCommandUnknown = "Команда не распознана!"

	MessageCallbackIncorrect = "Ошибка: некорректный callback: '%s'"
	Message
)

// Формат вывода сообщения с котировкой.
const (
	QuoteHead       = "USDT/RUB • "
	QuoteFailed     = "\nНе удалось получить котировку!"
	QuoteFormat     = "\n🔼 %s (%s)\n🔽 %s (%s)"
	QuoteStop       = "\n" + CommandStop + " ⏱%s"
	QuoteTimeLayout = "2006.01.02 15:04:05"
)
