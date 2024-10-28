package main

import (
    "fmt"
    "errors"
)
// Stack
/*
stack (ou pilha) é uma estrutura de dados que segue a regra LIFO (Last In, First Out),
ou seja, o último elemento a ser inserido é o primeiro a ser retirado. Pode ser implementada
de várias formas, como por exemplo, com um array ou com uma lista encadeada.
*/
type Stack interface { // tipo abstrato de dados
    Size() int
    Push(e int)
    Pop() (int, error)
    Top() (int, error)
}

type Node struct { // nó da lista encadeada
    v int          // elemento
    next *Node     // próximo nó
}

type LinkedStack struct { // pilha com lista encadeada
    top *Node             // topo, último elemento inserido
    inserted int          // quantidade de elementos inseridos
}

func (s *LinkedStack) Size() int { // O(1) Omega(1)
    return s.inserted
}

func (s *LinkedStack) Push(e int) { // O(1) Omega(1)
    // criar novo nó, atualizar próximo, atualizar topo
    newNode := &Node{v:e}    // instancia um novo nó
    if s.top != nil {        // se o topo não for nulo
        newNode.next = s.top // o próximo do novo nó é o topo, ou seja, o antigo topo é o próximo do novo nó
    }
    s.top = newNode          // o topo agora é o novo nó, funciona caso a pilha esteja vazia
    s.inserted++
}

func (s *LinkedStack) Pop() (int, error) { // O(1) Omega(1)
    // fazer o topo apontar pro próximo
    if s.top != nil {      // se o topo não for nulo
        s.inserted--
        val := s.top.v     // guarda o valor do topo
        s.top = s.top.next // o topo agora é o próximo do antigo topo
        return val, nil    // retorna o valor do antigo topo
    } else {
        return 0, errors.New("Empty Stack")
    }
}

func (s *LinkedStack) Top() (int, error) { // O(1) Omega(1)
    if s.top != nil {       // se o topo não for nulo
        return s.top.v, nil // retorna o valor do topo
    } else {                // se o topo for nulo
        return 0, errors.New("Empty Stack") // retorna erro
    }
}

func main(){
    s := &LinkedStack{}
    s.Push(1)
    s.Push(4)
    s.Push(8)
    s.Push(12)    
    for i := 0; i < 5; i++ {
        val, err := s.Pop()
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(val)
        }
    }
}