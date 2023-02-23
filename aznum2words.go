package aznum2words

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (

	//single digits numbers
	ZeroAsString  string = "sıfır"  // 0
	OneAsString   string = "bir"    // 1
	TwoAsString   string = "iki"    // 2
	ThreeAsString string = "üç"     // 3
	FourAsString  string = "dörd"   // 4
	FiveAsString  string = "beş"    // 5
	SixAsString   string = "altı"   // 6
	SevenAsString string = "yeddi"  // 7
	EightAsString string = "səkkiz" // 8
	NineAsString  string = "doqquz" // 9

	//two digits numbers
	TenAsString     string = "on"     // 10
	TwentyAsString  string = "iyirmi" // 20
	ThirtyAsString  string = "otuz"   // 30
	FortyAsString   string = "qırx"   // 40
	FiftyAsString   string = "əlli"   // 50
	SixtyAsString   string = "altmış" // 60
	SeventyAsString string = "yetmiş" // 70
	EightyAsString  string = "səksən" // 80
	NinetyAsString  string = "doxsan" // 90

	//three digits numbers
	HundredAsString string = "yüz" // 10^2

	//others
	ThousandAsString        string = "min"            // 10^3
	MillionAsString         string = "milyon"         // 10^6
	BillionAsString         string = "milyard"        // 10^9
	TrillionAsString        string = "trilyon"        // 10^12
	QuadrillionAsString     string = "katrilyon"      // 10^15
	QuintillionAsString     string = "kentilyon"      // 10^18
	SextillionAsString      string = "sekstilyon"     // 10^21
	SeptillionAsString      string = "septilyon"      // 10^24
	OctillionAsString       string = "oktilyon"       // 10^27
	NonillionAsString       string = "nonilyon"       // 10^30
	DecillionAsString       string = "desilyon"       //10^33
	UndecillionAsString     string = "undesilyon"     //10^36
	DodecillionAsString     string = "dodesilyon"     //10^39
	TredecillionAsString    string = "tredesilyon"    //10^42
	CathodecillionAsString  string = "katordesilyon"  //10^45
	KendecillionAsString    string = "kendesilyon"    //10^48
	SexdecillionAsString    string = "seksdesilyon"   //10^51
	SeptendecillionAsString string = "septendesilyon" //10^54
	OctodecillionAsString   string = "oktodesilyon"   //10^57
	NovedesillionAsString   string = "novemdesilyon"  //10^60
	VigintillionAsString    string = "vigintilyon"    //10^63
)

const (
	NegativeAsWord  string = "mənfi" // -
	SeparatorAsWord string = "tam"
)

var validNumberRegex = regexp.MustCompile("^-?\\d+(\\.\\d+)?$")

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

	1_000_000_000_000_000:   "bir katrilyonda",
	10_000_000_000_000_000:  "on katrilyonda",
	100_000_000_000_000_000: "yüz katrilyonda",
}

// holds an array of single digits as word representation
var digits = []string{
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
var tens = []string{
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
	if ok := validNumberRegex.MatchString(numberAsStr); !ok {
		return "", errors.New(fmt.Sprintf("Input: %s invalid", numberAsStr))
	}

	isFloatingNumber := strings.Contains(numberAsStr, ".")
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

	signKeyword, err := getSignKeywordAsWord(intValueAsStr)
	if err != nil {
		return "", err
	}

	if signKeyword != "" {
		wordBuilder = append(wordBuilder, signKeyword)
	}

	spelledInteger := convertIntPart(intValueAsStr)

	wordBuilder = append(wordBuilder, spelledInteger)

	return strings.Join(wordBuilder, " "), nil
}

// The function intended to handle the conversion of floating point numbers into words
// handleFloatingPointNumberConversion("-12.1") -> "mənfi on iki tam onda bir"
// handleFloatingPointNumberConversion("0.248551") -> "sıfır tam bir milyonda iki yüz qırx səkkiz min beş yüz əlli bir"
func handleFloatingPointNumberConversion(floatValueAsStr string) (string, error) {
	signKeyword, err := getSignKeywordAsWord(floatValueAsStr)
	if err != nil {
		return "", err
	}

	slices := strings.Split(floatValueAsStr, ".")

	intPartAsWord := convertIntPart(slices[0])

	floatingPart := slices[1]

	//sanitize floatingPart by removing trailing zeros if exist
	floatingPart = strings.TrimRight(floatingPart, "0")

	floatingPartAsIntegerWithWord := convertIntPart(floatingPart)

	cnt := len(floatingPart)
	separatorKey := int(math.Pow10(cnt))

	suffix, ok := floatingPointDict[separatorKey]

	if !ok {
		return "", errors.New(fmt.Sprintf("No separator found"))
	}

	wordBuilder := make([]string, 0)
	if len(signKeyword) != 0 {
		wordBuilder = append(wordBuilder, signKeyword)
	}
	wordBuilder = append(wordBuilder, intPartAsWord)
	wordBuilder = append(wordBuilder, SeparatorAsWord)
	wordBuilder = append(wordBuilder, suffix)
	wordBuilder = append(wordBuilder, floatingPartAsIntegerWithWord)

	return strings.Join(wordBuilder, " "), nil
}

// Extract the sign mark('-') and convert into word
// getSignKeywordAsWord("-123") -> "mənfi"
// getSignKeywordAsWord("123") -> ""
func getSignKeywordAsWord(str string) (string, error) {
	if str == "" {
		return "", errors.New("argument cannot be empty")
	}

	var sign string
	if len(str) > 1 && str[0:1] == "-" {
		sign = NegativeAsWord
	}

	return sign, nil
}

// Remove the sign mark('-') from str
// removeSignMarkIfExists("-123") -> "123"
// removeSignMarkIfExists("123") -> "123"
func removeSignMarkIfExists(str string) string {
	if len(str) > 1 && str[0:1] == "-" {
		return str[1:]
	}

	return str
}

// The function intended to convert any sized positive integer numbers into words
func convertIntPart(strArg string) string {
	str := removeSignMarkIfExists(strArg)

	//used to indicate level 10^3, 10^6
	pos := 1
	wordBuilder := make([]string, 0)

	//starting from right hand side, pick up triples and start process it
	for i := len(str); i >= 0; i = i - 3 {
		key, _ := getKeyword(pos)
		pos++

		var leftPointer int = i - 3
		if leftPointer < 0 {
			leftPointer = 0
		}
		var rightPointer int = i

		tripleAsWords, _ := tripleToWord(str[leftPointer:rightPointer])

		//insert triple `tripleAsWords` in the front of result
		if len(tripleAsWords) != 0 {
			if len(key) != 0 {
				wordBuilder = append([]string{tripleAsWords, key}, wordBuilder...)
			} else {
				wordBuilder = append([]string{tripleAsWords}, wordBuilder...)
			}
		}

	}

	return strings.Join(wordBuilder, " ")
}

// The function designed to get level-word of triples based on its position in representation when scanning from left to right
// e.g "1234", divide it into triples: "1" and "234"
// getKeyword(1) -> "" : keyword for first triple when scanning from right to left
// getKeyword(2) -> "min" : keyword for second triple when scanning from right to left
func getKeyword(position int) (keyword string, err error) {
	switch position {
	case 1:
		//first triple: 10^3
		return "", nil
	case 2:
		return ThousandAsString, nil
	case 3:
		return MillionAsString, nil
	case 4:
		return BillionAsString, nil
	case 5:
		return TrillionAsString, nil
	case 6:
		return QuadrillionAsString, nil
	case 7:
		return QuintillionAsString, nil
	case 8:
		return SextillionAsString, nil
	case 9:
		return SeptillionAsString, nil
	case 10:
		return OctillionAsString, nil
	case 11:
		return NonillionAsString, nil
	case 12:
		return DecillionAsString, nil
	case 13:
		return UndecillionAsString, nil
	case 14:
		return DodecillionAsString, nil
	case 15:
		return TredecillionAsString, nil
	case 16:
		return CathodecillionAsString, nil
	case 17:
		return KendecillionAsString, nil
	case 18:
		return SexdecillionAsString, nil
	case 19:
		return SeptendecillionAsString, nil
	case 20:
		return OctodecillionAsString, nil
	case 21:
		return NovedesillionAsString, nil
	case 22:
		return VigintillionAsString, nil
	default:
		return "", errors.New(fmt.Sprintf("Max supported number level: %s", VigintillionAsString))
	}
}

// The function intended to convert triples("1", "12", "123") into words' representation
// tripleToWord("") -> ""
// tripleToWord("1") -> "bir"
// tripleToWord("12") -> "on iki"
// tripleToWord("123") -> "bir yüz iyirmi üz"
func tripleToWord(triple string) (string, error) {
	if len(triple) == 0 {
		return "", nil
	}

	if len(triple) == 1 {
		return convertOneDigitIntoWord(triple)
	}

	if len(triple) == 2 {
		return convertTwoDigitsIntoWord(triple)
	}

	return convertThreeDigitsIntoWord(triple)
}

// convertOneDigitIntoWord("0") -> "sıfır"
// convertOneDigitIntoWord("1") -> "bir"
func convertOneDigitIntoWord(oneDigitWord string) (string, error) {
	if i, err := strconv.Atoi(oneDigitWord); err != nil {
		return "", err
	} else {
		return digits[i], nil
	}
}

// convertTwoDigitsIntoWord("10") -> "on"
// convertTwoDigitsIntoWord("25") -> "iyirmi beş"
func convertTwoDigitsIntoWord(twoDigitsWord string) (string, error) {
	var textBuilder []string
	var idxAt1 int
	if i, err := strconv.Atoi(twoDigitsWord[1:]); err != nil {
		return "", err
	} else {
		idxAt1 = i
	}
	var wordForIdxAt1 string
	if idxAt1 == 0 {
		wordForIdxAt1 = ""
	} else {
		wordForIdxAt1 = digits[idxAt1]
	}

	if len(wordForIdxAt1) != 0 {
		textBuilder = append([]string{wordForIdxAt1}, textBuilder...)
	}

	var idxAt0 int
	if i, err := strconv.Atoi(twoDigitsWord[0:1]); err != nil {
		return "", err
	} else {
		idxAt0 = i
	}

	wordForIdxAt0 := tens[idxAt0]

	if len(wordForIdxAt0) != 0 {
		textBuilder = append([]string{wordForIdxAt0}, textBuilder...)
	}

	return strings.Join(textBuilder, " "), nil
}

// convertThreeDigitsIntoWord("390") -> "üç yüz doxsan"
// convertThreeDigitsIntoWord("125") -> "bir yüz iyirmi beş"
func convertThreeDigitsIntoWord(threeDigitsWord string) (string, error) {
	var textBuilder []string
	for i := len(threeDigitsWord) - 1; i >= 0; i-- {

		var word string
		switch i {
		case 2:
			var ValAtIdx2 int
			if i, err := strconv.Atoi(threeDigitsWord[2:]); err != nil {
				return "", err
			} else {
				ValAtIdx2 = i
			}
			if ValAtIdx2 != 0 {
				word = digits[ValAtIdx2]
			}

		case 1:
			var ValAtIdx1 int
			if i, err := strconv.Atoi(threeDigitsWord[1:2]); err != nil {
				return "", err
			} else {
				ValAtIdx1 = i
			}
			if ValAtIdx1 != 0 {
				word = tens[ValAtIdx1]
			} else {
				word = ""
			}
		case 0:
			var ValAtIdx0 int
			if i, err := strconv.Atoi(threeDigitsWord[0:1]); err != nil {
				return "", err
			} else {
				ValAtIdx0 = i
			}
			isNonZero := ValAtIdx0 != 0

			var prefix string
			if isNonZero {
				prefix = digits[ValAtIdx0]
			} else {
				prefix = ""
			}

			if len(prefix) > 0 {
				word = prefix + " " + HundredAsString
			} else {
				word = ""
			}
		default:
			return "", errors.New("")
		}

		if len(word) != 0 {
			textBuilder = append([]string{word}, textBuilder...)
		}
	}

	return strings.Join(textBuilder, " "), nil
}
