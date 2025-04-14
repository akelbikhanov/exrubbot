// Package text содержит строковые константы,
// используемые в различных частях приложения.
package text

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
	MessageStart = "Вас приветствует сервис ExRubBot!\n\nЯ умею:\n" +
		"• получать актуальную котировку USDT/RUB с биржи Grinex (/price)\n" +
		"• периодически присылать её вам (/repeat)\n\n" +
		"Для подробностей используйте /help."
	MessageHelp = "Доступные команды:\n" +
		CommandStart + " – запуск/описание бота\n" +
		CommandHelp + " – список доступных команд\n" +
		CommandPrice + " – получить текущую котировку USDT/RUB с Grinex\n" +
		CommandRepeat + " – периодическое получение котировки\n" +
		CommandStop + " – остановить рассылку"
	MessageRepeat         = "Выберите интервал получения котировки:"
	MessageStopNo         = "У вас нет активной подписки." + MessageStopAdd
	MessageStopYes        = "Периодическое получение остановлено." + MessageStopAdd
	MessageStopAdd        = "\nВозобновить: " + CommandRepeat + ", разово: " + CommandPrice
	MessageUnknownCommand = "Команда не распознана!"

	MessageQuoteHead   = "USDT/RUB • "
	MessageQuoteFailed = "Не удалось получить котировку!"
	MessageQuoteFormat = "🔼 %s (%s)\n🔽 %s (%s)"
	MessageQuoteBottom = "\n" + CommandStop + " ⏱"
)
