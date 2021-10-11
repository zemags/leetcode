package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		actual   []int
		expected int
	}{
		{actual: []int{}, expected: 0},
		{actual: []int{1, 1, 2}, expected: 2},
	}

	for _, tt := range tests {
		t.Run("remove duplicates", func(t *testing.T) {
			actual := removeDuplicates(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		actual   []int
		expected int
	}{
		{actual: []int{}, expected: 0},
		{actual: []int{7, 1, 5, 3, 6, 4}, expected: 7},
		{actual: []int{1, 2, 3, 4, 5}, expected: 4},
		{actual: []int{7, 6, 4, 3, 1}, expected: 0},
	}

	for _, tt := range tests {
		t.Run("max profit", func(t *testing.T) {
			actual := maxProfit(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		actual   []int
		k        int
		expected []int
	}{
		{actual: []int{}, k: 3, expected: []int{}},
		{actual: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{5, 6, 7, 1, 2, 3, 4}},
		{actual: []int{-1, -100, 3, 99}, k: 2, expected: []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {
		t.Run("rotate", func(t *testing.T) {
			actual := rotate(tt.actual, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		actual   []int
		expected bool
	}{
		{actual: []int{}, expected: false},
		{actual: []int{1}, expected: false},
		{actual: []int{1, 2, 3, 1}, expected: true},
		{actual: []int{1, 2, 3, 4}, expected: false},
		{actual: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expected: true},
	}

	for _, tt := range tests {
		t.Run("contains duplicate", func(t *testing.T) {
			actual := containsDuplicate(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		actual   []int
		expected int
	}{
		{actual: []int{1}, expected: 1},
		{actual: []int{4, 1, 2, 1, 2}, expected: 4},
		{actual: []int{2, 2, 1}, expected: 1},
	}

	for _, tt := range tests {
		t.Run("single number", func(t *testing.T) {
			actual := singleNumber(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		actual1  []int
		actual2  []int
		expected []int
	}{
		{actual1: []int{1}, actual2: []int{1}, expected: []int{1}},
		{actual1: []int{1}, actual2: []int{2}, expected: []int{}},
		{actual1: []int{1, 2, 2, 1}, actual2: []int{2, 2}, expected: []int{2, 2}},
		{actual1: []int{4, 9, 5}, actual2: []int{9, 4, 9, 8, 4}, expected: []int{4, 9}},
		{actual1: []int{43, 85, 49, 2, 83, 2, 39, 99, 15, 70, 39, 27, 71, 3, 88, 5, 19, 5, 68, 34, 7, 41, 84, 2, 13, 85, 12, 54, 7, 9, 13, 19, 92}, actual2: []int{10, 8, 53, 63, 58, 83, 26, 10, 58, 3, 61, 56, 55, 38, 81, 29, 69, 55, 86, 23, 91, 44, 9, 98, 41, 48, 41, 16, 42, 72, 6, 4, 2, 81, 42, 84, 4, 13}, expected: []int{2, 83, 3, 41, 84, 13, 9}},
		{actual1: []int{84, 5, 30, 84, 67, 78, 73, 38, 93, 92, 15, 43, 38, 81, 68, 65, 62, 21, 16, 38, 95, 68, 60, 35, 43, 95, 67}, actual2: []int{82, 60, 70, 10, 94, 6, 44, 51, 1, 3, 97, 84, 3, 87, 91, 55, 81, 90, 45, 22, 18, 58, 62, 96, 27, 24, 16, 63, 30, 60, 29, 93, 27, 56, 79, 4, 69, 9, 21, 23, 7, 49, 62, 89, 22, 64, 85, 75, 55, 49, 57, 17, 84, 49, 8, 13, 94, 40, 75, 50, 93, 46, 36, 94, 50, 0, 3, 65, 49, 82, 45, 11, 53, 63, 27, 71, 45, 37, 45, 19, 21, 57, 66, 99, 94, 92, 44, 35, 84, 78, 80, 88}, expected: []int{84, 30, 84, 78, 93, 92, 81, 65, 62, 21, 16, 60, 35}},
	}

	for _, tt := range tests {
		t.Run("intersect", func(t *testing.T) {
			actual := intersect(tt.actual1, tt.actual2)
			assert.ElementsMatch(t, tt.expected, actual)
		})
	}
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		actual   []int
		expected []int
	}{
		{actual: []int{1}, expected: []int{2}},
		{actual: []int{0}, expected: []int{1}},
		{actual: []int{9}, expected: []int{1, 0}},
		{actual: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{actual: []int{1, 2, 9}, expected: []int{1, 3, 0}},
		{actual: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
		{actual: []int{9, 9}, expected: []int{1, 0, 0}},
		{actual: []int{9, 9, 9}, expected: []int{1, 0, 0, 0}},
		{actual: []int{8, 9}, expected: []int{9, 0}},
		{actual: []int{9, 8, 9}, expected: []int{9, 9, 0}},
	}

	for _, tt := range tests {
		t.Run("plusone", func(t *testing.T) {
			actual := plusOne(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		actual   []int
		expected []int
	}{
		{actual: []int{0}, expected: []int{0}},
		{actual: []int{0, 0}, expected: []int{0, 0}},
		{actual: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
	}

	for _, tt := range tests {
		t.Run("move zeros", func(t *testing.T) {
			actual := moveZeroes(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		actual   []int
		target   int
		expected []int
	}{
		{actual: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{actual: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{actual: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run("two sum", func(t *testing.T) {
			actual := twoSum(tt.actual, tt.target)
			assert.ElementsMatch(t, tt.expected, actual)
		})
	}
}

func TestReverseString(t *testing.T) {
	expected := []byte("abcd")
	actual := reverseString([]byte("dcba"))
	assert.Equal(t, expected, actual)

	expected = []byte("olleh")
	actual = reverseString([]byte("hello"))
	assert.Equal(t, expected, actual)

	expected = []byte("hannah")
	actual = reverseString([]byte("hannah"))
	assert.Equal(t, expected, actual)
}

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	tt := "nagaram"
	actual := isAnagram(s, tt)
	assert.Equal(t, true, actual)

	s = "rat"
	tt = "car"
	actual = isAnagram(s, tt)
	assert.Equal(t, false, actual)
}

func TestFibonacciRecursive(t *testing.T) {
	actual := fibonacciRecursive(4)
	assert.Equal(t, 3, actual)

	actual = fibonacciRecursive(1)
	assert.Equal(t, 1, actual)

	actual = fibonacciRecursive(5)
	assert.Equal(t, 5, actual)
}

func TestFibonacciLoop(t *testing.T) {
	actual := fibonacciLoop(4)
	assert.Equal(t, 3, actual)

	actual = fibonacciLoop(1)
	assert.Equal(t, 1, actual)

	actual = fibonacciLoop(5)
	assert.Equal(t, 5, actual)
}

func TestFibonacciWhile(t *testing.T) {
	actual := fibonacciWhile(4)
	assert.Equal(t, 3, actual)

	actual = fibonacciWhile(1)
	assert.Equal(t, 1, actual)

	actual = fibonacciWhile(5)
	assert.Equal(t, 5, actual)
}
