/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nds []*ListNode
	current := head
	nds = append(nds, current)
	for current.Next != nil {
		current = current.Next
		nds = append(nds, current)
	}
	if len(nds) == 1 {
		return nil
	} else if len(nds) == n {
		nds[0] = nil
		head = nds[1]
	} else if len(nds)-n+1 < len(nds) {
		nds[len(nds)-n-1].Next = nds[len(nds)-n+1]
	} else {
		nds[len(nds)-n-1].Next = nil
	}
	return head
}