package types

import (
	"testing"
)

var strNums = map[string]bool{
	"123":      true,
	".123":     false,
	"0.123":    true,
	"123456.0": true,
	"1.2034":   true,
	"1a.2":     false,
	"aaa":      false,
}

var strIdentifier = map[string]bool{
	"&&":     true,
	"||a":    false,
	"<=1":    false,
	"asdw":   true,
	"1.2034": false,
	".9":     false,
	"--":     false,
	"*":      true,
}

var strString = map[string]bool{
	"&&":             false,
	`"adwda\""`:      true,
	"<=1":            false,
	`"asdw"`:         true,
	"1.2034":         false,
	"\".9\"":         true,
	`"asadw""`:       false,
	`a = 3`:          false,
	`"aadw\\\n\"dw"`: true,
}

func TestRegxNum(t *testing.T) {
	t.Log("\ntest strNums")
	for str, flag := range strNums {
		match := RegNumber.MatchString(str)
		if flag == match {
			t.Log("pass\t", match, flag, str)
		} else {
			t.Error("fail\t", match, flag, str)
		}
	}

	t.Log("\ntest strIdentifier")
	for str, flag := range strIdentifier {
		match := RegIdentifier.MatchString(str)
		if flag == match {
			t.Log("pass\t", match, flag, str)
		} else {
			t.Error("fail\t", match, flag, str)
		}
	}
	t.Log("\ntest strString")
	for str, flag := range strString {
		match := RegString.MatchString(str)
		if flag == match {
			t.Log("pass\t", match, flag, str)
		} else {
			t.Error("fail\t", match, flag, str)
		}
	}
}
