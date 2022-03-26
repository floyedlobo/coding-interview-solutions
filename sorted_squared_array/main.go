package main

import "math"

//https://leetcode.com/problems/squares-of-a-sorted-array/

//Time O(n) Space O(n)
func SortedSquaredArray(array []int) []int {
	startIdx := 0
	endIdx := len(array) - 1
	index := endIdx
	output := make([]int, len(array))
	for index >= 0 {
		if math.Abs(float64(array[startIdx])) > math.Abs(float64(array[endIdx])) {
			output[index] = array[startIdx] * array[startIdx]
			startIdx++
		} else {
			output[index] = array[endIdx] * array[endIdx]
			endIdx--
		}
		index--
	}
	return output
}
