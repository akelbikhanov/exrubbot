package text

// Команды бота.
const (
	CommandStart = "/start"
	CommandQuote = "/quote"
	CommandStop  = "/stop"
)

// Элементы (составные части) callback-команд.
const (
	CallbackQuotePrefix = "quote"
	CallbackSeparator   = ":"
	CallbackBack        = "back"
)

// Кнопки выбора интервала.
const (
	TimerButtonOnce = "Разово"
	TimerButton1M   = "1м"
	TimerButton5M   = "5м"
	TimerButton30M  = "30м"
	TimerButton1H   = "1ч"
	TimerButton6H   = "6ч"
	TimerButton1D   = "1д"
	TimerButtonBack = "◀ Назад"
)
