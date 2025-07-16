package text

// Команды бота
const (
	CmdStart   = "/start"
	CmdQuote   = "/quote"
	CmdStop    = "/stop"
	CmdVersion = "/version"
)

// Callback префиксы и разделители
const (
	CallbackPrefixQuote = "quote"
	CallbackSeparator   = ":"
)

// Кнопки интерфейса
const (
	BtnQuoteOnce   = "Разово"
	BtnInterval1M  = "1м"
	BtnInterval5M  = "5м"
	BtnInterval30M = "30м"
	BtnInterval1H  = "1ч"
	BtnInterval6H  = "6ч"
	BtnInterval1D  = "1д"
	BtnBack        = "◀ Назад"
)

// Форматы callback-данных
const (
	CallbackFormatFeed     = "%s%s%s"     // prefix:feedID
	CallbackFormatInterval = "%s:%s:%d"   // prefix:feedID:seconds
	CallbackFormatFull     = "%s%s%s%s%d" // prefix:sep:feedID:sep:seconds
)

// UI символы
const (
	SymbolAskPrice = "🔼"
	SymbolBidPrice = "🔽"
)

// Суффиксы времени
const (
	SuffixDays    = "д"
	SuffixHours   = "ч"
	SuffixMinutes = "м"
	SuffixSeconds = "с"
	SuffixZero    = "0с"
)
