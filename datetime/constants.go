package datetime

const (
	// Months in Azerbaijani
	January   = "yanvar"
	February  = "fevral"
	March     = "mart"
	April     = "aprel"
	May       = "may"
	June      = "iyun"
	July      = "iyul"
	August    = "avqust"
	September = "sentyabr"
	October   = "oktyabr"
	November  = "noyabr"
	December  = "dekabr"

	// Time-related words
	Hour     = "saat"
	Minute   = "dəqiqə"
	Second   = "saniyə"
	AM       = "səhər"
	PM       = "axşam"
	Noon     = "günorta"
	Midnight = "gecəyarı"

	// Date-related words
	Year  = "il"
	Month = "ay"
	Day   = "gün"
	Of    = "i"
	The   = "nın"
)

// Months array for easy access
var Months = [12]string{
	January, February, March, April, May, June,
	July, August, September, October, November, December,
}
