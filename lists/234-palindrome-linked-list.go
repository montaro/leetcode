package main

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slice := convertListToSlice(head)
	front := 0
	back := len(slice) - 1
	for front < back {
		if slice[front] != slice[back] {
			return false
		}
		front++
		back--
	}
	return true
}

func convertListToSlice(head *ListNode) []int {
	var slice []int
	for true {
		slice = append(slice, head.Val)
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
	return slice
}

func main() {
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(isPalindrome(&list))
}
