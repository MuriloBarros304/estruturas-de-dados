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
    prev *Node
}

type LinkedList struct {
    head *Node
    tail *Node
    inserted int
}

func (l *LinkedList) Size() int {
    return l.inserted
}

func (l *LinkedList) Get(index int) (int, error){
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
    newNode := &Node{v:e} // Cria um novo nó
    aux := l.head
    prev := aux
    if aux == nil {
        l.head = newNode
        l.tail = newNode
    } else {
        aux := l.head
        for aux.next != nil {
            prev = aux
            aux = aux.next
        }
        aux.next = newNode
        newNode.prev = prev
    }
    l.inserted++
}

func (l *LinkedList) AddOnIndex(e int, index int) error {
    if index >= 0 && index <= l.inserted {
        newNode := &Node{v:e}
        if l.head == nil {
            l.head = newNode
            l.tail = newNode
        } else {
            aux := l.head
            for i := 0; i < index - 1; i++ {
                aux = aux.next
            }
            newNode.next = aux.next
            newNode.prev = aux
            aux.next = newNode
            if newNode.next != nil {
                newNode.next.prev = newNode
            } else {
                l.tail = newNode
            }
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
            aux.next.prev = aux
        }

        l.inserted--
        return nil
    } else {
        return errors.New("Index invalido")
    } 
}

func main(){
    l := &LinkedList{}
    for i := 0; i < 10; i++ {
        l.Add(i)
    }
    l.AddOnIndex(15,0)
    l.Remove(4)
    //fmt.Println(l.Get(2))
    for i := 0; i < l.Size(); i++ {
        fmt.Println(l.Get(i))
    }
    //l.AddOnIndex(-1,0)
}