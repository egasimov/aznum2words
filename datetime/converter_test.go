package datetime

import (
	"errors"
	"testing"
)

func TestConvertDate(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{
			input:    "2024-04-24",
			expected: "iyirmi dörd aprel iki min iyirmi dörd",
			err:      nil,
		},
		{
			input:    "1990-01-01",
			expected: "bir yanvar min doqquz yüz doxsan",
			err:      nil,
		},
		{
			input:    "invalid-date",
			expected: "",
			err:      errors.New("invalid date format"),
		},
		{
			input:    "2024-13-01", // Invalid month
			expected: "",
			err:      errors.New("invalid date format"),
		},
		{
			input:    "2024-02-30", // Invalid day
			expected: "",
			err:      errors.New("invalid date format"),
		},
	}

	for _, tc := range testCases {
		result, err := ConvertDate(tc.input)
		if tc.err == nil {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.input, err)
				continue
			}
			if result != tc.expected {
				t.Errorf("For input %s, expected %s, got %s", tc.input, tc.expected, result)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error for input %s, got none", tc.input)
			}
		}
	}
}

func TestConvertTime(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{
			input:    "14:30",
			expected: "on dörd otuz",
			err:      nil,
		},
		{
			input:    "09:15",
			expected: "doqquz on beş",
			err:      nil,
		},
		{
			input:    "00:00",
			expected: "sıfır sıfır",
			err:      nil,
		},
		{
			input:    "23:59",
			expected: "iyirmi üç əlli doqquz",
			err:      nil,
		},
		{
			input:    "invalid-time",
			expected: "",
			err:      errors.New("invalid time format"),
		},
		{
			input:    "25:00", // Invalid hour
			expected: "",
			err:      errors.New("invalid hour: must be between 0 and 23"),
		},
		{
			input:    "12:60", // Invalid minute
			expected: "",
			err:      errors.New("invalid minute: must be between 0 and 59"),
		},
	}

	for _, tc := range testCases {
		result, err := ConvertTime(tc.input)
		if tc.err == nil {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.input, err)
				continue
			}
			if result != tc.expected {
				t.Errorf("For input %s, expected %s, got %s", tc.input, tc.expected, result)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error for input %s, got none", tc.input)
			}
		}
	}
}

func TestConvertDateTime(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{
			input:    "2024-04-24 14:30",
			expected: "iyirmi dörd aprel iki min iyirmi dörd, on dörd otuz",
			err:      nil,
		},
		{
			input:    "1990-01-01 09:15",
			expected: "bir yanvar min doqquz yüz doxsan, doqquz on beş",
			err:      nil,
		},
		{
			input:    "invalid-datetime",
			expected: "",
			err:      errors.New("invalid datetime format"),
		},
		{
			input:    "2024-04-24", // Missing time
			expected: "",
			err:      errors.New("invalid datetime format"),
		},
		{
			input:    "14:30", // Missing date
			expected: "",
			err:      errors.New("invalid datetime format"),
		},
	}

	for _, tc := range testCases {
		result, err := ConvertDateTime(tc.input)
		if tc.err == nil {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.input, err)
				continue
			}
			if result != tc.expected {
				t.Errorf("For input %s, expected %s, got %s", tc.input, tc.expected, result)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error for input %s, got none", tc.input)
			}
		}
	}
}
