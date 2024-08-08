package problems

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePaths(root *TreeNode) {

	if root == nil {
		return
	}

	fmt.Println(root.Val)

	BinaryTreePaths(root.Left)
	BinaryTreePaths(root.Right)

}
