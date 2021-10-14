package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := LinkedList{}
	list.AddForward(9)
	list.AddForward(10)
	list.AddForward(11)
	list.AddForward(12)

	list.ReverseLinkedList()

	list.Display()
	fmt.Println("len", list.Len())
}

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
	m := make(map[byte]byte)
	for _, v := range []byte("abcdefghijklmnopqrstuvwxyz1234567890") {
		m[v] = v
	}
	ss := []byte(strings.ToLower(s))
	left, right := 0, len(ss)-1
	for left < right {
		for left < right {
			_, ex1 := m[ss[left]]
			if ex1 {
				break
			}
			left++
		}
		for right > left {
			_, ex2 := m[ss[right]]
			if ex2 {
				break
			}
			right--
		}
		if ss[left] != ss[right] {
			return false
		}
		left, right = left+1, right-1
	}
	return true
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

func countPrimes(n int) int {
	var count int
	if n <= 2 {
		return 0
	}
	arr := make([]bool, n)
	for idx := range arr {
		arr[idx] = true
	}
	for i := 2; i*i < n; i++ {
		if arr[i] {
			for j := i * i; j < n; j = j + i {
				arr[j] = false
			}
		}

	}
	for i := 2; i < n; i++ {
		if arr[i] {
			count++
		}
	}
	return count
}

// linked list
type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	length int
	head   *Node
}

func (l *LinkedList) Len() int {
	return l.length
}

func (l *LinkedList) Display() {
	for l.head != nil {
		fmt.Println(l.head.value, l.head)
		l.head = l.head.next
	}
}

func (l *LinkedList) AddBackward(i int) {
	n := &Node{value: i, next: l.head}
	l.head = n
	l.length++
}

func (l *LinkedList) AddForward(i int) {
	if l.length == 0 {
		l.head = &Node{value: i}
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &Node{value: i}
	}
	l.length++
}

func (l *LinkedList) RemoveNode(i int) {
	current := l.head
	prev := &Node{}
	for idx := 1; idx <= l.length; idx++ {
		if current.value == i && idx == 1 {
			l.head = current.next
			break
		} else if current.value == i {
			prev.next = current.next
			break
		}
		prev = current
		current = current.next
	}
	l.length--
}

func (l *LinkedList) ReverseLinkedList() {
	var prev *Node = nil
	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

// moving window find longest substring in string
func movingWindow(s string) int {
	b := []byte(s)
	n := len(b)
	right := 0 // window right border
	var window []byte
	maxLen := 0 // window len
	for right < n {
		if !contains(window, b[right]) {
			window = append(window, b[right])
		} else {
			if len(window) > maxLen {
				maxLen = len(window) // update window len if contains
			}
			for contains(window, b[right]) {
				window = window[1:]
			}
			window = append(window, b[right])
		}
		right++
	}
	var result int
	if maxLen > len(window) {
		result = maxLen
	} else {
		result = len(window)
	}
	return result
}

func contains(window []byte, i byte) bool {
	for _, j := range window {
		if j == i {
			return true
		}
	}
	return false
}

// stack
type Stack struct {
	stack  []int
	length int
}

func (s *Stack) Push(i int) {
	s.stack = append(s.stack, i)
	s.length++
}

func (s *Stack) Pop() int {
	last := s.stack[len(s.stack)-1:]
	s.stack = s.stack[:len(s.stack)-1]
	s.length--
	return last[0]
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// binary search
func binarySearchRecursive(nums []int, target int) bool {
	l := len(nums)
	if l == 1 && nums[0] == target {
		return true
	}
	if l == 0 {
		return false
	}
	mid := l / 2
	if target == nums[mid] {
		return true
	} else if len(nums) == 1 && nums[mid] != target {
		return false
	} else {
		if target > nums[mid] {
			return binarySearchRecursive(nums[mid+1:], target)
		} else {
			return binarySearchRecursive(nums[:mid], target)
		}
	}
}

// binary search non recursive
func binarySearch(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else {
			if nums[mid] < target {
				start++
			} else {
				end--
			}
		}
	}
	return -1
}

// queue
type Queue struct {
	queue  []int
	length int
}

func (q *Queue) Enqueue(i int) {
	q.queue = append(q.queue, i)
	q.length++
}

func (q *Queue) Dequeue() int {
	r := q.queue[q.length-1]
	q.queue = q.queue[:q.length-1]
	q.length--
	return r
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

// deque head is right side
type Deque struct {
	deque  []int
	length int
}

func (d *Deque) AddHead(i int) {
	d.deque = append(d.deque, i)
	d.length++
}

func (d *Deque) AddTail(i int) {
	d.deque = append([]int{i}, d.deque...)
	d.length++
}

func (d *Deque) GetHead() int {
	r := d.deque[d.length-1]
	d.deque = d.deque[:d.length-1]
	d.length--
	return r
}

func (d *Deque) GetTail() int {
	r := d.deque[0]
	d.deque = d.deque[1:]
	d.length--
	return r
}
