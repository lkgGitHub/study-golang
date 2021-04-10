package linkedList

import (
	"fmt"
	"testing"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headSize := 0
	for node := head; node.Next != nil; node = node.Next {
		headSize++
	}
	deleteIndex := headSize + 1 - n

	node := head
	for ; deleteIndex >= 0; deleteIndex-- {
		if deleteIndex == 0 {
			if node.Next == nil {
				node = nil
				if headSize == 0 {
					head = &ListNode{}
				}
				break
			} else {
				node.Val = node.Next.Val
				node.Next = node.Next.Next
				break
			}
		}
		node = node.Next
	}

	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head     *ListNode
		n        int
		expected *ListNode
	}{
		//{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		//	2,
		//	&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		//},
		//{
		//	&ListNode{1, nil},
		//	1,
		//	&ListNode{},
		//},
		{
			&ListNode{1, &ListNode{2, nil}},
			1,
			&ListNode{1, nil},
		},
	}

	for _, tt := range tests {
		actual := removeNthFromEnd(tt.head, tt.n)
		isPass := true
		for ; actual.Next != nil; actual, tt.expected = actual.Next, tt.expected.Next {
			if actual.Val != tt.expected.Val {
				isPass = false
				break
			}
		}
		if !isPass {
			t.Error(fmt.Sprintf("TestRotate failed. node: %v, actual: %v", tt.head, actual))
		}
	}
}
