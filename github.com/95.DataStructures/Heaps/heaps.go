//to do for MaxHeap:
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
