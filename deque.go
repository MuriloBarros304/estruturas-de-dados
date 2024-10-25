//Deque
package main
import (
    "fmt"
    "errors"
)

type Deque interface {
    Size() int
    EnqueueFront(e int) (error)
    Enqueue(e int) (error)
    EnqueueRear(e int) (error)
    DequeueFront() (int, error)
    DequeueRear() (int, error)
    Dequeue() (int, error)
    Front() (int, error)
    Rear() (int, error)
}

type ArrayDeque struct {
    v []int
    front int
    rear int
    inserted int
}

func (q *ArrayDeque) Init(size int) error {
    if size > 0 {
    q.v = make([]int, size)
        return nil
    } else {
        return errors.New("Size <= 0")
    }
}

func (q *ArrayDeque) Size() int {
    return q.inserted
}

func (q *ArrayDeque) EnqueueRear(e int) error {
    if q.inserted == len(q.v) {
        return errors.New("Deque está cheio")
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

func (q *ArrayDeque) EnqueueFront(e int) error {
    if q.inserted == len(q.v) {
        return errors.New("Deque está cheio")
    }
    if q.inserted == 0 {
        q.front++
        q.rear++
    } else {
        q.front = (q.front-1+len(q.v))%len(q.v) // verificar soma com len
    }
    q.v[q.front] = e
    q.inserted++
    return nil
}

func (q *ArrayDeque) DequeueFront() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Deque está vazio")
    } else if q.inserted == 1 {
        temp := q.v[q.front]
        q.front = -1
        q.rear = -1
        q.inserted--
        return temp, nil
    } else {
        temp := q.front
        q.inserted--
        q.front = (q.front+1)%len(q.v)
        return q.v[temp], nil
    }
}

func (q *ArrayDeque) DequeueRear() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Deque está vazio")
    } else if q.inserted == 1 {
        temp := q.v[q.rear]
        q.front = -1
        q.rear = -1
        q.inserted--
        return temp, nil
    } else {
        temp := q.rear
        q.inserted--
        q.rear = (q.rear-1+len(q.v))%len(q.v)
        return q.v[temp], nil
    }
}

func (q *ArrayDeque) Front() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.v[q.front], nil
    }
}

func (q *ArrayDeque) Rear() (int, error) {
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.v[q.rear], nil
    }
}
