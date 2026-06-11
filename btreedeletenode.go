package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			return nil
		}
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
		return root
	}

	if node.Left == nil || node.Right == nil {
		var child *TreeNode
		if node.Left != nil {
			child = node.Left
		} else {
			child = node.Right
		}

		if node.Parent == nil {
			child.Parent = nil
			return child
		}

		child.Parent = node.Parent
		if node.Parent.Left == node {
			node.Parent.Left = child
		} else {
			node.Parent.Right = child
		}
		return root
	}

	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	node.Data = successor.Data

	return BTreeDeleteNode(root, successor)
}
