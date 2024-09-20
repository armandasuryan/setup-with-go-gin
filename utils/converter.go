package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func Base64Encoded(Value string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(Value))
	return encoded
}

func Base64Decoded(encodedValue string) (string, error) {
	decodedBytes, _ := base64.StdEncoding.DecodeString(encodedValue)
	return string(decodedBytes), nil
}

// converting time
func ConvertStringToTime(date string) time.Time {
	var dateFormat string

	fmt.Println("date:", date)

	if strings.Contains(date, "T") {
		dateFormat = "2006-01-02T15:04:05Z"
	} else {
		dateFormat = "2006-01-02 15:04:05"
	}

	// Parse the input string to time.Time
	parsedTime, err := time.Parse(dateFormat, date)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Return the parsed time
	return parsedTime
}

// Get time as string
func ConvertTimeToString(datetime time.Time, rules string) string {

	// Format the time
	if rules == "datetime" {
		formattedTime := datetime.Format("at Jan 2, 2006 15:04:05")
		return formattedTime
	}

	if rules == "default" {
		formattedTime := datetime.Format("2006-01-02 15:04:05")
		return formattedTime
	}

	if rules == "normal" {
		formattedTime := datetime.Format("01-02-2006 15:04:05")
		return formattedTime
	}

	if rules == "fullname" {
		formattedTime := datetime.Format("02 January 2006")
		return formattedTime
	}

	// default value as string "-"
	formatedTime := datetime.Format("-")
	return formatedTime
}

func JsonParseString(jsonString string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		return nil, err
	}
	return data, nil
}
