package main

import (
    "fmt"
    "errors"
)

type List interface {
    Size() int
    Get(index int) (int,error)
    Add(e int) 
    AddOnIndex(e int, index int) error
    Remove(index int) error
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
        return errors.New("Size <= 0")
    }
}

func (l *ArrayList) doubleV() {
    newSize := len(l.v)*2
    newV := make([]int, newSize)
    for i:=0; i < len(l.v); i++ {
        newV[i] = l.v[i]
    } 
    l.v = newV
}

func (l *ArrayList) Size() int{
    return l.inserted
}

func (l *ArrayList) Get(index int) (int,error){
    if index>=0 && index < l.inserted {
        return l.v[index], nil
    } else {
        return index, errors.New("Index fora dos limites da lista") 
    }
}

func (l *ArrayList) Add(e int) {
    if l.inserted == len(l.v) {
        l.doubleV()
    }
    l.v[l.inserted] = e
    l.inserted++
}

func (l *ArrayList) AddOnIndex(e int, index int) error {
    if index>=0 && index <= l.inserted {
        if l.inserted == len(l.v) {
            l.doubleV()
        }
        for i:=l.inserted; i > index; i--{
            l.v[i] = l.v[i-1]
        }

        l.v[index] = e
        l.inserted++
        return nil
    }
    return errors.New("Index invalido") 
}

func (l *ArrayList) Remove(index int)  error {
    if index>=0 && index < l.inserted {
        for i:=index; i < l.inserted-1; i++{
            l.v[i] = l.v[i+1]
        }

        l.inserted--
        return nil
    }
    return errors.New("Index invalido") 
}

func main(){
    l := &ArrayList{}
    l.Init(10)
    l.Add(1)
    l.Add(2)
    l.Add(3)
    l.Add(4)
    l.Add(5)
    l.Add(6)
    l.Add(7)
    l.Add(8)
    l.Add(9)
    l.Add(10)
    l.AddOnIndex(0,0)
    l.Remove(4)
    fmt.Println(l)
    //l.AddOnIndex(-1,0)
}