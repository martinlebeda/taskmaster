package tools

import "time"

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
