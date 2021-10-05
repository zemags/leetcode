package main

import "fmt"

func main() { fmt.Println(intersect([]int{1}, []int{2})) }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	leftIndex := 1
	for rightIndex := 1; rightIndex < len(nums); rightIndex++ {
		if nums[rightIndex] != nums[rightIndex-1] {
			nums[leftIndex] = nums[rightIndex]
			leftIndex++
		}
	}
	return leftIndex
}

func maxProfit(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	profit := 0
	for i := 0; i != len(nums)-1; i++ {
		currentValue := nums[i]
		nextValue := nums[i+1]

		currentProfit := 0

		if currentValue < nextValue {
			tempProfit := (nextValue - currentValue)
			if tempProfit > currentProfit {
				currentProfit = tempProfit
			}
		}
		profit += currentProfit
	}
	return profit
}

func rotate(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	k = k % len(nums)
	firstTemp, secondTemp := nums[len(nums)-k:], nums[:len(nums)-k]
	return append(firstTemp, secondTemp...)
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	for idx, i := range nums {
		for jdx, j := range nums {
			if idx != jdx {
				if i == j {
					return true
				}
			}

		}
	}
	return false
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	temp := nums[0]
	for i := 1; i != len(nums); i++ {
		temp = temp ^ nums[i]
	}
	return temp
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	maps1 := make(map[int]int)
	nums1Max := 0
	for _, i := range nums1 {
		if i > nums1Max {
			nums1Max = i
		}
		maps1[i]++
	}
	maps2 := make(map[int]int)
	nums2Max := 0
	for _, i := range nums2 {
		if i > nums2Max {
			nums2Max = i
		}
		maps2[i]++
	}
	max := nums2Max
	if nums1Max > nums2Max {
		max = nums1Max
	}

	var result []int
	for i := 0; i < max+1; i++ {
		comp1, exist1 := maps1[i]
		comp2, exist2 := maps2[i]
		if exist1 && exist2 {
			rng := comp1
			if comp1 > comp2 {
				rng = comp2
			}

			for j := 0; j < rng; j++ {
				result = append(result, i)
			}
		}
	}
	return result
}
