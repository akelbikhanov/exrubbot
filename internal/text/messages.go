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
	QuoteTimeFormat = "02.01.2006 15:04:05 MSK"
	QuoteTimeZone   = "MSK"
)
