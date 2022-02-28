package queue

import (
	"errors"
	"fmt"
)

//CircleQueue ring queue
type CircleQueue struct {
	MaxSize int
	Array   []string
	Front   int
	Rear    int
}

func Init(size int) (q *CircleQueue, err error) {
	q = &CircleQueue{MaxSize: size + 1}
	q.Array = make([]string, size+1)
	q.Front = 0
	q.Rear = 0
	return
}

//Push adds a value to the queue
func (q *CircleQueue) Push(val string) (err error) {
	//First determine whether the queue is full
	if q.IsFull() {
		return errors.New("Queue is full")
	}
	q.Array[q.Rear] = val
	//The end of the line does not contain elements
	//q.Rear++
	q.Rear = (q.Rear + 1) % q.MaxSize
	return
}

//Pop gets a value
func (q *CircleQueue) Pop() (val string, err error) {
	if q.IsEmpty() {
		return "", errors.New("Queue is empty")
	}
	//Head of the team contains elements
	val = q.Array[q.Front]
	//q.Front++
	q.Front = (q.Front + 1) % q.MaxSize
	return val, err
}

//IsFull The queue is full
func (q *CircleQueue) IsFull() bool {
	return (q.Rear+1)%q.MaxSize == q.Front
}

//IsEmpty whether the queue is empty
func (q *CircleQueue) IsEmpty() bool {
	return q.Front == q.Rear
}

//Size The size of the queue
func (q *CircleQueue) Size() int {
	return (q.Rear + q.MaxSize - q.Front) % q.MaxSize
}

//Show show queue
func (q *CircleQueue) Show() {
	//Get how many elements in the current queue
	size := q.Size()
	if size == 0 {
		fmt.Println("The queue is empty")
	}
	tmpArray := []string{}
	//Auxiliary variable, pointing to Front
	tmpFront := q.Front
	for i := 0; i < size; i++ {
		//fmt.Printf("queue[%d]=%v\t", tmpFront, q.Array[tmpFront])
		tmpArray = append(tmpArray, q.Array[tmpFront])
		tmpFront = (tmpFront + 1) % q.MaxSize
	}
	fmt.Printf("queue %v, size %v\n", tmpArray, len(tmpArray))
}
func (q *CircleQueue) ShowFull() {
	//Get how many elements in the current queue
	size := q.Size()
	if size == 0 {
		fmt.Println("The queue is empty")
	}
	//Auxiliary variable, pointing to Front
	tmpFront := q.Front
	for i := 0; i < size; i++ {
		fmt.Printf("queue[%d]=%v\t", tmpFront, q.Array[tmpFront])
		tmpFront = (tmpFront + 1) % q.MaxSize
	}

}
