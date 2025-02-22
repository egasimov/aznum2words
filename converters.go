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

// Remove the sign mark:('-' or '+') from str
// removeSignMarkIfExists("-123") -> "123"
// removeSignMarkIfExists("+123") -> "123"
// removeSignMarkIfExists("123") -> "123"
func removeSignMarkIfExists(str string) string {
	if len(str) > 1 && (str[0:1] == "-" || str[0:1] == "+") {
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
func convertIntPart(strArg string) (string, error) {
	str := removeSignMarkIfExists(strArg)
	length := len(str)

	if length == 0 {
		return "", nil
	}

	// Calculate how many triples will be processed
	tripleCount := (length + 2) / 3
	words := make([]string, tripleCount*2) // Worst case: each triple + a keyword
	wordIndex := tripleCount*2 - 1         // Fill from the end to avoid reversing

	pos := 1
	for i := length; i > 0; i -= 3 {
		key, _ := getKeyword(pos)
		pos++

		leftPointer := i - 3
		if leftPointer < 0 {
			leftPointer = 0
		}

		tripleAsWords, err := tripleToWord(str[leftPointer:i])
		if err != nil {
			return "", err
		}

		if tripleAsWords != "" {
			if key != "" {
				words[wordIndex] = key
				wordIndex--
			}
			words[wordIndex] = tripleAsWords
			wordIndex--
		}
	}

	// Construct the final string using strings.Builder
	var builder strings.Builder
	for i := wordIndex + 1; i < len(words); i++ {
		if words[i] != "" {
			if builder.Len() > 0 {
				builder.WriteByte(' ')
			}
			builder.WriteString(words[i])
		}
	}

	return builder.String(), nil
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

	var builder strings.Builder
	first := true // To manage spacing between words

	// Process tens place (index 0)
	if val, err := strconv.Atoi(twoDigitsWord[0:1]); err != nil {
		return "", err
	} else if val != 0 {
		builder.WriteString(tens[val])
		first = false
	}

	// Process ones place (index 1)
	if val, err := strconv.Atoi(twoDigitsWord[1:]); err != nil {
		return "", err
	} else if val != 0 {
		if !first {
			builder.WriteString(" ")
		}
		builder.WriteString(digits[val])
	}

	return builder.String(), nil
}

// convertThreeDigitsIntoWord("390") -> "üç yüz doxsan"
// convertThreeDigitsIntoWord("125") -> "bir yüz iyirmi beş"
func convertThreeDigitsIntoWord(threeDigitsWord string) (string, error) {
	if len(threeDigitsWord) != 3 {
		return "", ErrInvalidArgument
	}

	// Use a strings.Builder to avoid heap allocations during string construction
	var builder strings.Builder
	first := true // To manage spacing between words

	// Process hundreds place
	if val, err := strconv.Atoi(threeDigitsWord[0:1]); err != nil {
		return "", err
	} else if val != 0 {
		builder.WriteString(digits[val])
		builder.WriteString(" ")
		builder.WriteString(HundredAsString)
		first = false
	}

	// Process tens place
	if val, err := strconv.Atoi(threeDigitsWord[1:2]); err != nil {
		return "", err
	} else if val != 0 {
		if !first {
			builder.WriteString(" ")
		}
		builder.WriteString(tens[val])
		first = false
	}

	// Process ones place
	if val, err := strconv.Atoi(threeDigitsWord[2:]); err != nil {
		return "", err
	} else if val != 0 {
		if !first {
			builder.WriteString(" ")
		}
		builder.WriteString(digits[val])
	}

	return builder.String(), nil
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
