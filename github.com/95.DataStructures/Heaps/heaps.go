//heaps can used for priority queues[highest priority served first]
//Heap is a datastructure, that can be expressed as a complete tree.
//a complete tree means - all the level[left and right] in the tree are full, except for the lowest level.
//where some nodes can be empty. but empty nodes should in the right side.
// two type: max heap and min heap

//max heap:
//largest key will be stored in the root and all the nodes should be larger than its child node, so getting the highest key would be so fast

//min heap:
//root should be smallest key, and each node should be smaller then the child nodes

//below shows the max heaps example. Keys are stored as an array
//left index are odd index and right index are even index

//left child index = [[parent index numnber]*2 +1 = left child index]
//right child index = [[parent index numnber]*2 +2 = right child index]

//inserting : whenever adding the new node that should go bottom right side of the heap.
//after adding at the bottom right, we need to rearrange the node to maintain the heap[heapify]. if the child node is larger than the parent, have to swap the numbers. repeat this process until the number reached the correct place.
//Heapify process is required if we extract the largest node[root node] from the array. once removed the largest number, fill the root node with the last node. start swap process starting from the top

//heapify process will be depends on how high the size of the tree is. O(h)=> h is height of the tree [time complexity of heap to insert and extract - O(log n)]

//to do:
//struct for maxheap
//insert method
//extract method

package main

import "fmt"

//MaxHeap, Insert, Extract[uppercase- visible from outside]

//MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

//Insert method adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

//Extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1
	//when array is empty
	if len(h.array) == 0 {
		fmt.Println("Cannpt extract because length of the array is 0")
		return -1
	}
	//take out the last index and put it in the root
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)
	return extracted
}

//maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//loop while index has atleast one child
	for l <= lastIndex {
		if l == lastIndex { //when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { //when left child is larger then right child
			childToCompare = l
		} else { //when right child is larger then left child
			childToCompare = r
		}

		//compare array value of current index to larger child and swap if smaller[index is smaller than the childToCompare so swaped]
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			//after swaped we need to update the variable in the loop
			index = childToCompare
			l, r = left(index), right(index)
		} else { // if the value of the current index is larger than the larger child, (that is the values are in correct position)
			return

		}
	}
}

//get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//get the left child index
func left(i int) int { //i is a parent index
	return i*2 + 1
}

//get the right child index
func right(i int) int {
	return i*2 + 2
}

//swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	fmt.Println(buildHeap)
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
