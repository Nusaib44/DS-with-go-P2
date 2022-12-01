package main

import "fmt"

type node struct {
	key   int
	left  *node
	right *node
}

type tree struct {
	root *node
}

// tree
func (t *tree) insertValue(val3 int) {
	if t.root == nil {
		t.root = &node{key: val3}
	} else {
		t.root.insertValue(val3)
	}
}

// Insert
func (n *node) insertValue(val1 int) {

	// [value < key :=left]
	if val1 > n.key {
		if n.right == nil {
			n.right = &node{key: val1}
		} else {
			n.right.insertValue(val1)
		}
	}

	// [value > key :=right]
	if val1 < n.key {
		if n.left == nil {
			n.left = &node{key: val1}
		} else {
			n.left.insertValue(val1)
		}

	}
}

// search
func (n *node) search(val2 int) bool {

	if n == nil {
		return false
	}

	if val2 > n.key {
		return n.right.search(val2)
	}

	if val2 < n.key {
		return n.left.search(val2)
	}
	return true
}

func print(n *node) {
	if n == nil {
		return
	} else {
		print(n.left)
		print(n.right)
		fmt.Printf("%d", n.key)
	}
}

func printPreOrder(n *node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%d ", n.key)
	}
	return
}

func printInOrder(n *node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%d ", n.key)
		printInOrder(n.right)
	}

}
func printRight(n *node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		printRight(n.right)
	}
}

func minValued(root *node) *node {
	temp := root
	for nil != temp && temp.left != nil {
		temp = temp.left
	}
	return temp
}

// Delete node from binary search tree
func Delete(root *node, val int) *node {
	if nil == root {
		return root
	}
	if root.key > val {
		root.left = Delete(root.left, val)
	}
	if root.key < val {
		root.right = Delete(root.right, val)
	}
	if root.key == val {
		if root.left == nil && root.right == nil {
			root = nil
			return root
		}
		if root.left == nil && root.right != nil {
			temp := root.right
			root = nil
			root = temp
			return root
		}
		if root.left != nil && root.right == nil {
			temp := root.left
			root = nil
			root = temp
			return root
		}

		tempNode := minValued(root.right)
		root.key = tempNode.key
		root.right = Delete(root.right, tempNode.key)
	}
	return root
}

func main() {

	myTree := node{key: 100}

	myTree.insertValue(200)
	myTree.insertValue(20)
	myTree.insertValue(500)
	myTree.insertValue(600)
	myTree.insertValue(300)
	myTree.insertValue(40)
	myTree.insertValue(50)
	myTree.insertValue(60)
	myTree.insertValue(2)
	myTree.insertValue(1)

	printPreOrder(&myTree)
	println("")
	printPostOrder(&myTree)
	println("")
	printInOrder(&myTree)
	println("")
	printRight(&myTree)
	println("")
	Delete(&myTree, 20)
	printInOrder(&myTree)
	println("")
	fmt.Println(myTree.search(400))
	println("")
}
