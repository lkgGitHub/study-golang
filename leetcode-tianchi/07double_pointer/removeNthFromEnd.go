package main

import (
	"fmt"
)

/*
18. 删除链表的倒数第 N 个结点
难度：中等
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]

提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

进阶：你能尝试使用一趟扫描实现吗？
*/

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	head := &ListNode{1, &ListNode{2, nil}}
	head = removeNthFromEnd(head, 1)

	for ; head != nil; head = head.Next {
		fmt.Printf("%d ", head.Val)
	}
}

// 同速指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}

	var node2 *ListNode

	p1 := 0
	node1 := &ListNode{head.Val, head.Next}
	for {
		p1++
		node1 = node1.Next

		if node1 == nil {
			break
		}

		if node1.Next == nil {
			if p1-n+1 == 0 {
				node2 = &ListNode{head.Val, head.Next}
			}

			if node2 == nil {

			}

			next := node2.Next
			if next != nil {
				node2.Next = next.Next
				break
			} else {
				node1 = nil
			}
		}

		if p1-n+1 == 0 {
			node2 = &ListNode{head.Val, head.Next}
		} else if p1-n+1 > 0 {
			node2 = node2.Next
		}

	}

	return head
}
