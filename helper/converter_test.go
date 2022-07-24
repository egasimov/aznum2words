package helper

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
		actual := ConvertIntPart(testCaseConvertIntPart.given)

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
		actual := ConvertIntPart(testCaseConvertIntPart.given)

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
		actual := ConvertIntPart(testCaseConvertIntPart.given)

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
		actual := ConvertIntPart(testCaseConvertIntPart.given)

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
		actual := ConvertIntPart(testCaseConvertIntPart.given)

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

	fiveDigitsCases := testCaseSpellNumberForNegativeIntegerNumbers()
	for _, testCaseConvertIntPart := range fiveDigitsCases {
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
		//testCaseRegexValid{
		//	given:    "10,4.5",
		//	expected: true,
		//},
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
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       0,
			expected:    ZeroAsString,
		},
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       1,
			expected:    OneAsString,
		},
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       2,
			expected:    TwoAsString,
		},
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       3,
			expected:    ThreeAsString,
		},
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       4,
			expected:    FourAsString,
		}, testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       5,
			expected:    FiveAsString,
		},
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       6,
			expected:    SixAsString,
		},
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       7,
			expected:    SevenAsString,
		}, testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       8,
			expected:    EightAsString,
		},
		testCaseConvertIntPart{
			description: "Test Case: Single digit number",
			given:       9,
			expected:    NineAsString,
		},
	}
}

func testCaseConvertIntPartsForTwoDigitsNumbers() []testCaseConvertIntPart {
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       10,
			expected:    "on",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       11,
			expected:    "on bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       19,
			expected:    "on doqquz",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       27,
			expected:    "iyirmi yeddi",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       32,
			expected:    "otuz iki",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       48,
			expected:    "qırx səkkiz",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       54,
			expected:    "əlli dörd",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       67,
			expected:    "altmış yeddi",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       70,
			expected:    "yetmiş",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
			given:       82,
			expected:    "səksən iki",
		},
		testCaseConvertIntPart{
			description: "Test Case: Single two digits number",
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
	return []testCaseConvertIntPart{
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       10000,
			expected:    "on min",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       11111,
			expected:    "on bir min bir yüz on bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       10192,
			expected:    "on min bir yüz doxsan iki",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       27021,
			expected:    "iyirmi yeddi min iyirmi bir",
		},
		testCaseConvertIntPart{
			description: "Test Case: four digits number",
			given:       13210,
			expected:    "on üç min iki yüz on",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       48550,
			expected:    "qırx səkkiz min beş yüz əlli",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       95412,
			expected:    "doxsan beş min dörd yüz on iki",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       79594,
			expected:    "yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       86070,
			expected:    "səksən altı min yetmiş",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       70010,
			expected:    "yetmiş min on",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       81200,
			expected:    "səksən bir min iki yüz",
		},
		testCaseConvertIntPart{
			description: "Test Case: three digits number",
			given:       10001,
			expected:    "on min bir",
		}}
}

func testCaseSpellNumberForNegativeIntegerNumbers() []testCaseSpellNumber {
	return []testCaseSpellNumber{
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-10000",
			expected:    "mənfi on min",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-11111",
			expected:    "mənfi on bir min bir yüz on bir",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-10192",
			expected:    "mənfi on min bir yüz doxsan iki",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-27021",
			expected:    "mənfi iyirmi yeddi min iyirmi bir",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-13210",
			expected:    "mənfi on üç min iki yüz on",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-48550",
			expected:    "mənfi qırx səkkiz min beş yüz əlli",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-95412",
			expected:    "mənfi doxsan beş min dörd yüz on iki",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-79594",
			expected:    "mənfi yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-86070",
			expected:    "mənfi səksən altı min yetmiş",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-700",
			expected:    "mənfi yeddi yüz",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-81",
			expected:    "mənfi səksən bir",
		},
		testCaseSpellNumber{
			description: "Test Case: negative number",
			given:       "-10000",
			expected:    "mənfi on min",
		}}
}
