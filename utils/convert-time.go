package utils

import (
	"fmt"
	"time"
)

func ConvertStringDateIntoGolangDateTime(date string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("Error while parsing time, ", err)
	}
	return parsedTime
}
