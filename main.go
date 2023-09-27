package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := new(ListNode)
	write := pre
	cur := &ListNode{Next: head}
	for {
		if cur.Next.Next == nil {
			if cur.Next.Val != cur.Val {
				write.Next = cur.Next
			} else {
				write.Next = nil
			}
			break
		}
		if cur.Next.Val != cur.Val && cur.Next.Val != cur.Next.Next.Val {
			write.Next = cur.Next
			write = write.Next
		}
		cur = cur.Next
	}
	return pre.Next
}
