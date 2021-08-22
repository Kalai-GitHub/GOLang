//Binary Search Tree-  each parent will have two child nodes [ Roor, Parent, Childrean, leaves]
//A tree that has nodes, with no more than two children is called binary tree [ smaller nodes placed left side and larger nodes will be on the right side]
//some of the right or left node can be empty
//advantage of binary tree is speed. insert and search the data into a binary tree is O(h) this is better than lon(n)

package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//insert will add value to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(v int) {
	if n.Key < v {
		if n.Right == nil {
			//n.right.Key = v
			n.Right = &Node{Key: v}
		} else {
			n.Right.Insert(v)
		}
	} else if n.Key > v {
		if n.Left == nil {
			n.Left = &Node{Key: v}
		} else {
			n.Left.Insert(v)
		}
	}
}

//Binary search take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k { //move right
		return n.Right.Search(k)
	} else if n.Key > k { //move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	myTree := &Node{Key: 56}

	fmt.Println(myTree)
	myTree.Insert(10)
	fmt.Println(myTree)

	myTree.Insert(20)
	myTree.Insert(30)
	myTree.Insert(40)
	myTree.Insert(50)
	myTree.Insert(60)
	myTree.Insert(70)

	fmt.Println(myTree.Search(60)) //present in the tree
	fmt.Println("count", count)
	fmt.Println(myTree.Search(100)) //not present in the tree

}
