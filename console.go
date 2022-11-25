package flog

const (
	colorReset     = "\033[0m"
	colorBold      = "\033[1m"
	colorBlack     = "\033[30m"
	colorRed       = "\033[31m"
	colorGreen     = "\033[32m"
	colorYellow    = "\033[33m"
	colorBlue      = "\033[34m"
	colorMagenta   = "\033[35m"
	colorCyan      = "\033[36m"
	colorWhite     = "\033[37m"
	colorBlackHigh = "\033[90m"
)
const (
	colorFatal = "\033[5;41;39m"
	colorError = "\033[1;31m"
	colorWarn  = "\033[1;33m"
	colorInfo  = "\033[1;32m"
	colorDebug = "\033[2;37m"
)

var levelConsoleMap = []string{
	colorFatal + "FATAL" + colorReset,
	colorError + "ERR" + colorReset,
	colorWarn + "WRN" + colorReset,
	colorInfo + "INF" + colorReset,
	colorDebug + "DBG" + colorReset,
}

const SEG_LEN_LEVEL = 7 + 3 + 4
