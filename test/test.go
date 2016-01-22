package test

import (
	"fmt"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}, msg string, args ...interface{}) {
	if expected != actual {
		msg := fmt.Sprintf(msg, args...)
		t.Errorf("%v: Expected %v but found %v", msg, expected, actual)
	}
}

func NotNil(t *testing.T, v interface{}, msg string, args ...interface{}) {
	if v == nil {
		msg := fmt.Sprintf(msg, args...)
		t.Errorf("%v: Should not be nil", msg)	
	}
}