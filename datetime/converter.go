package datetime

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/egasimov/aznum2words"
)

// ConvertDate converts a date to words in Azerbaijani
// Example: "2024-04-24" -> "iyirmi dörd aprel iki min iyirmi dörd"
func ConvertDate(dateStr string) (string, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %v", err)
	}

	day := date.Day()
	month := date.Month()
	year := date.Year()

	// Convert day to words
	dayWords, err := aznum2words.SpellNumber(strconv.Itoa(day))
	if err != nil {
		return "", err
	}

	// Get month name
	monthName := Months[month-1]

	// Convert year to words
	yearWords, err := aznum2words.SpellNumber(strconv.Itoa(year))
	if err != nil {
		return "", err
	}

	// Remove extra "bir" from year if it starts with "bir"
	if strings.HasPrefix(yearWords, "bir ") {
		yearWords = yearWords[4:]
	}

	// Combine all parts
	result := fmt.Sprintf("%s %s %s", dayWords, monthName, yearWords)
	return result, nil
}

// ConvertTime converts a time to words in Azerbaijani
// Example: "14:30" -> "on dörd otuz"
func ConvertTime(timeStr string) (string, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid time format")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid hour: %v", err)
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid minute: %v", err)
	}

	// Validate hour and minute ranges
	if hour < 0 || hour > 23 {
		return "", fmt.Errorf("invalid hour: must be between 0 and 23")
	}
	if minute < 0 || minute > 59 {
		return "", fmt.Errorf("invalid minute: must be between 0 and 59")
	}

	// Convert hour to words
	hourWords, err := aznum2words.SpellNumber(strconv.Itoa(hour))
	if err != nil {
		return "", err
	}

	// Convert minute to words
	minuteWords, err := aznum2words.SpellNumber(strconv.Itoa(minute))
	if err != nil {
		return "", err
	}

	// Combine all parts
	result := fmt.Sprintf("%s %s", hourWords, minuteWords)
	return result, nil
}

// ConvertDateTime converts both date and time to words in Azerbaijani
// Example: "2024-04-24 14:30" -> "iyirmi dörd aprel iki min iyirmi dörd, on dörd otuz"
func ConvertDateTime(datetimeStr string) (string, error) {
	parts := strings.Split(datetimeStr, " ")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid datetime format")
	}

	dateWords, err := ConvertDate(parts[0])
	if err != nil {
		return "", err
	}

	timeWords, err := ConvertTime(parts[1])
	if err != nil {
		return "", err
	}

	// Combine date and time
	result := fmt.Sprintf("%s, %s", dateWords, timeWords)
	return result, nil
}
