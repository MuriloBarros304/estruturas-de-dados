package main

import (
    "fmt"
    "errors"
)

type Queue interface {
    Enqueue(e int)
    Dequeue() (int, error)
    Peek() (int, error)
    IsEmpty() bool
    Size() int
}

type SimpleQueue struct {
    v []int
    size int
}


