package service

import (
	"testing"
)

func TestRemovePrioFromDesc(t *testing.T) {

	if RemovePrioFromDesc("xxx xxx") != "xxx xxx" {
		t.Error("remove from desc without priority")
	}

	if RemovePrioFromDesc("xxx (A) xxx") != "xxx (A) xxx" {
		t.Error("remove from desc without priority")
	}

	desc := RemovePrioFromDesc("(A) xxx xxx")
	if desc != "xxx xxx" {
		t.Error("remove from desc with A priority: ", desc)
	}

	//no shift
	//timeResult := getDateTime("", "17:10", "2018-07-11")
	//
	//if timeResult.Format("2006-01-02 15:04") != "2018-07-11 17:10" {
	//	t.Error("Expected 2018-07-11 17:10, got ", timeResult.Format("2006-01-02 15:04"))
	//}
	//
	//with shift
	//timeResult = getDateTime("-1h", "17:10", "2018-07-11")
	//
	//if timeResult.Format("2006-01-02 15:04") != "2018-07-11 16:10" {
	//	t.Error("Expected 2018-07-11 16:10, got ", timeResult.Format("2006-01-02 15:04"))
	//}

}
