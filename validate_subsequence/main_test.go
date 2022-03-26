package main

import (
	"testing"
)

func TestValidateSubsequence(t *testing.T) {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	subArr := []int{1, 6, -1, 10}

	if IsValidSubsequence(arr, subArr) != true {
		t.Errorf("Expected true but output is false")
	}
}
