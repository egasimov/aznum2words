package aznum2words

import (
	"reflect"
	"testing"
)

func Test_checkConstraints(t *testing.T) {

	type testCaseCheckConstraint struct {
		given    string
		expected error
	}
	testCases := []testCaseCheckConstraint{
		{
			given:    "123456789123456789123456789123456789123456789123456789123456789123",
			expected: nil,
		},
		{
			given:    "-123456789123456789123456789123456789123456789123456789123456789123",
			expected: nil,
		},
		{
			given:    "1234567891234567891234567891234567891234567891234567891234567891234",
			expected: ErrTooBigNumber,
		},
		{
			given:    "-1234567891234567891234567891234567891234567891234567891234567891234",
			expected: ErrTooBigNumber,
		},
		{
			given:    "0.12345678912345",
			expected: nil,
		},
		{
			given:    "-0.123456789123456",
			expected: nil,
		},
		{
			given:    "-0.12345678912345000000",
			expected: nil,
		},
		{
			given:    "0.1234567891234567",
			expected: ErrTooBigScale,
		},
		{
			given:    "-0.1234567891234567",
			expected: ErrTooBigScale,
		},
	}

	for _, testCase := range testCases {
		actual := checkConstraints(testCase.given)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Error("For", "Testing of  `checkConstraint` ",
				"\n Given: ", testCase.given,
				"\n Expected: ", testCase.expected,
				"\n Got: ", actual)
		}
	}

}

func Test_validNumberRegex(t *testing.T) {

	type testCaseRegexValid struct {
		given    string
		expected bool
	}
	testCases := []testCaseRegexValid{
		{
			given:    "10",
			expected: true,
		},
		{
			given:    "-10",
			expected: true,
		},
		{
			given:    "10.5",
			expected: true,
		},
		{
			given:    "-10.5",
			expected: true,
		},
		{
			given:    "0",
			expected: true,
		},
		{
			given:    "0.01",
			expected: true,
		},
		{
			given:    "-9999",
			expected: true,
		},
		{
			given:    "1000000000.00001",
			expected: true,
		},
		{
			given:    "0.0",
			expected: true,
		},
		{
			given:    "00.5",
			expected: false,
		},
		{
			given:    "some-test",
			expected: false,
		},
		{
			given:    "12912%2",
			expected: false,
		},
	}

	for _, testCase := range testCases {
		actual := validateNumberRegex2.MatchString(testCase.given)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Error("For", "Validation of  `validNumberRegex` ",
				"\n Given: ", testCase.given,
				"\n Expected: ", testCase.expected,
				"\n Got: ", actual)
		}
	}
}

func Test_validateInput(t *testing.T) {

	type testCaseValidateInput struct {
		given       string
		expectedErr error
	}
	testCases := []testCaseValidateInput{
		{
			given:       "10",
			expectedErr: nil,
		},
		{
			given:       "-10",
			expectedErr: nil,
		},
		{
			given:       "10.5",
			expectedErr: nil,
		},
		{
			given:       "-10.5",
			expectedErr: nil,
		},
		{
			given:       "0",
			expectedErr: nil,
		},
		{
			given:       "0.01",
			expectedErr: nil,
		},
		{
			given:       "-9999",
			expectedErr: nil,
		},
		{
			given:       "1000000000.00001",
			expectedErr: nil,
		},
		{
			given:       "some-test",
			expectedErr: ErrInvalidArgument,
		},
		{
			given:       "12912%2",
			expectedErr: ErrInvalidArgument,
		},
	}

	for _, testCase := range testCases {
		actual := validateInput(testCase.given)

		if !reflect.DeepEqual(actual, testCase.expectedErr) {
			t.Error("For", "Validation of  `validateInput` ",
				"\n Given: ", testCase.given,
				"\n Expected: ", testCase.expectedErr,
				"\n Got: ", actual)
		}
	}
}
