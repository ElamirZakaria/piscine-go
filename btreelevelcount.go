package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	RightCount := BTreeLevelCount(root.Right)
	LeftCount := BTreeLevelCount(root.Left)
	if RightCount > LeftCount {
		return 1 + RightCount
	} else {
		return 1 + LeftCount
	}
}
