package aznum2words

import (
	"regexp"
	"strings"
)

var validNumberRegex = regexp.MustCompile("^-?\\d+(\\.\\d+)?$")

func validateInput(numberAsStr string) error {
	if ok := validNumberRegex.MatchString(numberAsStr); !ok {
		return ErrInvalidArgument
	}
	return nil
}

// checkConstraints as name explains, intended to check predefined constraints for given input number
// basically checks whether given number satisfy the limitations for conversion
// There is two basic limitation:
// Max allowed precision for real number should be less or equal to 15.
// Max allowed number digits - only ones from left side of decimal point should be less or equal to 66.
func checkConstraints(numberAsStr string) error {
	numberAsStr = removeSignMarkIfExists(numberAsStr)
	isFloatingNumber := strings.Contains(numberAsStr, ".")

	if !isFloatingNumber {
		if len(numberAsStr) > MaxNumberDigitCount {
			return ErrTooBigNumber

		}
		return nil
	}

	slices := strings.Split(numberAsStr, ".")

	//check number of digits on the left side of decimal point
	if len(slices[0]) > MaxNumberDigitCount {
		return ErrTooBigNumber
	}

	//sanitize floatingPart by removing trailing zeros if exist
	floatingPart := strings.TrimRight(slices[1], "0")
	if len(floatingPart) > MaxNumberScaleCount {
		return ErrTooBigScale
	}

	return nil
}
