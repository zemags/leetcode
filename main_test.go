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

func TestIsPalindrome(t *testing.T) {
	actual := isPalindrome("A man, a plan, a canal: Panama")
	assert.Equal(t, true, actual)
	actual = isPalindrome("race a car")
	assert.Equal(t, false, actual)
}

func TestFizzBuzz(t *testing.T) {
	actual := fizzBuzz(3)
	assert.Equal(t, []string{"1", "2", "Fizz"}, actual)

	actual = fizzBuzz(5)
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz"}, actual)

	actual = fizzBuzz(15)
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, actual)
}

func TestCcountPrimes(t *testing.T) {
	actual := countPrimes(10)
	assert.Equal(t, 4, actual)

	actual = countPrimes(0)
	assert.Equal(t, 0, actual)

	actual = countPrimes(1)
	assert.Equal(t, 0, actual)

	actual = countPrimes(30)
	assert.Equal(t, 10, actual)
}

func TestMmovingWindow(t *testing.T) {
	tests := []struct {
		actual   string
		expected int
	}{
		{actual: "abcabcbb", expected: 3},
		{actual: "bbbbb", expected: 1},
		{actual: "pwwkew", expected: 3},
		{actual: "adeafgdc", expected: 6},
	}

	for _, tt := range tests {
		t.Run("moving window", func(t *testing.T) {
			actual := movingWindow(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestStack(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	assert.Equal(t, []int{1, 2}, s.stack)
	s.Push(3)
	s.Push(4)
	actual := s.Pop()
	assert.Equal(t, 4, actual)
}

func TestBinarySearchRecursive(t *testing.T) {
	actual := binarySearchRecursive([]int{-1, 0, 3, 5, 9, 12}, 9)
	assert.Equal(t, true, actual)
	actual = binarySearchRecursive([]int{-1, 0, 3, 5, 9, 12}, 2)
	assert.Equal(t, false, actual)
}

func TestBbinarySearch(t *testing.T) {
	actual := binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9)
	assert.Equal(t, 4, actual)
	actual = binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2)
	assert.Equal(t, -1, actual)
}

func TestQueue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, []int{1, 2, 3}, q.queue)

	actual := q.Dequeue()
	assert.Equal(t, 3, actual)
	assert.Equal(t, []int{1, 2}, q.queue)

	q.Dequeue()
	q.Dequeue()
	actualBool := q.IsEmpty()
	assert.Equal(t, true, actualBool)
}

func TestDeque(t *testing.T) {
	d := &Deque{}

	d.AddHead(1)
	d.AddHead(2)
	d.AddHead(3)
	actual := d.GetHead()
	assert.Equal(t, 3, actual)
	assert.Equal(t, 2, d.length)

	d.AddTail(4)
	d.AddTail(5)
	actual = d.GetTail()
	assert.Equal(t, 5, actual)
	assert.Equal(t, 3, d.length)
}

func TestInsertToHeap(t *testing.T) {
	h := &Heap{}
	for _, i := range []int{1, 2, 3, 5, 2, 4} {
		h.Insert(i)
	}
	assert.Equal(t, []int{5, 3, 4, 1, 2, 2}, h.heap)
}

func TestSortColors(t *testing.T) {
	actual := sortColors([]int{5, 3, 2, 6, 10})
	assert.Equal(t, []int{2, 3, 5, 6, 10}, actual)
}

func TestSelectionSort(t *testing.T) {
	actual := selectionSort([]int{5, 3, 2, 6, 10})
	assert.Equal(t, []int{2, 3, 5, 6, 10}, actual)

	actual = selectionSort([]int{5, 4, 3, 2})
	assert.Equal(t, []int{2, 3, 4, 5}, actual)
}

func TestQuickSort(t *testing.T) {
	actual := quickSort([]int{6})
	assert.Equal(t, []int{6}, actual)

	actual = quickSort([]int{6, 10})
	assert.Equal(t, []int{6, 10}, actual)

	actual = quickSort([]int{16, 10})
	assert.Equal(t, []int{10, 16}, actual)

	actual = quickSort([]int{5, 3, 2, 6, 10})
	assert.Equal(t, []int{2, 3, 5, 6, 10}, actual)
}

func TestLlongestSubarray(t *testing.T) {
	tests := []struct {
		actual   []int
		expected int
	}{
		{actual: []int{1, 1, 0, 1}, expected: 3},
		{actual: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, expected: 5},
		{actual: []int{1, 1, 1}, expected: 2},
		{actual: []int{0, 0, 0}, expected: 0},
	}

	for _, tt := range tests {
		t.Run("longets subarray", func(t *testing.T) {
			actual := longestSubarray(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMergeSort(t *testing.T) {
	actual := mergeSort([]int{6})
	assert.Equal(t, []int{6}, actual)

	actual = mergeSort([]int{6, 10})
	assert.Equal(t, []int{6, 10}, actual)

	actual = mergeSort([]int{16, 10})
	assert.Equal(t, []int{10, 16}, actual)

	actual = mergeSort([]int{3, 2, 4, 3})
	assert.Equal(t, []int{2, 3, 3, 4}, actual)

	actual = mergeSort([]int{5, 3, 2, 6, 10})
	assert.Equal(t, []int{2, 3, 5, 6, 10}, actual)
}

func TestFirstBadVersion(t *testing.T) {
	actual := firstBadVersion(5, 4)
	assert.Equal(t, 4, actual)
	actual = firstBadVersion(1, 1)
	assert.Equal(t, 1, actual)
}

func TestMmaxSubArrayQuadric(t *testing.T) {
	tests := []struct {
		actual   []int
		expected int
	}{
		{actual: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
		{actual: []int{1}, expected: 1},
		{actual: []int{5, 4, -1, 7, 8}, expected: 23},
	}

	for _, tt := range tests {
		t.Run("max subarray", func(t *testing.T) {
			actual := maxSubArrayQuadric(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMmaxSubArrayLinear(t *testing.T) {
	tests := []struct {
		actual   []int
		expected int
	}{
		{actual: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
		{actual: []int{1}, expected: 1},
		{actual: []int{5, 4, -1, 7, 8}, expected: 23},
	}

	for _, tt := range tests {
		t.Run("max subarray", func(t *testing.T) {
			actual := maxSubArrayLinear(tt.actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSimpleBinaryTree(t *testing.T) {
	b := CreateTree(3)
	b.Insert(5)
	b.Insert(4)
	b.Insert(2)
	b.Insert(6)
	b.Insert(3)

	actualArray := []int{}
	actual := PreOrder(b, &actualArray)
	assert.Equal(t, []int{3, 5, 6, 4, 3, 2}, actual)

	actualArray = []int{}
	actual = SymmetricOrder(b, &actualArray)
	assert.Equal(t, []int{6, 5, 4, 3, 3, 2}, actual)

	actualArray = []int{}
	actual = ReverseSymmetricOrder(b, &actualArray)
	assert.Equal(t, []int{2, 3, 3, 4, 5, 6}, actual)

	actualArray = []int{}
	actual = PostOrder(b, &actualArray)
	assert.Equal(t, []int{6, 3, 4, 5, 2, 3}, actual)

	actualDepth := GetTreeDepth(b)
	assert.Equal(t, 4, actualDepth)
}

func TestOrderedTree(t *testing.T) {
	b := CreateTree(3)
	b.Insert(5)
	b.Insert(4)
	b.Insert(2)
	b.Insert(6)
	b.Insert(3)
}

func TestMmissingNumber(t *testing.T) {
	actual := missingNumber([]int{3, 0, 1})
	assert.Equal(t, 2, actual)

	actual = missingNumber([]int{0, 1})
	assert.Equal(t, 2, actual)

	actual = missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	assert.Equal(t, 8, actual)

	actual = missingNumber([]int{0})
	assert.Equal(t, 1, actual)

	actual = missingNumber([]int{1, 2})
	assert.Equal(t, 0, actual)
}

func TestRemoveNthFromEnd(t *testing.T) {
	l := &LinkedList{}
	for i := 1; i < 6; i++ {
		l.AddForward(i)
	}
	l.RemoveNthFromEnd(2)
}

func TestRob(t *testing.T) {
	actual := rob([]int{1, 2, 3, 1})
	assert.Equal(t, 4, actual)
	actual = rob([]int{2, 7, 9, 3, 1})
	assert.Equal(t, 12, actual)
	actual = rob([]int{1, 1})
	assert.Equal(t, 1, actual)
	actual = rob([]int{1, 1, 2})
	assert.Equal(t, 3, actual)
	actual = rob([]int{1})
	assert.Equal(t, 1, actual)
	actual = rob([]int{1, 3, 1})
	assert.Equal(t, 3, actual)
	actual = rob([]int{1, 2, 1, 1})
	assert.Equal(t, 3, actual)
	actual = rob([]int{1, 3, 4, 1})
	assert.Equal(t, 5, actual)
	actual = rob([]int{2, 3, 4, 6})
	assert.Equal(t, 9, actual)
}
