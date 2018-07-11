package service

import (
	"testing"
)

func TestGetDateTime(t *testing.T) {

	// no shift
	timeResult := getDateTime("", "17:10", "2018-07-11")

	if timeResult.Format("2006-01-02 15:04") != "2018-07-11 17:10" {
		t.Error("Expected 2018-07-11 17:10, got ", timeResult.Format("2006-01-02 15:04"))
	}

	// with shift
	timeResult = getDateTime("-1h", "17:10", "2018-07-11")

	if timeResult.Format("2006-01-02 15:04") != "2018-07-11 16:10" {
		t.Error("Expected 2018-07-11 16:10, got ", timeResult.Format("2006-01-02 15:04"))
	}

}
