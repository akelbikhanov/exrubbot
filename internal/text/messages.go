package text

// Основные сообщения бота
const (
	MsgWelcome = "Вас приветствует ExRubBot!\n\n" +
		"Получаю курс USDT/RUB с бирж и агрегаторов.\n\n" +
		"Для запроса котировки: " + CmdQuote

	MsgSelectProvider = "Выберите поставщика:"
	MsgSelectInterval = "Выберите интервал:"
	MsgProviderError  = "Список поставщиков недоступен"
	MsgUnknownCommand = "Неизвестная команда"
	MsgCallbackError  = "Некорректные данные: %s"

	MsgUnsubscribeOK   = "Подписка отменена.\nВозобновить: " + CmdQuote
	MsgUnsubscribeNone = "Активных подписок нет.\nСоздать: " + CmdQuote
)

// Формат котировки
const (
	QuoteHeader     = "USDT/RUB • "
	QuoteError      = "\nОшибка получения котировки"
	QuoteStopHint   = "\n" + CmdStop + " ⏱%s"
	QuoteTimeFormat = "2006.01.02 15:04:05 MST (UTC -07)"
	QuoteTimeZone   = "MSK"
)
