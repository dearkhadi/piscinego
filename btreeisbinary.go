package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return checkBST(root, "", "")
}

func checkBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	if min != "" && node.Data <= min {
		return false
	}

	if max != "" && node.Data >= max {
		return false
	}

	return checkBST(node.Left, min, node.Data) && checkBST(node.Right, node.Data, max)
}
