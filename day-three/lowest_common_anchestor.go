package main

import (
	"fmt"
)

// TreeNode definition is comment out cause already this declared in this package  but this is must need to make this program workable
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LCA function
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

// Main function
func main() {
	// Construct the tree manually
	n3 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 5}
	n1 := &TreeNode{Val: 1}
	n6 := &TreeNode{Val: 6}
	n2 := &TreeNode{Val: 2}
	n0 := &TreeNode{Val: 0}
	n8 := &TreeNode{Val: 8}
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}
	n9 := &TreeNode{Val: 9}
	n11 := &TreeNode{Val: 11}
	n10 := &TreeNode{Val: 10}
	n15 := &TreeNode{Val: 15}
	n16 := &TreeNode{Val: 16}

	n3.Left, n3.Right = n5, n1
	n5.Left, n5.Right = n6, n2
	n1.Left, n1.Right = n0, n8
	n2.Left, n2.Right = n7, n4
	n0.Left, n0.Right = n9, n15
	n4.Left, n4.Right = n11, n10
	n15.Left, n15.Right = n16, nil

	// Example 1: p = 6, q = 4
	lca := lowestCommonAncestor(n3, n6, n4)
	fmt.Printf("LCA of %d and %d is: %d\n", n6.Val, n4.Val, lca.Val)

	// Example 2: p = 5, q = 1
	lca2 := lowestCommonAncestor(n3, n9, n16)
	fmt.Printf("LCA of %d and %d is: %d\n", n9.Val, n16.Val, lca2.Val)

	// Example 3: p = 5, q = 4
	lca3 := lowestCommonAncestor(n3, n10, n7)
	fmt.Printf("LCA of %d and %d is: %d\n", n10.Val, n7.Val, lca3.Val)
}
