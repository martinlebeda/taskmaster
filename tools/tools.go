package tools

import "time"

// base format for date and time in minutes
const BaseDateTimeFormat = "2006-01-02 15:04"

// get zero datetime (1.1.0001 00:00) for initialize variable
func GetZeroTime() time.Time {
	return time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// check if exists error and throw panic
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// parse time from BaseDateTimeFormat (2006-01-02 15:04) in locale location
func ParseDateTimeMinutes(dateTimeStr string) (time.Time, error) {
	return time.ParseInLocation(BaseDateTimeFormat, dateTimeStr, time.Now().Location())

}
