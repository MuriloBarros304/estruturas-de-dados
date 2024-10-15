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

// doubly Linked List
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

// linked list 
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

func (l *LinkedList) AddOnIndex(e int, index int) error {
    if index >= 0 && index <= l.inserted {
        newNode := &Node{v:e}
        if l.head == nil {
            l.head = newNode
        } else {
            if index == 0 {
                newNode.next = l.head
                l.head = newNode
            } else {
                aux := l.head
                for i := 0; i < index - 1; i++ {
                    aux = aux.next
                }
                newNode.next = aux.next
                aux.next = newNode
            }
        }
        l.inserted++
        return nil
    } else {
        return errors.New("Index invalido") 
    }
}

func (l *LinkedList) Add(e int) {
    l.AddOnIndex(e, l.inserted)
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

// main
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