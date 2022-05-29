// Daily practice
// 234. Palindrome Linked List
// https://leetcode.com/problems/palindrome-linked-list/
package leetcode

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// newList creates a new list from passed values and
// returns pointer to its head.
func newList(vals []int) *ListNode {
	var list *ListNode
	if len(vals) > 0 {
		list = &ListNode{vals[0], nil}
	}
	var cur *ListNode
	prev := list
	for i := 1; i < len(vals); i++ {
		cur = &ListNode{vals[i], nil}
		prev.Next = cur
		prev = cur
	}
	return list
}

// isPalindrome returns true if the list is a palindrome:
// 1 -> 2 -> 2 -> 1: true
// 1 -> 2: false
// 05/29/2022 08:00 ✅
// Runtime: 150 ms, faster than 84.59% of Go online submissions for Palindrome Linked List.
// Memory Usage: 11.2 MB, less than 36.59% of Go online submissions for Palindrome Linked List.
func isPalindrome(list *ListNode) bool {
	length := 0
	cur := list
	for cur != nil {
		length++
		cur = cur.Next
	}
	halfIndex := length / 2
	if length%2 != 0 {
		halfIndex++
	}
	// search for head of the second part of linked list
	cur = list
	for i := 0; i < halfIndex; i++ {
		cur = cur.Next
	}
	// reverse second part of linked list
	half := reverse(cur)
	cur = list
	for i := 0; i < length/2; i++ {
		if cur.Val != half.Val {
			return false
		}
		cur = cur.Next
		half = half.Next
	}
	return true
}

// reverse reverses single linked list and returns a new head
func reverse(list *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	cur := list
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		vals         []int
		isPalindrome bool
	}{
		{[]int{1, 2}, false},
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
	}
	for _, data := range testData {
		if actual := isPalindrome(newList(data.vals)); actual != data.isPalindrome {
			t.Errorf("isPalindrome1(%v): want %v, but was %v", data.vals, data.isPalindrome, actual)
		}
	}
}

func TestReverse(t *testing.T) {
	list := newList([]int{5, 4, 3, 2, 1})
	list = reverse(list)
	for i := 1; i <= 5; i++ {
		if list.Val != i {
			t.Errorf("Val: want %d, but was %d", i, list.Val)
		}
		list = list.Next
	}
}

func TestNewList(t *testing.T) {
	vals := []int{5, 4, 3, 2, 1}
	list := newList(vals)
	for _, val := range vals {
		if val != list.Val {
			t.Error("Val: want", val, "but was", list.Val)
		}
		list = list.Next
	}
}
