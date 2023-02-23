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

func SpellNumber(str string) (string, error) {
	//check whether number is in valid format
	if ok := validNumberRegex.MatchString(str); !ok {
		return "", errors.New(fmt.Sprintf("Input: %s invalid", str))
	}

	//it may contain separator like `,` placed every 3 decimal places for numbers larger than 999
	str = strings.ReplaceAll(str, ",", "")

	isFloatingNumber := strings.Contains(str, ".")
	if isFloatingNumber {
		return handleFloatingPointNumberConversion(str)
	} else {
		return handleIntegerNumberConversion(str)
	}
}

func handleIntegerNumberConversion(intValueAsStr string) (string, error) {
	wordBuilder := make([]string, 0)

	sign, err := getSignKeywordAsWord(intValueAsStr)
	if err != nil {
		return "", err
	}

	if sign != "" {
		wordBuilder = append(wordBuilder, sign)
	}

	spelledInteger := convertIntPart(intValueAsStr)

	wordBuilder = append(wordBuilder, spelledInteger)

	return strings.Join(wordBuilder, " "), nil
}

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

func removeSignMarkIfExists(str string) string {
	if len(str) > 1 && str[0:1] == "-" {
		return str[1:]
	}

	return str
}

// Convert integer given in str to word
func convertIntPart(strArg string) string {
	str := removeSignMarkIfExists(strArg)

	//starting from right hand side, pick up triples and start process it

	//used to indicate level 10^3, 10^6
	pos := 1
	wordBuilder := make([]string, 0)
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

func tripleToWord(triple string) (string, error) {

	if len(triple) == 0 {
		return "", nil
	}

	if len(triple) == 1 {
		if i, err := strconv.Atoi(triple); err != nil {
			return "", err
		} else {
			return digits[i], nil
		}
	}

	if len(triple) == 2 {
		var textBuilder []string
		var idxAt1 int
		if i, err := strconv.Atoi(triple[1:]); err != nil {
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
		if i, err := strconv.Atoi(triple[0:1]); err != nil {
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

	var textBuilder []string
	for i := len(triple) - 1; i >= 0; i-- {

		var word string
		switch i {
		case 2:
			var ValAtIdx2 int
			if i, err := strconv.Atoi(triple[2:]); err != nil {
				return "", err
			} else {
				ValAtIdx2 = i
			}
			if ValAtIdx2 != 0 {
				word = digits[ValAtIdx2]
			}

		case 1:
			var ValAtIdx1 int
			if i, err := strconv.Atoi(triple[1:2]); err != nil {
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
			if i, err := strconv.Atoi(triple[0:1]); err != nil {
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
