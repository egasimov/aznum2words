package aznum2words

import (
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
