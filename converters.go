package aznum2words

import (
	"strconv"
	"strings"
)

// Extract the sign mark('-') and convert into word
// getSignSymbolAsWord("-123") -> "mənfi"
// getSignSymbolAsWord("123") -> ""
func getSignSymbolAsWord(str string) (string, error) {
	if str == "" {
		return "", ErrInvalidArgument
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
		return "", ErrTooBigNumber
	}
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

// convertOneDigitIntoWord("0") -> "sıfır"
// convertOneDigitIntoWord("1") -> "bir"
func convertOneDigitIntoWord(oneDigitWord string) (string, error) {
	if len(oneDigitWord) != 1 {
		return "", ErrInvalidArgument
	}

	if i, err := strconv.Atoi(oneDigitWord); err != nil {
		return "", err
	} else {
		return digits[i], nil
	}
}

// convertTwoDigitsIntoWord("10") -> "on"
// convertTwoDigitsIntoWord("25") -> "iyirmi beş"
func convertTwoDigitsIntoWord(twoDigitsWord string) (string, error) {
	if len(twoDigitsWord) != 2 {
		return "", ErrInvalidArgument
	}

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
	if len(threeDigitsWord) != 3 {
		return "", ErrInvalidArgument
	}

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
			return "", ErrUnexpectedBehaviour
		}

		if len(word) != 0 {
			textBuilder = append([]string{word}, textBuilder...)
		}
	}

	return strings.Join(textBuilder, " "), nil
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
