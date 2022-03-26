package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	input := []int{3, 5, -4, 8, 11, 1, -1, 6}
	expected := []int{11, -1}

	sort.Ints(expected)

	output := TwoNumberSum(input, 10)

	sort.Ints(output)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v Output %v", expected, output)
	}
}

func TestTwoNumberSum2(t *testing.T) {
	input := []int{3, 5, -4, 8, 11, 1, -1, 6}
	expected := []int{11, -1}

	sort.Ints(expected)

	output := TwoNumberSum2(input, 10)

	sort.Ints(output)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v Output %v", expected, output)
	}
}
