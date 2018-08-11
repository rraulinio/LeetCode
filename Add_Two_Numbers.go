func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	out := &ListNode{Val: -1, Next: nil}
	lst := false
	for true {
		var st int
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			st += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			st += l2.Val
			l2 = l2.Next
		}
		if lst == true {
			st++
			lst = false
		}
		if st > 9 {
			st = st % 10
			lst = true
		}
		d := &ListNode{Val: st, Next: nil}
		if out.Val == -1 {
			out = d
		} else {
			current := out
			for current.Next != nil {
				current = current.Next
			}
			current.Next = d
		}
		if l1 == nil && l2 == nil && lst == true {
			l1 = &ListNode{Val: 1, Next: nil}
			lst = false
		}
	}
	return out
}