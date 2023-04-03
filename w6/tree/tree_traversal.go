// Vish Siriwatana
// 64050229

package main

import "fmt"

func main() {
	n := 10
	node := []byte(" ABCDEFGHIJ") // keep binary tree in 1D array
	fmt.Print("DFS in-order   : ")
	inorder(1, n, node)
	fmt.Println()
	fmt.Print("DFS pre-order  : ")
	preorder(1, n, node)
	fmt.Println()
	fmt.Print("DFS post-order : ")
	postorder(1, n, node)
	fmt.Println()
}

func inorder(i, n int, node []byte) {
	if i != -1 {
		inorder(leftchild(i, n), n, node)
		fmt.Printf("%c ", node[i])
		inorder(rightchild(i, n), n, node)
	}
}

func leftchild(i, n int) int {
	if 2*i <= n {
		return 2 * i
	} else {
		return -1
	}
}

func rightchild(i, n int) int {
	if 2*i+1 <= n {
		return 2*i + 1
	} else {
		return -1
	}
}

func preorder(i, n int, node []byte) {
	if i != -1 {
		fmt.Printf("%c ", node[i])
		preorder(leftchild(i, n), n, node)
		preorder(rightchild(i, n), n, node)
	}
}

func postorder(i, n int, node []byte) {
	if i != -1 {
		postorder(leftchild(i, n), n, node)
		postorder(rightchild(i, n), n, node)
		fmt.Printf("%c ", node[i])
	}
}
