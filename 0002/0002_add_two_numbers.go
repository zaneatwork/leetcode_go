// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order, and each of their nodes
// contains a single digit. Add the two numbers and return the sum as a linked
// list.
//
// Constraints:
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have
// leading zeros.

// You may assume the two numbers do not contain any leading zero, except the
// number 0 itself.
package main

import "fmt"


func main() {
	vals1 := []int{2, 4, 3}
	head1 := listNodeFromIntSlice(vals1)
  printListNode(&head1)

  fmt.Printf("+\n")

	vals2 := []int{5, 6, 4}
	head2 := listNodeFromIntSlice(vals2)
  printListNode(&head2)
  fmt.Printf("=\n")

  newHead1 := addTwoNumbers(&head1, &head2)
  printListNode(newHead1)

  fmt.Println()

	vals3 := []int{9,9,9,9,9,9,9}
	head3 := listNodeFromIntSlice(vals3)
  printListNode(&head3)

  fmt.Printf("+\n")

	vals4 := []int{9,9,9,9}
	head4 := listNodeFromIntSlice(vals4)
  printListNode(&head4)
  fmt.Printf("=\n")

  newHead2 := addTwoNumbers(&head3, &head4)
  printListNode(newHead2)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Take an int slice and return the head of a linked list with the values of 
// that slice.
func listNodeFromIntSlice(vals []int) ListNode {
	head := ListNode{Val: vals[0], Next: nil}
	tail := &head
	for i := 1; i < len(vals); i++ {
		newNode := ListNode{Val: vals[i]}
		tail.Next = &newNode
		tail = &newNode
	}
	return head
}

func printListNode(ln *ListNode) {
  tail := ln
  for tail != nil {
    fmt.Printf("(%v)", tail.Val)
    tail = tail.Next
    if tail != nil {
      fmt.Printf("->")
    } else {
      fmt.Printf("\n")
    }
  }
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  sum := l1.Val + l2.Val

  if l1.Next == nil && l2.Next == nil && sum < 10 {
    sumNode := ListNode{Val:l1.Val+l2.Val}
    return &sumNode
  } 

  left := l1.Next
  right := l2.Next

  if left == nil {
    blank := ListNode{Val:0} 
    left = &blank 
  } 

  if right == nil {
    blank := ListNode{Val:0} 
    right = &blank
  }

  if sum >= 10 {
    sum = sum % 10
    left.Val += 1
  }

  sumNode := ListNode{Val:sum, Next:addTwoNumbers(left, right)}
  return &sumNode
}
