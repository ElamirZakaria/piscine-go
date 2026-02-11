package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		if root.Right == nil {
			return BTreeIsBinary(root.Left)
		} else if root.Left == nil {
			return BTreeIsBinary(root.Right)
		} else if root.Right != nil && root.Left != nil {
			if !(root.Right.Data > root.Data && root.Left.Data < root.Data) {
				return false
			} else {
				return BTreeIsBinary(root.Right) && BTreeIsBinary(root.Left)
			}
		} else if root.Right != nil {
			if !(root.Data < root.Right.Data) {
				return false
			} else {
				return BTreeIsBinary(root.Right)
			}
		} else {
			if !(root.Data > root.Left.Data) {
				return false
			} else {
				return BTreeIsBinary(root.Left)
			}
		}
	}
}
