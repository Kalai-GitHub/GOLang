//Queue worked as FIFO, has two methods and enqueue and dequeue
package main

import "fmt"

//Queue struct for holding the items in a slice
type Queue struct {
	items []int
}

//Enqueue adds the value at the end
func (q *Queue) enqueue(val int) {
	q.items = append(q.items, val)
}

//Dequeue
func (q *Queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(100)
	myQueue.enqueue(200)
	myQueue.enqueue(300)
	myQueue.enqueue(400)
	myQueue.enqueue(500)
	fmt.Println(myQueue)

	myQueue.dequeue()
	fmt.Println(myQueue)

	myQueue.dequeue()
	fmt.Println(myQueue)

}
