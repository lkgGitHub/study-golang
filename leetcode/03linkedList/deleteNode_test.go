package linkedList

import (
	"fmt"
	"testing"
)

/*
删除链表中的节点

这道题感觉很扯，因为head和ListNode在官方代码中。线下手撸没法测试。
线下自己构造出head和ListNode 撸完又不对。哎~理解思想就好
*/
// 提交的代码
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

var head = &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}

// 实际逻辑和测试
func myDeleteNode(node *ListNode) {
	h := &ListNode{head.Val, head.Next}
	for ; h.Next != nil; h = h.Next {
		if h.Val == node.Val {
			next := h.Next
			h.Val = next.Val
			h.Next = next.Next
			break
		}
	}
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		node     *ListNode
		expected *ListNode
	}{
		{&ListNode{5, nil},
			&ListNode{4, &ListNode{1, &ListNode{9, nil}}}},

		{
			&ListNode{1, nil},
			&ListNode{4, &ListNode{5, &ListNode{9, nil}}},
		},
	}
	for _, tt := range tests {
		head = &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
		deleteNode(tt.node)
		isPass := true
		for ; head.Next != nil; head, tt.expected = head.Next, tt.expected.Next {
			if head.Val != tt.expected.Val {
				isPass = false
				break
			}
		}
		if !isPass {
			t.Error(fmt.Sprintf("TestRotate failed. node: %d, actual: %v", tt.node.Val, head))
		}
	}
}
