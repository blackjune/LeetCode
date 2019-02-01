package main

import (
	"fmt"
)

type MyCircularQueue struct {
	data   []int
	head   int
	tail   int
	length int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{data: make([]int, k)}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.tail] = value
	this.tail = (this.tail + 1) % len(this.data)
	this.length++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.data)
	this.length--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	rear := this.tail - 1
	if rear < 0 {
		rear = this.length - 1
	}
	return this.data[rear]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.length == len(this.data)
}

func main() {
	queue := Constructor(10)
	fmt.Println(queue.EnQueue(1))
	fmt.Println(queue.EnQueue(2))
	fmt.Println(queue.EnQueue(3))
	fmt.Printf("front %d, rear %d, %#v\n", queue.Front(), queue.Rear(), queue)
	queue.DeQueue()
	fmt.Printf("front %d, rear %d, %#v\n", queue.Front(), queue.Rear(), queue)
	queue.DeQueue()
	fmt.Printf("front %d, rear %d, %#v\n", queue.Front(), queue.Rear(), queue)
	queue.DeQueue()
	fmt.Printf("front %d, rear %d, %#v\n", queue.Front(), queue.Rear(), queue)
	fmt.Println(queue.EnQueue(4))
	queue.DeQueue()
	queue.DeQueue()

	fmt.Println(queue.EnQueue(5))
	fmt.Printf("front %d, rear %d, %#v\n", queue.Front(), queue.Rear(), queue)
}
