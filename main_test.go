package main

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	expected := []string{"abcde", "abcd", "abc", "ab", "a", "bcde", "bcd", "bc", "b", "cde", "cd", "c", "de", "d", "e"}
	input := "abcde"
	var actual []string
	length := len(input)
	for i := 0; i <= length; i++ {
		str := removeRightChars(input)
		actual = append(actual, str...)
		input = removeLeftChars(input)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected and Actual are not equal")

	}
}
