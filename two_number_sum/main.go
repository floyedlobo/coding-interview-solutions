package main

import (
	"sort"
)

//https://leetcode.com/problems/two-sum/

//Time: O(n) & Space: O(n)
func TwoNumberSum(array []int, target int) []int {

	complementList := make(map[int]bool, len(array))

	for _, num := range array {
		compNum := target - num
		if complementList[num] == true {
			return []int{num, compNum}
		}
		complementList[compNum] = true
	}

	return []int{}
}

//Time O(nlogn) Space O(1)

func TwoNumberSum2(array []int, target int) []int {
	startIdx := 0
	endIdx := len(array) - 1
	sort.Ints(array)
	for startIdx < endIdx {
		sum := array[startIdx] + array[endIdx]
		if sum > target {
			endIdx--
		} else if sum < target {
			startIdx++
		} else {
			return []int{array[startIdx], array[endIdx]}
		}
	}
	return []int{}
}
