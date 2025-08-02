package main

import "fmt"

func iterativeInorderTreeTraversal(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		// Step 1: Go as far left as possible and push all nodes to stack
		for curr != nil {
			stack = append(stack, curr) // push current node
			curr = curr.Left            // move left
		}

		// Step 2: Now pop the top node from stack (leftmost)
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Step 3: Visit the current node
		result = append(result, curr.Val)

		// Step 4: Move to the right subtree
		curr = curr.Right
	}
	return result
}

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	res := iterativeInorderTreeTraversal(root)
	fmt.Println(res)
}
