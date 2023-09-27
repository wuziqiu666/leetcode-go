package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}
	head := &ListNode{Val: nodes[0]}
	curr := head
	for _, v := range nodes[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	strList := make([]string, 0)
	for l != nil {
		strList = append(strList, fmt.Sprintf("%d", l.Val))
		l = l.Next
	}
	str := "[" + strings.Join(strList, ", ") + "]"
	return str
}
