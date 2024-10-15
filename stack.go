package main

import (
    "fmt"
    "errors"
)

type Stack interface {
    Size() int
    Push(e int)
    Pop() (int, error)
    Top() (int, error)
}

type Node struct {
    v int
    next *Node
}

type LinkedStack struct {
    top *Node
    inserted int
}

func (s *LinkedStack) Size() int {
    return s.inserted
}

func (s *LinkedStack) Push(e int) {
    // criar novo nó, atualizar próximo, atualizar topo
    newNode := &Node{v:e}
    if s.top != nil {
        newNode.next = s.top
    }
    s.top = newNode
    s.inserted++
}

func (s *LinkedStack) Pop() (int, error) {
    // fazer o topo apontar pro próximo
    if s.top != nil {
        s.inserted--
        val := s.top.v
        s.top = s.top.next
        return val, nil
    } else {
        return 0, errors.New("Empty Stack")
    }
}

func (s *LinkedStack) Top() (int, error) {
    if s.top != nil {
        return s.top.v, nil
    } else {
        return 0, errors.New("Empty Stack")
    }
}

func main(){
    s := &LinkedStack{}
    s.Push(1)
    s.Push(4)
    s.Push(8)
    s.Push(12)    
    fmt.Println(s.Top())
}