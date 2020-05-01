package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	myTree := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println("InOrder Traversal: ")
	inOrderTraversal(&myTree)

	fmt.Println("PreOrder Traversal: ")
	preOrderTraversal(&myTree)

	fmt.Println("PostOrder Traversal: ")
	postOrderTraversal(&myTree)

}

func inOrderTraversal(tn *TreeNode) {
	if tn != nil {
		inOrderTraversal(tn.Left)
		visit(tn)
		inOrderTraversal(tn.Right)
	}
}

func preOrderTraversal(tn *TreeNode) {
	if tn != nil {
		visit(tn)
		preOrderTraversal(tn.Left)
		preOrderTraversal(tn.Right)
	}
}

func postOrderTraversal(tn *TreeNode) {
	if tn != nil {
		postOrderTraversal(tn.Left)
		postOrderTraversal(tn.Right)
		visit(tn)
	}
}

func visit(tn *TreeNode) {
	fmt.Println(tn.Val)
}
