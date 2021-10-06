package main

import "fmt"

func main() { fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12})) }

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

func plusOne(nums []int) []int {
	length := len(nums)
	if length == 1 && nums[0] == 9 {
		return []int{1, 0}
	}
	if nums[length-1] == 9 {
		increment := 0
		for i := length - 1; i >= 0; i-- {
			if nums[i] == 9 {
				nums[i] = 0
				increment = 1
			} else {
				nums[i] += increment
				increment = 0
				break
			}
		}
	} else {
		nums[length-1]++
	}
	if nums[0] == 0 {
		return append([]int{1}, nums...)
	}
	return nums
}

func moveZeroes(nums []int) []int {
	var zeros []int
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < len(nums) {
			if nums[i] == 0 {
				zeros = append(zeros, 0)
				nums = append(nums[:i], nums[i+1:]...)
			}
		}
	}
	return append(nums, zeros...)
}

func twoSum(nums []int, target int) []int {
	var r []int
	for i := 0; i != len(nums); i++ {

		start := nums[i]
		end := nums[len(nums)-(i+1)]

		fmt.Println(start, end)

		if start+end == target {
			r = []int{i, len(nums) - (i + 1)}
		} else if start+nums[i+1] == target {
			r = []int{i, i + 1}
		} else if end+nums[len(nums)-(i+1)] == target {
			r = []int{len(nums) - (i + 1), len(nums) - (i + 1)}
		}
	}
	return r
}
