//Linked list - adding and deleting the values at the beginning of the list is faster [time complexity - O(1)]. this is complecated in Array

//Doubly linked list - nodes contains the address of the next node and the previous node. Adding and deleting the values at the end of the list is faster.

package main

import "fmt"

type node struct {
	data int
	next *node //address of the next node
}

type linkedlist struct {
	head   *node //head holds pointer of the node
	length int   //length of the linked list
}

//adding node in the beginning
func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

//print out the data of everynode in the linkedlist
func (l linkedlist) printLinkedList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%v\t", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

//delete the given value from the list
func (l *linkedlist) deleteValue(value int) {
	//if linkedlist is empty
	if l.length == 0 {
		return
	}

	//if the given value is in the first index of the linkedlist
	if l.head.data == value {
		l.head = l.head.next
		l.length--
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		//if the given value is not present
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedlist{}
	node1 := &node{data: 20}
	node2 := &node{data: 30}
	node3 := &node{data: 40}
	node4 := &node{data: 50}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	myList.printLinkedList()

	myList.deleteValue(40)
	myList.printLinkedList()

	myList.deleteValue(20) //value is present in the first node
	myList.printLinkedList()

	myList.deleteValue(60) //value is not present in myList
	myList.printLinkedList()

	emptylist := linkedlist{}
	emptylist.deleteValue(90) //deleting the value from empty list

}
