package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
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

// letcode problem
func (l *LinkedList) RemoveNthFromEnd(n int) {
	dummy := &Node{next: l.head}
	left, right := dummy, l.head

	for i := 0; i <= n; i++ {
		right = right.next
	}

	for right != nil {
		left = left.next
		right = right.next
	}

	left.next = left.next.next
	l.head = dummy.next
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

// the hanoi tower
func hanoiTower(height int, fromTower, midTower, toTower string) { //nolint
	if height >= 1 {
		hanoiTower(height-1, fromTower, midTower, toTower)
		fmt.Println(fromTower, toTower)
		hanoiTower(height-1, midTower, toTower, fromTower)
	}
}

// binary heap
type Heap struct {
	heap []int
	size int
}

func (h *Heap) Insert(i int) {
	h.heap = append(h.heap, i)
	h.size++
	index := h.size - 1
	for h.heap[parent(index)] < h.heap[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

// bubble sort
func sortColors(nums []int) []int { //nolint
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// selection sort
func selectionSort(nums []int) []int { //nolint
	for i := len(nums) - 1; i >= 0; i-- {
		currentMax := 0

		for j := 1; j < i+1; j++ {
			if nums[j] > nums[currentMax] {
				currentMax = j
			}
		}
		temp := nums[i]
		nums[i] = nums[currentMax]
		nums[currentMax] = temp
	}
	return nums
}

// quicksort
func quickSort(nums []int) []int { //nolint
	l := len(nums)
	if l < 2 {
		return nums
	}
	if l == 2 {
		if nums[0] <= nums[1] {
			return nums
		} else {
			return []int{nums[1], nums[0]}
		}
	}
	pivot := nums[0]
	var left, right []int
	for _, i := range nums[1:] {
		if i <= pivot {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

func longestSubarray(nums []int) int {
	var res []int
	var countZeros, countOnes int
	for stop := 0; stop < len(nums); stop++ {
		var maxCount, count, zeros int
		if nums[stop] == 0 {
			countZeros++
		} else {
			countOnes++
		}
		for start := stop; start < len(nums); start++ {
			if nums[start] == 1 {
				count++
			} else {
				maxCount += count
				zeros++
				count = 0
			}
			if zeros == 1 {
				if len(res) != 0 {
					if res[len(res)-1] < maxCount+count {
						res[len(res)-1] = maxCount + count
					}
				} else {
					res = append(res, maxCount+count)
				}
			}
			if zeros > 1 {
				if len(res) != 0 {
					if res[len(res)-1] < maxCount {
						res[len(res)-1] = maxCount
					}
				} else {
					res = append(res, maxCount+count)
				}
				break
			}
		}
	}
	if countOnes == len(nums) {
		return countOnes - 1
	}
	if countZeros == len(nums) {
		return 0
	}
	return res[len(res)-1]
}

func mergeSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	if l == 2 {
		if nums[0] <= nums[1] {
			return nums
		} else {
			return []int{nums[1], nums[0]}
		}
	}
	left, right := mergeSort(nums[:l/2]), mergeSort(nums[l/2:])
	r := len(left)
	if len(left) > len(right) {
		r = len(right)
	}
	var res []int
	for i := 0; i < r; i++ {
		if left[i] < right[i] {
			res = append(res, left[i])
			res = append(res, right[i])

		} else {
			res = append(res, right[i])
			res = append(res, left[i])
		}
	}
	return append(res, right[r:]...)
}

func firstBadVersion(n int, bad int) int {
	var low, mid int
	var result, hight int = n, n
	for low <= hight {
		mid = low + (hight-low)/2
		if mid == bad {
			result = mid
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

func maxSubArrayQuadric(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var result int
	for start := 0; start < len(nums); start++ {
		current := 0
		for stop := start; stop < len(nums); stop++ {
			current += nums[stop]
			if current > result {
				result = current
			}
		}
	}
	return result
}

func maxSubArrayLinear(nums []int) int {
	current, result := nums[0], nums[0]
	for start := 1; start < len(nums); start++ {
		current += nums[start]
		if current < nums[start] {
			current = nums[start]
		}
		if result < current {
			result = current
		}
	}
	return result
}

// simple binary tree on lists
type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func CreateTree(val int) *BinaryTree {
	return &BinaryTree{
		value: val,
	}
}

func (b *BinaryTree) Insert(val int) {
	if b.value <= val {
		if b.left == nil {
			b.left = &BinaryTree{value: val}
		} else {
			b.left.Insert(val)
		}
	} else {
		if b.right == nil {
			b.right = &BinaryTree{value: val}
		} else {
			b.right.Insert(val)
		}
	}
}

// direct tree walking
func PreOrder(child *BinaryTree, a *[]int) []int {
	if child != nil {
		*a = append(*a, child.value)
		if child.left != nil {
			PreOrder(child.left, a)
		}
		if child.right != nil {
			PreOrder(child.right, a)
		}
	}
	return *a
}

// symmetric tree walking
func SymmetricOrder(child *BinaryTree, a *[]int) []int {
	if child != nil {
		if child.left != nil {
			SymmetricOrder(child.left, a)
		}
		*a = append(*a, child.value)
		if child.right != nil {
			SymmetricOrder(child.right, a)
		}
	}
	return *a
}

func ReverseSymmetricOrder(child *BinaryTree, a *[]int) []int {
	if child != nil {
		if child.right != nil {
			ReverseSymmetricOrder(child.right, a)
		}
		*a = append(*a, child.value)
		if child.left != nil {
			ReverseSymmetricOrder(child.left, a)
		}
	}
	return *a
}

// riverse tree walking
func PostOrder(child *BinaryTree, a *[]int) []int {
	if child != nil {
		if child.left != nil {
			PostOrder(child.left, a)
		}
		if child.right != nil {
			PostOrder(child.right, a)
		}
		*a = append(*a, child.value)
	}
	return *a
}

func (b *BinaryTree) GetMinValue() int {
	return 0
}

func (b *BinaryTree) GetMaxValue() int {
	return 0
}

func GetTreeDepth(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}
	left := GetTreeDepth(tree.left)
	right := GetTreeDepth(tree.right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// from min to max
func (b *BinaryTree) InsertOrderedAscending(val int) {

}

// from max to min
func (b *BinaryTree) InsertOrderedDescending(val int) {

}

// Validate Binary Search Tree
func isValidBST(root *BinaryTree) bool { // nolint
	a := []int{}
	SymmetricOrder(root, &a)
	cur := a[0]
	for i := 1; i < len(a); i++ {
		if cur < a[i] {
			cur = a[i]
		} else {
			return false
		}
	}
	return true
}

func missingNumber(nums []int) int {
	nums = missingNumberSort(nums)
	var res int
	for i := 0; i <= len(nums); i++ {
		if i < len(nums) {
			if i != nums[i] {
				res = i
				break
			}
		} else {
			res = i
			break
		}
	}
	return res
}

func missingNumberSort(nums []int) []int {
	// selection sort
	for i := len(nums) - 1; i >= 0; i-- {
		currentMax := 0
		for j := 1; j < i+1; j++ {
			if nums[j] > nums[currentMax] {
				currentMax = j
			}
		}
		temp := nums[i]
		nums[i] = nums[currentMax]
		nums[currentMax] = temp
	}
	return nums
}

func rob(nums []int) int {
	var left, right int

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			left = max(left+nums[i], right)
		} else {
			right = max(right+nums[i], left)
		}
	}
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
