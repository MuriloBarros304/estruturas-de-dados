package main

import(
    "fmt"
    "errors"
)

type ArrayList interface {
    v []int
    size int
}

func (l *ArrayList) Init(size int) error {
    if size > 0 {
        l.v = make([]int, size)
        return nil
    } else {
        return errors.New("Não é possível criar um vetor de tamanho negativo")
    }
}

func (l *ArrayList) Add(e int) {

}