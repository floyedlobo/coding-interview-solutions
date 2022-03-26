package main

import (
	"reflect"
	"testing"
)

func TestSortedSquaredArray(t *testing.T) {
	input := []int{-10, 2, 3, 5, 6, 8, 9}
	expected := []int{4, 9, 25, 36, 64, 81, 100}

	output := SortedSquaredArray(input)

	if !reflect.DeepEqual(expected, output) {
		t.Error("Expected array did not match output")
	}
}
