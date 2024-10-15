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
        return q.v[q.front-1], nil // verificar se é isso mesmo
    }
}

func (q *ArrayQueue) Front() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.v[q.front], nil
    }
}

func main(){
    q := &ArrayQueue{}
    q.Init(6)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)
    q.Enqueue(5)
    q.Enqueue(6)
    val, _ := q.Dequeue()
    //fmt.Println(val)
    val, _ = q.Dequeue()
    //fmt.Println(val)
    val, _ = q.Dequeue()
    //fmt.Println(val)
    val, _ = q.Front()
    fmt.Println(val)
    q.Enqueue(7)
    q.Enqueue(8)
    val, _ = q.Dequeue()
    //fmt.Println(val)
    val, _ = q.Dequeue()
    val, _ = q.Dequeue()
    val, _ = q.Dequeue()
    //fmt.Println(val)
    val, _ = q.Front()
    fmt.Println(val)
    fmt.Println(q.Size())
}