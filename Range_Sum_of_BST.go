// Runtime: 200 ms, faster than 100.00% of Go online submissions for Range Sum of BST.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, L, R int, sum *int) {
	if root == nil {
		return
	} else if root.Val >= L && root.Val <= R {
		*sum += root.Val
	}
	helper(root.Left, L, R, sum)
	helper(root.Right, L, R, sum)
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int
	helper(root, L, R, &sum)
	return sum
}