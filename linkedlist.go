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

type Node struct {
    v int
    next *Node
}

type LinkedList struct {
    head *Node
    inserted int
}

func (l *LinkedList) Size() int {
    return l.inserted
}

func (l *LinkedList) Get(index int) (int,error){
    if index>=0 && index < l.inserted {
        aux := l.head

        for i := 0; i < index; i++ {
            aux = aux.next
        }
        return aux.v, nil

    } else {
        return index, errors.New("Index fora dos limites da lista") 
    }
}

func (l *LinkedList) Add(e int) {
    newNode := &Node{v:e}
    aux := l.head
    if aux == nil {
        l.head = newNode
    } else {
        aux := l.head
        for aux.next != nil {
            aux = aux.next
        }
        aux.next = newNode
    }
    l.inserted++
}

func (l *LinkedList) AddOnIndex(e int, index int) error {
    if index >= 0 && index <= l.inserted {
        newNode := &Node{v:e}
        if l.head == nil {
            l.head = newNode
        } else {
            aux := l.head
            for i := 0; i < index - 1; i++ {
                aux = aux.next
            }
            newNode.next = aux.next
            aux.next = newNode
        }
        l.inserted++
        return nil
    } else {
        return errors.New("Index invalido") 
    }
}

func (l *LinkedList) Remove(index int) error {
    if index >= 0 && index < l.inserted {
        if index == 0 {
            l.head = l.head.next
        } else {
            aux := l.head
            for i := 0; i < index - 1; i++ {
                aux = aux.next
            }
            aux.next = aux.next.next
        }

        l.inserted--
        return nil
    } else {
        return errors.New("Index invalido")
    } 
}
func main(){
    l := &LinkedList{}
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
    //fmt.Println(l.Get(2))
    fmt.Println(l)
    //l.AddOnIndex(-1,0)
}