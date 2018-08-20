/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var lst []*ListNode
	for true {
		if head != nil {
			lst = append(lst, head)
			head = head.Next
			continue
		}
		break
	}
	return lst[len(lst)/2]
}