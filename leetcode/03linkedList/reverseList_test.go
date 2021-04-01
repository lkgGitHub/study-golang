package linkedList

import (
	"fmt"
	"testing"
)

/*
反转链表
反转一个单链表。

示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

相关标签:	链表
*/
func reverseList(head *ListNode) *ListNode {

}

func TestReverseList(t *testing.T) {
	tests := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			&ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
	}

	for _, tt := range tests {
		actual := reverseList(tt.head)
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
