package main

func main() {

}

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	var output []int

	for len(stack) != 0 {
		n := len(stack)
		root := stack[n-1]
		stack = stack[:n-1]
		if root != nil {
			output = append(output, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
	return output
}
