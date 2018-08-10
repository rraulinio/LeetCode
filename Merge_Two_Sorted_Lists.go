func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	c := &ListNode{}
	var lst []int
	for true {
		if l1 != nil {
			lst = append(lst, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			lst = append(lst, l2.Val)
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			break
		}
	}
	sort.Ints(lst)
	for _, x := range lst {
		d := &ListNode{Val: x, Next: nil}
		if c.Val == 0 {
			c = d
		} else {
			current := c
			for current.Next != nil {
				current = current.Next
			}
			current.Next = d
		}
	}
	return c
}