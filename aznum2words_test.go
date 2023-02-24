package aznum2words

import (
	"reflect"
	"strconv"
	"testing"
)

type testCaseConvertIntPart struct {
	description string
	given       int
	expected    string
}

type testCaseSpellNumber struct {
	description string
	given       string
	expected    string
}

func Test_convertIntPart_WhereTwoDigitsNumber(t *testing.T) {
	twoDigitsCases := testCaseConvertIntPartsForTwoDigitsNumbers()
	for _, testCaseConvertIntPart := range twoDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))
		}
	}
}

func Test_convertIntPart_WhereSingleDigitsNumber(t *testing.T) {
	singleDigitsCases := testCaseConvertIntPartsForSingleDigitsNumbers()
	for _, testCaseConvertIntPart := range singleDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))
		}
	}
}

func Test_convertIntPart_WhereThreeDigitsNumber(t *testing.T) {
	threeDigitsCases := testCaseConvertIntPartsForThreeDigitsNumbers()
	for _, testCaseConvertIntPart := range threeDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereFourDigitsNumber(t *testing.T) {
	fourDigitsCases := testCaseConvertIntPartsForFourDigitsNumbers()
	for _, testCaseConvertIntPart := range fourDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereFiveDigitsNumber(t *testing.T) {
	fiveDigitsCases := testCaseConvertIntPartsForFiveDigitsNumbers()
	for _, testCaseConvertIntPart := range fiveDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereSixDigitsNumber(t *testing.T) {
	sixDigitsCases := testCaseConvertIntPartsForSixDigitsNumbers()
	for _, testCaseConvertIntPart := range sixDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereSevenDigitsNumber(t *testing.T) {
	sevenDigitsCases := testCaseConvertIntPartsForSevenDigitsNumbers()
	for _, testCaseConvertIntPart := range sevenDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereEightDigitsNumber(t *testing.T) {
	eightDigitsCases := testCaseConvertIntPartsForEightDigitsNumbers()
	for _, testCaseConvertIntPart := range eightDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereNineDigitsNumber(t *testing.T) {
	nineDigitsCases := testCaseConvertIntPartsForNineDigitsNumbers()
	for _, testCaseConvertIntPart := range nineDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_convertIntPart_WhereTenDigitsNumber(t *testing.T) {
	tenDigitsCases := testCaseConvertIntPartsForTenDigitsNumbers()
	for _, testCaseConvertIntPart := range tenDigitsCases {
		actual := convertIntPart(strconv.Itoa(testCaseConvertIntPart.given))

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %d, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(strconv.Itoa(testCaseConvertIntPart.given)),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_SpellNumber_WhereNegativeIntegerNumber(t *testing.T) {

	testCases := testCaseSpellNumberForNegativeIntegerNumbers()
	for _, testCaseConvertIntPart := range testCases {
		actual, err := SpellNumber(testCaseConvertIntPart.given)

		if err != nil {
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(testCaseConvertIntPart.given),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				err.Error())
		}

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseConvertIntPart.description,
				testCaseConvertIntPart.given, len(testCaseConvertIntPart.given),
				testCaseConvertIntPart.expected, len(testCaseConvertIntPart.expected),
				actual, len(actual))

		}
	}
}

func Test_SpellNumber_WherePositiveIntegerNumber(t *testing.T) {

	testCases := testCaseSpellNumberForPositiveIntegerNumbers()
	for _, testCase := range testCases {
		actual, err := SpellNumber(testCase.given)

		if err != nil {
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s",
				testCase.description,
				testCase.given, len(testCase.given),
				testCase.expected, len(testCase.expected),
				err.Error())
			continue
		}

		if !reflect.DeepEqual(actual, testCase.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.description,
				testCase.given, len(testCase.given),
				testCase.expected, len(testCase.expected),
				actual, len(actual))

		}
	}
}

func Test_SpellNumber_WhereNegativeFloatingPointNumber(t *testing.T) {

	testCases := testCaseSpellNumberForNegativeFloatingPointNumbers()
	for _, testCaseSpellNumber := range testCases {
		actual, err := SpellNumber(testCaseSpellNumber.given)

		if err != nil {
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s",
				testCaseSpellNumber.description,
				testCaseSpellNumber.given, len(testCaseSpellNumber.given),
				testCaseSpellNumber.expected, len(testCaseSpellNumber.expected),
				err.Error())
		}

		if !reflect.DeepEqual(actual, testCaseSpellNumber.expected) {
			//t.Error("For", testCaseConvertIntPart.description,
			//	"\n Given: ", testCaseConvertIntPart.given,
			//	"\n Expected: ", testCaseConvertIntPart.expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseSpellNumber.description,
				testCaseSpellNumber.given, len(testCaseSpellNumber.given),
				testCaseSpellNumber.expected, len(testCaseSpellNumber.expected),
				actual, len(actual))

		}
	}
}

func Test_validNumberRegex(t *testing.T) {

	type testCaseRegexValid struct {
		given    string
		expected bool
	}
	testCases := []testCaseRegexValid{
		testCaseRegexValid{
			given:    "10",
			expected: true,
		},
		testCaseRegexValid{
			given:    "-10",
			expected: true,
		},
		testCaseRegexValid{
			given:    "10.5",
			expected: true,
		},
		testCaseRegexValid{
			given:    "-10.5",
			expected: true,
		},
		testCaseRegexValid{
			given:    "0",
			expected: true,
		},
		testCaseRegexValid{
			given:    "0.01",
			expected: true,
		},
		testCaseRegexValid{
			given:    "-9999",
			expected: true,
		},
		testCaseRegexValid{
			given:    "1000000000.00001",
			expected: true,
		},
		testCaseRegexValid{
			given:    "some-test",
			expected: false,
		},
		testCaseRegexValid{
			given:    "12912%2",
			expected: false,
		},
	}

	for _, testCaseConvertIntPart := range testCases {
		actual := validNumberRegex.MatchString(testCaseConvertIntPart.given)

		if !reflect.DeepEqual(actual, testCaseConvertIntPart.expected) {
			t.Error("For", "Validation of  `validNumberRegex` ",
				"\n Given: ", testCaseConvertIntPart.given,
				"\n Expected: ", testCaseConvertIntPart.expected,
				"\n Got: ", actual)
		}
	}

}

func testCaseConvertIntPartsForSingleDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: Single digit number"

	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       0,
			expected:    ZeroAsString,
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1,
			expected:    OneAsString,
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       2,
			expected:    TwoAsString,
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       3,
			expected:    ThreeAsString,
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       4,
			expected:    FourAsString,
		}, testCaseConvertIntPart{
			description: testCaseDescription,
			given:       5,
			expected:    FiveAsString,
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       6,
			expected:    SixAsString,
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       7,
			expected:    SevenAsString,
		}, testCaseConvertIntPart{
			description: testCaseDescription,
			given:       8,
			expected:    EightAsString,
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       9,
			expected:    NineAsString,
		},
	}
}

func testCaseConvertIntPartsForTwoDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: Two digits number"
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       10,
			expected:    "on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       11,
			expected:    "on bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       19,
			expected:    "on doqquz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       27,
			expected:    "iyirmi yeddi",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       32,
			expected:    "otuz iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       48,
			expected:    "qırx səkkiz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       54,
			expected:    "əlli dörd",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       67,
			expected:    "altmış yeddi",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       70,
			expected:    "yetmiş",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       82,
			expected:    "səksən iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       96,
			expected:    "doxsan altı",
		}}
}

func testCaseConvertIntPartsForThreeDigitsNumbers() []testCaseConvertIntPart {
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       100,
			expected:    "bir yüz",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       111,
			expected:    "bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       119,
			expected:    "bir yüz on doqquz",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       271,
			expected:    "iki yüz yetmiş bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       321,
			expected:    "üç yüz iyirmi bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       485,
			expected:    "dörd yüz səksən beş",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       541,
			expected:    "beş yüz qırx bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       594,
			expected:    "beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single three digits number",
			given:       670,
			expected:    "altı yüz yetmiş",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       701,
			expected:    "yeddi yüz bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       820,
			expected:    "səkkiz yüz iyirmi",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       906,
			expected:    "doqquz yüz altı",
		}}
}

func testCaseConvertIntPartsForFourDigitsNumbers() []testCaseConvertIntPart {
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       1000,
			expected:    "bir min",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       1111,
			expected:    "bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       1019,
			expected:    "bir min on doqquz",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       2701,
			expected:    "iki min yeddi yüz bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       3210,
			expected:    "üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       4850,
			expected:    "dörd min səkkiz yüz əlli",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       5412,
			expected:    "beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       9594,
			expected:    "doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       8670,
			expected:    "səkkiz min altı yüz yetmiş",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       7001,
			expected:    "yeddi min bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       8200,
			expected:    "səkkiz min iki yüz",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       9526,
			expected:    "doqquz min beş yüz iyirmi altı",
		}}
}

func testCaseConvertIntPartsForFiveDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: five digits number"
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       10000,
			expected:    "on min",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       11111,
			expected:    "on bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       10192,
			expected:    "on min bir yüz doxsan iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       27021,
			expected:    "iyirmi yeddi min iyirmi bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       13210,
			expected:    "on üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       48550,
			expected:    "qırx səkkiz min beş yüz əlli",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       95412,
			expected:    "doxsan beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       79594,
			expected:    "yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       86070,
			expected:    "səksən altı min yetmiş",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       70010,
			expected:    "yetmiş min on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       81200,
			expected:    "səksən bir min iki yüz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       10001,
			expected:    "on min bir",
		}}
}

func testCaseConvertIntPartsForSixDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: six digits number"
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       100000,
			expected:    "bir yüz min",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       111111,
			expected:    "bir yüz on bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       110192,
			expected:    "bir yüz on min bir yüz doxsan iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       127021,
			expected:    "bir yüz iyirmi yeddi min iyirmi bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       113210,
			expected:    "bir yüz on üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       148550,
			expected:    "bir yüz qırx səkkiz min beş yüz əlli",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       195412,
			expected:    "bir yüz doxsan beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       179594,
			expected:    "bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       186070,
			expected:    "bir yüz səksən altı min yetmiş",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       170010,
			expected:    "bir yüz yetmiş min on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       181200,
			expected:    "bir yüz səksən bir min iki yüz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       110001,
			expected:    "bir yüz on min bir",
		}}
}

func testCaseConvertIntPartsForSevenDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: seven digits number"
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1000000,
			expected:    "bir milyon",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1111111,
			expected:    "bir milyon bir yüz on bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1110192,
			expected:    "bir milyon bir yüz on min bir yüz doxsan iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1127021,
			expected:    "bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1113210,
			expected:    "bir milyon bir yüz on üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1148550,
			expected:    "bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1195412,
			expected:    "bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1179594,
			expected:    "bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1186070,
			expected:    "bir milyon bir yüz səksən altı min yetmiş",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1170010,
			expected:    "bir milyon bir yüz yetmiş min on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1181200,
			expected:    "bir milyon bir yüz səksən bir min iki yüz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1110001,
			expected:    "bir milyon bir yüz on min bir",
		}}
}

func testCaseConvertIntPartsForEightDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: eight digits number"
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       11000000,
			expected:    "on bir milyon",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       11111111,
			expected:    "on bir milyon bir yüz on bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       11110192,
			expected:    "on bir milyon bir yüz on min bir yüz doxsan iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       11127021,
			expected:    "on bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       11113210,
			expected:    "on bir milyon bir yüz on üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       21148550,
			expected:    "iyirmi bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       31195412,
			expected:    "otuz bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       81179594,
			expected:    "səksən bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       41186070,
			expected:    "qırx bir milyon bir yüz səksən altı min yetmiş",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       91170010,
			expected:    "doxsan bir milyon bir yüz yetmiş min on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       71181200,
			expected:    "yetmiş bir milyon bir yüz səksən bir min iki yüz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       61110001,
			expected:    "altmış bir milyon bir yüz on min bir",
		}}
}

func testCaseConvertIntPartsForNineDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: eight digits number"
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       211000000,
			expected:    "iki yüz on bir milyon",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       211111111,
			expected:    "iki yüz on bir milyon bir yüz on bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       411110192,
			expected:    "dörd yüz on bir milyon bir yüz on min bir yüz doxsan iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       111127021,
			expected:    "bir yüz on bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       611113210,
			expected:    "altı yüz on bir milyon bir yüz on üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       721148550,
			expected:    "yeddi yüz iyirmi bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       831195412,
			expected:    "səkkiz yüz otuz bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       881179594,
			expected:    "səkkiz yüz səksən bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       141186070,
			expected:    "bir yüz qırx bir milyon bir yüz səksən altı min yetmiş",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       191170010,
			expected:    "bir yüz doxsan bir milyon bir yüz yetmiş min on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       171181200,
			expected:    "bir yüz yetmiş bir milyon bir yüz səksən bir min iki yüz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       961110001,
			expected:    "doqquz yüz altmış bir milyon bir yüz on min bir",
		}}
}

func testCaseConvertIntPartsForTenDigitsNumbers() []testCaseConvertIntPart {
	testCaseDescription := "Test Case: eight digits number"
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1211000000,
			expected:    "bir milyard iki yüz on bir milyon",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       3211111111,
			expected:    "üç milyard iki yüz on bir milyon bir yüz on bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       7411110192,
			expected:    "yeddi milyard dörd yüz on bir milyon bir yüz on min bir yüz doxsan iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       1111127021,
			expected:    "bir milyard bir yüz on bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       5611113210,
			expected:    "beş milyard altı yüz on bir milyon bir yüz on üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       9721148550,
			expected:    "doqquz milyard yeddi yüz iyirmi bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       8831195412,
			expected:    "səkkiz milyard səkkiz yüz otuz bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       5881179594,
			expected:    "beş milyard səkkiz yüz səksən bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       4141186070,
			expected:    "dörd milyard bir yüz qırx bir milyon bir yüz səksən altı min yetmiş",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       6091170010,
			expected:    "altı milyard doxsan bir milyon bir yüz yetmiş min on",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       7171181200,
			expected:    "yeddi milyard bir yüz yetmiş bir milyon bir yüz səksən bir min iki yüz",
		},
		testCaseConvertIntPart{
			description: testCaseDescription,
			given:       3961110001,
			expected:    "üç milyard doqquz yüz altmış bir milyon bir yüz on min bir",
		}}
}

func testCaseSpellNumberForNegativeIntegerNumbers() []testCaseSpellNumber {
	testCaseDescription := "Test Case: negative integer number"

	return []testCaseSpellNumber{
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-10000",
			expected:    "mənfi on min",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-11111",
			expected:    "mənfi on bir min bir yüz on bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-10192",
			expected:    "mənfi on min bir yüz doxsan iki",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-27021",
			expected:    "mənfi iyirmi yeddi min iyirmi bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-13210",
			expected:    "mənfi on üç min iki yüz on",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-48550",
			expected:    "mənfi qırx səkkiz min beş yüz əlli",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-95412",
			expected:    "mənfi doxsan beş min dörd yüz on iki",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-79594",
			expected:    "mənfi yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-86070",
			expected:    "mənfi səksən altı min yetmiş",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-700",
			expected:    "mənfi yeddi yüz",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-81",
			expected:    "mənfi səksən bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-10000",
			expected:    "mənfi on min",
		}}
}

func testCaseSpellNumberForPositiveIntegerNumbers() []testCaseSpellNumber {
	testCaseDescription := "Test Case: positive integer number"

	return []testCaseSpellNumber{
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "493882371553121860890561055192142938414552660618128252927700430053",
			expected:    "dörd yüz doxsan üç vigintilyon səkkiz yüz səksən iki novemdesilyon üç yüz yetmiş bir oktodesilyon beş yüz əlli üç septendesilyon bir yüz iyirmi bir seksdesilyon səkkiz yüz altmış kendesilyon səkkiz yüz doxsan katordesilyon beş yüz altmış bir tredesilyon əlli beş dodesilyon bir yüz doxsan iki undesilyon bir yüz qırx iki desilyon doqquz yüz otuz səkkiz nonilyon dörd yüz on dörd oktilyon beş yüz əlli iki septilyon altı yüz altmış sekstilyon altı yüz on səkkiz kvintilyon bir yüz iyirmi səkkiz kvadrilyon iki yüz əlli iki trilyon doqquz yüz iyirmi yeddi milyard yeddi yüz milyon dörd yüz otuz min əlli üç",
		},
	}
}

func testCaseSpellNumberForNegativeFloatingPointNumbers() []testCaseSpellNumber {
	testCaseDescription := "Test Case: negative floating point number"

	return []testCaseSpellNumber{
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-12.1",
			expected:    "mənfi on iki tam onda bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-111.11",
			expected:    "mənfi bir yüz on bir tam yüzdə on bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-10.192",
			expected:    "mənfi on tam bir mində bir yüz doxsan iki",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-2.7021",
			expected:    "mənfi iki tam on mində yeddi min iyirmi bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.13211",
			expected:    "mənfi sıfır tam yüz mində on üç min iki yüz on bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.248551",
			expected:    "mənfi sıfır tam bir milyonda iki yüz qırx səkkiz min beş yüz əlli bir",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.6195412",
			expected:    "mənfi sıfır tam on milyonda altı milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.79354594",
			expected:    "mənfi sıfır tam yüz milyonda yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.179354594",
			expected:    "mənfi sıfır tam bir milyardda bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.1179354594",
			expected:    "mənfi sıfır tam on milyardda bir milyard bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.41179354594",
			expected:    "mənfi sıfır tam yüz milyardda qırx bir milyard bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		testCaseSpellNumber{
			description: testCaseDescription,
			given:       "-0.541179354594",
			expected:    "mənfi sıfır tam bir trilyonda beş yüz qırx bir milyard bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		}}
}
