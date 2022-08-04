package main

import "fmt"

type Queue struct {
	Values []int
}

//Enqueue

func (q *Queue) Enqueue(v int) {
	q.Values = append(q.Values, v)
}

//Dequeue

func (q *Queue) Dequeue() int {
	toRemove := q.Values[0]
	q.Values = q.Values[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Enqueue(400)
	myQueue.Enqueue(500)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
