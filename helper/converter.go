package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (

	//single digits numbers
	ZeroAsString  string = "sıfır"
	OneAsString   string = "bir"
	TwoAsString   string = "iki"
	ThreeAsString string = "üç"
	FourAsString  string = "dörd"
	FiveAsString  string = "beş"
	SixAsString   string = "altı"
	SevenAsString string = "yeddi"
	EightAsString string = "səkkiz"
	NineAsString  string = "doqquz"

	//two digits numbers
	TenAsString     string = "on"
	TwentyAsString  string = "iyirmi"
	ThirtyAsString  string = "otuz"
	FortyAsString   string = "qırx"
	FiftyAsString   string = "əlli"
	SixtyAsString   string = "altmış"
	SeventyAsString string = "yetmiş"
	EightyAsString  string = "səksən"
	NinetyAsString  string = "doxsan"

	//three digits numbers
	HundredAsString  string = "yüz"
	ThousandAsString string = "min"
	MillionAsString  string = "milyon"
	BillionAsString  string = "milyard"
)

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

// Convert integer given in str to word
func ConvertIntPart(str string) string {

	//it contains separator like `,`
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
	default:
		return "", errors.New(fmt.Sprintf("Position: %d not supported", position))
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
