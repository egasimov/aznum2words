package helper

import (
	"reflect"
	"testing"
)

type testCase struct {
	description string
	given       string
	expected    string
}

func Test_convertIntPart_WhereTwoDigitsNumber(t *testing.T) {
	twoDigitsCases := testCasesForTwoDigitsNumbers()
	for _, testCase := range twoDigitsCases {
		actual := ConvertIntPart(testCase.given)

		if !reflect.DeepEqual(actual, testCase.expected) {
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

func Test_convertIntPart_WhereSingleDigitsNumber(t *testing.T) {
	singleDigitsCases := testCasesForSingleDigitsNumbers()
	for _, testCase := range singleDigitsCases {
		actual := ConvertIntPart(testCase.given)

		if !reflect.DeepEqual(actual, testCase.expected) {
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

func Test_convertIntPart_WhereThreeDigitsNumber(t *testing.T) {
	threeDigitsCases := testCasesForThreeDigitsNumbers()
	for _, testCase := range threeDigitsCases {
		actual := ConvertIntPart(testCase.given)

		if !reflect.DeepEqual(actual, testCase.expected) {
			//t.Error("For", testCase.description,
			//	"\n Given: ", testCase.given,
			//	"\n Expected: ", testCase.expected,
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

func testCasesForSingleDigitsNumbers() []testCase {
	return []testCase{
		testCase{
			description: "Test Case: Single digit number",
			given:       "0",
			expected:    ZeroAsString,
		},
		testCase{
			description: "Test Case: Single digit number",
			given:       "1",
			expected:    OneAsString,
		},
		testCase{
			description: "Test Case: Single digit number",
			given:       "2",
			expected:    TwoAsString,
		},
		testCase{
			description: "Test Case: Single digit number",
			given:       "3",
			expected:    ThreeAsString,
		},
		testCase{
			description: "Test Case: Single digit number",
			given:       "4",
			expected:    FourAsString,
		}, testCase{
			description: "Test Case: Single digit number",
			given:       "5",
			expected:    FiveAsString,
		},
		testCase{
			description: "Test Case: Single digit number",
			given:       "6",
			expected:    SixAsString,
		},
		testCase{
			description: "Test Case: Single digit number",
			given:       "7",
			expected:    SevenAsString,
		}, testCase{
			description: "Test Case: Single digit number",
			given:       "8",
			expected:    EightAsString,
		},
		testCase{
			description: "Test Case: Single digit number",
			given:       "9",
			expected:    NineAsString,
		},
	}
}

func testCasesForTwoDigitsNumbers() []testCase {
	return []testCase{
		testCase{
			description: "Test Case: Single two digits number",
			given:       "10",
			expected:    "on",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "11",
			expected:    "on bir",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "19",
			expected:    "on doqquz",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "27",
			expected:    "iyirmi yeddi",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "32",
			expected:    "otuz iki",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "48",
			expected:    "qırx səkkiz",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "54",
			expected:    "əlli dörd",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "54",
			expected:    "əlli dörd",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "67",
			expected:    "altmış yeddi",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "70",
			expected:    "yetmiş",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "82",
			expected:    "səksən iki",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "96",
			expected:    "doxsan altı",
		}}
}

func testCasesForThreeDigitsNumbers() []testCase {
	return []testCase{
		testCase{
			description: "Test Case: Single three digits number",
			given:       "100",
			expected:    "bir yüz",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "111",
			expected:    "bir yüz on bir",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "119",
			expected:    "bir yüz on doqquz",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "271",
			expected:    "iki yüz yetmiş bir",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "321",
			expected:    "üç yüz iyirmi bir",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "485",
			expected:    "dörd yüz səksən beş",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "541",
			expected:    "beş yüz qırx bir",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "594",
			expected:    "beş yüz doxsan dörd",
		},
		testCase{
			description: "Test Case: Single three digits number",
			given:       "670",
			expected:    "altı yüz yetmiş",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "701",
			expected:    "yeddi yüz bir",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "820",
			expected:    "səkkiz yüz iyirmi",
		},
		testCase{
			description: "Test Case: Single two digits number",
			given:       "906",
			expected:    "doqquz yüz altı",
		}}
}
