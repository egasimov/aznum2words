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


func Test_convertIntPart_WhereFourDigitsNumber(t *testing.T) {
	fourDigitsCases := testCasesForFourDigitsNumbers()
	for _, testCase := range fourDigitsCases {
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

func Test_convertIntPart_WhereFiveDigitsNumber(t *testing.T) {
	fiveDigitsCases := testCasesForFiveDigitsNumbers()
	for _, testCase := range fiveDigitsCases {
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

func testCasesForFourDigitsNumbers() []testCase {
	return []testCase{
		testCase{
			description: "Test Case: four digits number",
			given:       "1000",
			expected:    "bir min",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "1111",
			expected:    "bir min bir yüz on bir",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "1019",
			expected:    "bir min on doqquz",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "2701",
			expected:    "iki min yeddi yüz bir",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "3210",
			expected:    "üç min iki yüz on",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "4850",
			expected:    "dörd min səkkiz yüz əlli",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "5412",
			expected:    "beş min dörd yüz on iki",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "9594",
			expected:    "doqquz min beş yüz doxsan dörd",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "8670",
			expected:    "səkkiz min altı yüz yetmiş",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "7001",
			expected:    "yeddi min bir",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "8200",
			expected:    "səkkiz min iki yüz",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "9526",
			expected:    "doqquz min beş yüz iyirmi altı",
		}}
}

func testCasesForFiveDigitsNumbers() []testCase {
	return []testCase{
		testCase{
			description: "Test Case: four digits number",
			given:       "10000",
			expected:    "on min",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "11111",
			expected:    "on bir min bir yüz on bir",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "10192",
			expected:    "on min bir yüz doxsan iki",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "27021",
			expected:    "iyirmi yeddi min iyirmi bir",
		},
		testCase{
			description: "Test Case: four digits number",
			given:       "13210",
			expected:    "on üç min iki yüz on",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "48550",
			expected:    "qırx səkkiz min beş yüz əlli",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "95412",
			expected:    "doxsan beş min dörd yüz on iki",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "79594",
			expected:    "yetmiş doqquz min beş yüz doxsan dörd",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "86070",
			expected:    "səksən altı min yetmiş",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "70010",
			expected:    "yetmiş min on",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "81200",
			expected:    "səksən bir min iki yüz",
		},
		testCase{
			description: "Test Case: three digits number",
			given:       "10001",
			expected:    "on min bir",
		}}
}