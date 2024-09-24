package main

import(
    "fmt"
    "errors"
)

type List interface {
    Add(e int)
    AddOnIndex(e int, index int) error
    RemoveOnIndex(index int) error
    Get(index int) (int, error)
    Set(e int, index int) error
    Size() int
}

type ArrayList struct {
    v []int
    inserted int
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
    if l.inserted == len(l.v) {
        l.doubleV()
    }
    l.v[l.inserted] = e
    l.inserted++
}

func (l *ArrayList) doubleV() {
    newSize := make([]int, len(l.v)*2)
    for i := 0; i < len(l.v); i++ {
        newSize[i] = l.v[i]
    }
    l.v = newSize
}

func (l *ArrayList) Size() int {
    return l.inserted
}

func (l* ArrayList) Get(index int) (int, error) {
    if index >= 0 && index < l.inserted {
        return l.v[index], nil
    } else {
        return 0, errors.New("Fora dos limites do vetor")
    }
}

func (l *ArrayList) AddOnIndex(e int, index int) error {
    if l.inserted == len(l.v) {
        l.doubleV()
    }
    for i := l.inserted; i > index; i-- {
        l.v[i] = l.v[i-1]
    }

    l.v[index] = e
    l.inserted++
    if index > len(l.v) {
        return errors.New("Fora dos limites do vetor")
    } else {
        return nil
    }
}

func main() {
    l := &ArrayList{}
    l.Init(1)
    l.Add(1)
    l.Add(2)
    l.Add(3)
    l.Add(4)
    l.Add(5)
    l.Add(6)
    l.Add(7)
    l.Add(8)
    l.Add(9)
    l.AddOnIndex(0, 0)
    fmt.Println(l)
}

