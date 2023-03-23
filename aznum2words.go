package aznum2words

import (
	"math"
	"strings"
)

// Floating point number indicators
// count of zeros : suffix
var floatingPointDict = map[int]string{
	10:  "onda",
	100: "yüzdə",

	1000:    "bir mində",
	10_000:  "on mində",
	100_000: "yüz mində",

	1_000_000:   "bir milyonda",
	10_000_000:  "on milyonda",
	100_000_000: "yüz milyonda",

	1_000_000_000:   "bir milyardda",
	10_000_000_000:  "on milyardda",
	100_000_000_000: "yüz milyardda",

	1_000_000_000_000:   "bir trilyonda",
	10_000_000_000_000:  "on trilyonda",
	100_000_000_000_000: "yüz trilyonda",

	1_000_000_000_000_000:   "bir kvadrilyonda",
	10_000_000_000_000_000:  "on kvadrilyonda",
	100_000_000_000_000_000: "yüz kvadrilyonda",
}

// holds an array of single digits as word representation
var digits = [10]string{
	ZeroAsString,
	OneAsString,
	TwoAsString,
	ThreeAsString,
	FourAsString,
	FiveAsString,
	SixAsString,
	SevenAsString,
	EightAsString,
	NineAsString,
}

// holds the array of tens as word representation
var tens = [10]string{
	"",
	TenAsString,
	TwentyAsString,
	ThirtyAsString,
	FortyAsString,
	FiftyAsString,
	SixtyAsString,
	SeventyAsString,
	EightyAsString,
	NinetyAsString,
}

// SpellNumber takes a string input numberAsStr and returns a string and an error.
// The purpose of the function is to convert a given number string into its equivalent Azerbaijani words.
// The function starts by validating the input string using a regular expression. If the input string is not in a valid format, it returns an error.
// If the input string is valid, it then checks whether it is a floating-point number or an integer number.
// Depending on this, it calls either `handleFloatingPointNumberConversion` or `handleIntegerNumberConversion` to convert the number into its equivalent Azerbaijani words.
// If the conversion is successful, the function returns the result as a string. If an error occurs during the conversion, the function returns an error.
func SpellNumber(numberAsStr string) (string, error) {
	//check whether number is in valid format
	if err := validateInput(numberAsStr); err != nil {
		return "", err
	}

	if err := checkConstraints(numberAsStr); err != nil {
		return "", err
	}

	isFloatingNumber := strings.Contains(numberAsStr, DecimalPointSeparator)
	if isFloatingNumber {
		return handleFloatingPointNumberConversion(numberAsStr)
	}

	return handleIntegerNumberConversion(numberAsStr)
}

// The function intended to handle the conversion of integer numbers into words
// handleIntegerNumberConversion("-79594") -> "mənfi yetmiş doqquz min beş yüz doxsan dörd"
// handleIntegerNumberConversion("81") -> "səksən bir"
func handleIntegerNumberConversion(intValueAsStr string) (string, error) {
	wordBuilder := make([]string, 0)

	signKeyword, err := getSignSymbolAsWord(intValueAsStr)
	if err != nil {
		return "", err
	}

	if signKeyword != "" {
		wordBuilder = append(wordBuilder, signKeyword)
	}

	intValueWithoutSign := removeSignMarkIfExists(intValueAsStr)
	spelledInteger := convertIntPart(intValueWithoutSign)

	// handling case: intValueAsStr is '-0'
	if intValueWithoutSign == "0" {
		return spelledInteger, nil
	}

	wordBuilder = append(wordBuilder, spelledInteger)

	return strings.Join(wordBuilder, " "), nil
}

// The function intended to handle the conversion of floating point numbers into words
// handleFloatingPointNumberConversion("-12.1") -> "mənfi on iki tam onda bir"
// handleFloatingPointNumberConversion("0.248551") -> "sıfır tam bir milyonda iki yüz qırx səkkiz min beş yüz əlli bir"
func handleFloatingPointNumberConversion(floatValueAsStr string) (string, error) {
	signKeyword, err := getSignSymbolAsWord(floatValueAsStr)
	if err != nil {
		return "", err
	}

	floatValueAsStr = removeSignMarkIfExists(floatValueAsStr)
	slices := strings.Split(floatValueAsStr, DecimalPointSeparator)

	intPartAsWord := convertIntPart(slices[0])

	floatingPart := slices[1]

	//sanitize floatingPart by removing trailing zeros if exist
	floatingPart = strings.TrimRight(floatingPart, "0")

	// handling the special case: floatValueAsStr is "0.0"
	if slices[0] == "0" && len(floatingPart) == 0 {
		return intPartAsWord, nil
	}

	var floatingPartAsIntegerWithWord string
	var suffix string
	if len(floatingPart) != 0 {
		floatingPartAsIntegerWithWord = convertIntPart(floatingPart)
		cnt := len(floatingPart)
		separatorKey := int(math.Pow10(cnt))
		resultSuffix, ok := floatingPointDict[separatorKey]

		if !ok {
			return "", ErrTooBigScale
		}
		suffix = resultSuffix
	}

	wordBuilder := make([]string, 0)
	if len(signKeyword) != 0 {
		wordBuilder = append(wordBuilder, signKeyword)
	}
	wordBuilder = append(wordBuilder, intPartAsWord)

	// handling for cases: floatValueAsStr is XXX.00
	if len(suffix) != 0 {
		wordBuilder = append(wordBuilder, SeparatorAsWord)
		wordBuilder = append(wordBuilder, suffix)
		wordBuilder = append(wordBuilder, floatingPartAsIntegerWithWord)
	}

	return strings.Join(wordBuilder, " "), nil
}
