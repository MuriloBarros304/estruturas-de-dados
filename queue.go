package main

import (
    "fmt"
    "errors"
)

type Queue interface {
    Size() int
    Enqueue(e int) (error)
    Dequeue() (int, error)
    Front() (int, error)
}

type ArrayQueue struct {
    v []int
    front int
    rear int
    inserted int
}

func (q *ArrayQueue) Init(size int) error {
    if size > 0 {
    q.v = make([]int, size)
        return nil
    } else {
        return errors.New("Size <= 0")
    }
}

func (q *ArrayQueue) Size() int {
    return q.inserted
}

func (q *ArrayQueue) Enqueue(e int) error {
    if q.inserted == len(q.v) {
        return errors.New("Fila está cheia")
    }
    if q.inserted == 0 {
        q.front++
        q.rear++
    } else {
        q.rear = (q.rear+1)%len(q.v) // circular caso chegue no final, volta pro início
    }
    q.v[q.rear] = e
    q.inserted++
    return nil
}

func (q *ArrayQueue) Dequeue() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Fila está vazia")
    } else if q.inserted == 1 {
        temp := q.v[q.front]
        q.front = -1
        q.rear = -1
        q.inserted--
        return temp, nil
    } else {
        temp := (q.front+1)%len(q.v) // circular caso chegue no final, volta pro início
        q.inserted--
        q.front = temp
        return q.v[q.front], nil // verificar se é isso mesmo
    }
}

func (q *ArrayQueue) Front() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.v[q.front], nil
    }
}

// Linked List Queue
type Node struct {
    val int
    next *Node
}

type LinkedListQueue struct {
    front *Node
    rear *Node
    inserted int
}

func (q *LinkedListQueue) Size() int {
    return q.inserted
}

func (q *LinkedListQueue) Enqueue(e int) error {
    newNode := &Node{val: e, next: nil}
    if q.inserted == 0 {
        q.front = newNode
        q.rear = newNode
    } else {
        q.rear.next = newNode
        q.rear = newNode
    }
    q.inserted++
    return nil
}

func (q *LinkedListQueue) Dequeue() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Fila está vazia")
    } else {
        temp := q.front
        q.front = q.front.next
        q.inserted--
        return temp.val, nil
    }
}

func (q *LinkedListQueue) Front() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.front.val, nil
    }
}

func main(){
    q := &ArrayDeque{}
    q.Init(7)
    q.EnqueueRear(1)
    q.EnqueueRear(2)
    q.EnqueueRear(3)
    q.EnqueueRear(4)
    q.EnqueueRear(5)
    q.EnqueueRear(6)    
    q.EnqueueRear(7)
    for i:=0; i<7; i++ {
        val, _ := q.DequeueFront()
        fmt.Println(val)
    }
}