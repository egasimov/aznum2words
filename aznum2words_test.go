package aznum2words

import (
	"fmt"
	"github.com/egasimov/aznum2words/fixtures"
	"reflect"
	"testing"
)

func Test_SpellNumber_WhereNegativeIntegerNumber(t *testing.T) {

	testCases := fixtures.SpellNumberForNegativeIntegerNumbers()
	for _, testCase := range testCases {
		actual, err := SpellNumber(testCase.Given)

		if err != nil {
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s",
				testCase.Description,
				testCase.Given, len(testCase.Given),
				testCase.Expected, len(testCase.Expected),
				err.Error())
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(testCase.Given),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_SpellNumber_WherePositiveIntegerNumber(t *testing.T) {

	testCases := fixtures.SpellNumberForPositiveIntegerNumbers()
	for _, testCase := range testCases {
		actual, err := SpellNumber(testCase.Given)

		if err != nil {
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s",
				testCase.Description,
				testCase.Given, len(testCase.Given),
				testCase.Expected, len(testCase.Expected),
				err.Error())
			continue
		}

		if !reflect.DeepEqual(actual, testCase.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCase.Description,
				testCase.Given, len(testCase.Given),
				testCase.Expected, len(testCase.Expected),
				actual, len(actual))

		}
	}
}

func Test_SpellNumber_WhereNegativeFloatingPointNumber(t *testing.T) {

	testCases := fixtures.SpellNumberForNegativeFloatingPointNumbers()
	for _, testCaseSpellNumber := range testCases {
		actual, err := SpellNumber(testCaseSpellNumber.Given)

		if err != nil {
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s",
				testCaseSpellNumber.Description,
				testCaseSpellNumber.Given, len(testCaseSpellNumber.Given),
				testCaseSpellNumber.Expected, len(testCaseSpellNumber.Expected),
				err.Error())
		}

		if !reflect.DeepEqual(actual, testCaseSpellNumber.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseSpellNumber.Description,
				testCaseSpellNumber.Given, len(testCaseSpellNumber.Given),
				testCaseSpellNumber.Expected, len(testCaseSpellNumber.Expected),
				actual, len(actual))

		}
	}
}

func Test_SpellNumber_WhereSpecialInputProvided(t *testing.T) {

	testCases := fixtures.SpellNumberForSpecialInput()
	for _, testCaseSpellNumber := range testCases {
		actual, err := SpellNumber(testCaseSpellNumber.Given)

		if err != nil {
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s",
				testCaseSpellNumber.Description,
				testCaseSpellNumber.Given, len(testCaseSpellNumber.Given),
				testCaseSpellNumber.Expected, len(testCaseSpellNumber.Expected),
				err.Error())
		}

		if !reflect.DeepEqual(actual, testCaseSpellNumber.Expected) {
			//t.Error("For", testCase.Description,
			//	"\n Given: ", testCase.Given,
			//	"\n Expected: ", testCase.Expected,
			//	"\n Got: ", actual)
			t.Errorf("For %s "+
				"\n Given: %s, len: %d "+
				"\n Expected: %s, len: %d"+
				"\n Got: %s, len: %d",
				testCaseSpellNumber.Description,
				testCaseSpellNumber.Given, len(testCaseSpellNumber.Given),
				testCaseSpellNumber.Expected, len(testCaseSpellNumber.Expected),
				actual, len(actual))

		}
	}
}

func BenchmarkSpellNumber(b *testing.B) {
	dataTable := []struct {
		input string
	}{
		{
			input: "12345",
		},
		{
			input: "1234567890",
		},
		{
			input: "1234567890123456789012345"},
		{
			input: "12345678901234567890123456789012345678901234567890",
		},
		{
			input: "493882371553121860890561055192142938414552660618128252927700430053",
		},
	}

	for _, v := range dataTable {
		testName := fmt.Sprintf("digitCnt-%d_gomaxproc", len(v.input))
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SpellNumber(v.input)
			}
		})
	}
}
