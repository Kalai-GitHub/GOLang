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
