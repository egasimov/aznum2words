package aznum2words

import (
	"github.com/egasimov/aznum2words/fixtures"
	"reflect"
	"strconv"
	"testing"
)

func Test_convertIntPart_WhereTwoDigitsNumber(t *testing.T) {
	twoDigitsCases := fixtures.ConvertIntPartsForTwoDigitsNumbers()
	for _, testcase := range twoDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testcase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testcase.Description, err)
		}

		if !reflect.DeepEqual(actual, testcase.Expected) {
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testcase.Description,
				testcase.Given, len(strconv.Itoa(testcase.Given)),
				testcase.Expected, len(testcase.Expected),
				actual, len(actual))
		}
	}
}

func Test_convertIntPart_WhereSingleDigitsNumber(t *testing.T) {
	singleDigitsCases := fixtures.ConvertIntPartsForSingleDigitsNumbers()
	for _, testCase := range singleDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}
		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))
		}
	}
}

func Test_convertIntPart_WhereThreeDigitsNumber(t *testing.T) {
	threeDigitsCases := fixtures.ConvertIntPartsForThreeDigitsNumbers()
	for _, testCase := range threeDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereFourDigitsNumber(t *testing.T) {
	fourDigitsCases := fixtures.ConvertIntPartsForFourDigitsNumbers()
	for _, testCase := range fourDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereFiveDigitsNumber(t *testing.T) {
	fiveDigitsCases := fixtures.ConvertIntPartsForFiveDigitsNumbers()
	for _, testCase := range fiveDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereSixDigitsNumber(t *testing.T) {
	sixDigitsCases := fixtures.ConvertIntPartsForSixDigitsNumbers()
	for _, testCase := range sixDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereSevenDigitsNumber(t *testing.T) {
	sevenDigitsCases := fixtures.ConvertIntPartsForSevenDigitsNumbers()
	for _, testCase := range sevenDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereEightDigitsNumber(t *testing.T) {
	eightDigitsCases := fixtures.ConvertIntPartsForEightDigitsNumbers()
	for _, testCase := range eightDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereNineDigitsNumber(t *testing.T) {
	nineDigitsCases := fixtures.ConvertIntPartsForNineDigitsNumbers()
	for _, testCase := range nineDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereTenDigitsNumber(t *testing.T) {
	tenDigitsCases := fixtures.ConvertIntPartsForTenDigitsNumbers()
	for _, testCase := range tenDigitsCases {
		actual, err := convertIntPart(strconv.Itoa(testCase.Given))
		if err != nil {
			t.Errorf("For %s, err: %v", testCase.Description, err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(strconv.Itoa(testCase.Given)),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_convertOneDigitIntoWord(t *testing.T) {
	testCases := fixtures.ConvertOneDigitIntoWord()

	for _, testCase := range testCases {
		actual, err := convertOneDigitIntoWord(testCase.Given)

		if err != nil {
			if testCase.IsErrExpected {
				continue
			}
			t.Error("For", "Testing of  `convertOneDigitIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n İsErrExpected: ", testCase.IsErrExpected,
				"\n Got Result: ", actual,
				"\n Got Err: ", err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Error("For", "Testing of  `convertOneDigitIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n Got: ", actual)
		}
	}
}

func Test_convertTwoDigitsIntoWord(t *testing.T) {
	testCases := fixtures.ConvertTwoDigitsIntoWord()

	for _, testCase := range testCases {
		actual, err := convertTwoDigitsIntoWord(testCase.Given)

		if err != nil {
			if testCase.IsErrExpected {
				continue
			}
			t.Error("For", "Testing of  `convertOneDigitIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n İsErrExpected: ", testCase.IsErrExpected,
				"\n Got Result: ", actual,
				"\n Got Err: ", err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Error("For", "Testing of  `convertOneDigitIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n Got: ", actual)
		}
	}
}

func Test_convertThreeDigitsIntoWord(t *testing.T) {
	testCases := fixtures.ConvertThreeDigitsIntoWord()

	for _, testCase := range testCases {
		actual, err := convertThreeDigitsIntoWord(testCase.Given)

		if err != nil {
			if testCase.IsErrExpected {
				continue
			}
			t.Error("For", "Testing of  `convertThreeDigitsIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n İsErrExpected: ", testCase.IsErrExpected,
				"\n Got Result: ", actual,
				"\n Got Err: ", err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Error("For", "Testing of  `convertThreeDigitsIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n Got: ", actual)
		}
	}
}

func Test_tripleToWord(t *testing.T) {
	testCases := fixtures.TripleToWord()

	for _, testCase := range testCases {
		actual, err := tripleToWord(testCase.Given)

		if err != nil {
			if testCase.IsErrExpected {
				continue
			}
			t.Error("For", "Testing of  `convertOneDigitIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n İsErrExpected: ", testCase.IsErrExpected,
				"\n Got Result: ", actual,
				"\n Got Err: ", err)
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Error("For", "Testing of  `convertOneDigitIntoWord` ",
				"\n Given: ", testCase.Given,
				"\n Expected: ", testCase.Expected,
				"\n Got: ", actual)
		}
	}
}
