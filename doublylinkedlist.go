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
    newNode := &Node{v:e} // cria um novo nó
    if l.head == nil { // verifica se a lista é vazia
        l.head = newNode
    } else {
        l.tail.next = newNode
        newNode.prev = l.tail // ajustar prev de newnode
    }
    l.tail = newNode
    l.inserted++
}

func (l *LinkedList) AddOnIndex(e int, index int) error {
    if index >= 0 && index <= l.inserted {
        newNode := &Node{v:e}
        if index == 0 { // se for uma inserção no início
            if l.head == nil {
                l.head = newNode
                l.tail = newNode
            } else {
                newNode.next = l.head
                l.head.prev = newNode
                l.head = newNode
            }
        } else {
            aux := l.head
            for i := 0; i < index - 1; i++ {
                aux = aux.next
            }
            newNode.next = aux.next
            newNode.prev = aux
            aux.next = newNode
            if newNode.next != nil { // se não estiver no final da lista
                newNode.next.prev = newNode // o anterior do próximo vai ser o novo nó
            } else { // se estiver no final
                l.tail = newNode // atualizar tail para o novo nó
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
    for i := 0; i < 6; i++ {
        l.Add(i)
    }
    l.AddOnIndex(15,0)
    l.AddOnIndex(6,3)
    //l.Remove(4)
    //fmt.Println(l.Get(2))
    for i := 0; i < l.Size(); i++ {
        fmt.Println(l.Get(i))
    }
    //l.AddOnIndex(-1,0)
}