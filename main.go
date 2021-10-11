package main

import (
	"fmt"
	"strconv"
)

func main() { fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12})) }

// arrays
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
	var x, y int
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		res := target - nums[i]
		_, ex := m[res]
		if ex {
			x, y = i, m[res]
		}
		m[nums[i]] = i
	}
	return []int{x, y}
}

// strings
func reverseString(s []byte) []byte {
	l := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[l-i], s[i] = s[i], s[l-i]
	}
	return s
}

func isAnagram(s string, t string) bool {
	f := func() map[byte]int {
		m := make(map[byte]int)
		s := []byte("abcdefghijklmnopqrstuvwxyz")
		for _, v := range s {
			m[v] = 0
		}
		return m
	}
	sm, tm := f(), f()
	for _, v := range []byte(s) {
		sm[v]++

	}
	for _, v := range []byte(t) {
		tm[v]++
	}
	for k := range f() {
		if sm[k] != tm[k] {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	// m := make(map[byte]byte)
	// for _, v := range []byte("abcdefghijklmnopqrstuvwxyz") {
	// 	m[v] = v
	// }
	// ss := []byte(strings.ToLower(s))
	// for i := 0; i < len(ss)-1; i++ {

	// }

	return false
}

// other
func fibonacciRecursive(x int) int {
	if x == 1 || x == 2 {
		return 1
	}
	return fibonacciRecursive(x-1) + fibonacciRecursive(x-2)
}

func fibonacciLoop(x int) int {
	f1, f2 := 1, 1
	for i := 2; i < x; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

func fibonacciWhile(x int) int {
	f1, f2 := 1, 1
	x = x - 2
	for x > 0 {
		f1, f2 = f2, f1+f2
		x--
	}
	return f2
}

// math
func fizzBuzz(n int) []string {
	var r []string
	for i := 1; i <= n; i++ {
		fmt.Println(i, i%3, i%5)
		if i%3 == 0 && i%5 == 0 {
			r = append(r, "FizzBuzz")
		} else if i%3 == 0 {
			r = append(r, "Fizz")
		} else if i%5 == 0 {
			r = append(r, "Buzz")
		} else {
			r = append(r, strconv.Itoa(i))
		}
	}
	return r
}
