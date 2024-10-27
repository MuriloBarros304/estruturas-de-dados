package main

import (
    "fmt"
    "errors"
)
// Queue
/*
Queue (ou fila) é uma estrutura de dados que segue a regra FIFO (First In, First Out),
ou seja, o primeiro elemento a ser inserido é o primeiro a ser retirado. Pode ser implementada
de várias formas, como por exemplo, com um array ou com uma lista encadeada.
*/

type Queue interface { // tipo abstrato de dados
    Size() int
    Enqueue(e int) (error)
    Dequeue() (int, error)
    Front() (int, error)
}

type ArrayQueue struct { // estrutura de dados, implementação de fila com array
    v []int              // array
    front int            // frente da fila, que deve ser retirada primeiro
    rear int             // traseira da fila, onde os elementos são inseridos
    inserted int         // quantidade de elementos inseridos
}

func (q *ArrayQueue) Init(size int) error { // O(1) Omega(1)
    if size > 0 {
    q.v = make([]int, size) // cria um array com o tamanho especificado
        return nil
    } else {
        return errors.New("Size <= 0")
    }
}

func (q *ArrayQueue) Size() int { // O(1) Omega(1)
    return q.inserted
}

func (q *ArrayQueue) Enqueue(e int) error {  // O(1) Omega(1)
    if q.inserted == len(q.v) {              // se a quantidade de elementos inseridos for igual ao tamanho do array
        return errors.New("Fila está cheia")
    }
    if q.inserted == 0 {                     // se a quantidade de elementos inseridos for 0, ou seja, fila vazia
        q.front++                            // incrementa a frente, que é a posição do primeiro elemento
        q.rear++                             // incrementa a traseira
    } else {                                 // se a fila não estiver vazia
        q.rear = (q.rear+1)%len(q.v)         // circular caso chegue no final, volta pro início
    }
    q.v[q.rear] = e                          // insere o elemento na traseira
    q.inserted++
    return nil
}

func (q *ArrayQueue) Dequeue() (int,error) { // O(1) Omega(1)
    if q.inserted == 0 {              // se a quantidade de elementos inseridos for 0, ou seja, fila vazia
        return 0, errors.New("Impossível desenfileirar de fila vazia!")
    } else if q.inserted == 1 {       // se a quantidade de elementos inseridos for 1
        temp := q.v[q.front]          // guarda o valor do elemento da frente, o único elemento
        q.front = -1                  // a frente agora é -1, pois não há mais elementos
        q.rear = -1                   // a traseira agora é -1
        q.inserted--
        return temp, nil
    } else {                          // se a quantidade de elementos inseridos for maior que 1
        temp := q.v[q.front]          // guarda o valor do elemento da frente
        q.front = (q.front+1)%len(q.v)// incrementa a frente, circular caso chegue no final, volta pro início
        q.inserted--
        return temp, nil
    }
}

func (q *ArrayQueue) Front() (int, error) { // O(1) Omega(1)
    if q.inserted == 0 {         // se a quantidade de elementos inseridos for 0, ou seja, fila vazia
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.v[q.front], nil // retorna o valor do elemento da frente
    }
}

// Linked List Queue
type Node struct { // nó da fila encadeada
    val int        // valor do nó
    next *Node     // ponteiro para o próximo nó
}

type LinkedListQueue struct { // estrutura de dados, implementação de fila com lista encadeada
    front *Node               // frente da fila
    rear *Node                // cauda da fila
    inserted int              // quantidade de elementos inseridos
}

func (q *LinkedListQueue) Size() int { // O(1) Omega(1)
    return q.inserted
}

func (q *LinkedListQueue) Enqueue(e int) error { // O(1) Omega(1)
    newNode := &Node{val: e, next: nil} // cria um novo nó
    if q.inserted == 0 {                // se a quantidade de elementos inseridos for 0, ou seja, fila vazia
        q.front = newNode               // a frente é o novo nó
        q.rear = newNode                // a traseira é o novo nó, pois é o único elemento
    } else {                            // se a fila não estiver vazia
        q.rear.next = newNode           // o próximo do nó da traseira é o novo nó
        q.rear = newNode                // a traseira é o novo nó, atualizando a traseira
    }
    q.inserted++
    return nil
}

func (q *LinkedListQueue) Dequeue() (int, error) { // O(1) Omega(1)
    if q.inserted == 0 {                           // se a quantidade de elementos inseridos for 0, ou seja, fila vazia
        return 0, errors.New("Fila está vazia")
    } else {                                       // se a fila não estiver vazia
        temp := q.front                            // guarda o nó da frente
        q.front = q.front.next                     // a frente agora é o próximo nó, liberando o nó da frente
        q.inserted--
        return temp.val, nil                       // retorna o valor do nó da frente
    }
}

func (q *LinkedListQueue) Front() (int, error) { // O(1) Omega(1)
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.front.val, nil                  // retorna o valor do nó da frente
    }
}

func main(){
    q := &ArrayDeque{}
    q.Init(7)
    q.EnqueueRear(1)
    q.EnqueueRear(2)
    q.EnqueueRear(3)
    q.EnqueueRear(4)
    q.EnqueueRear(5)
    q.EnqueueRear(6)    
    q.EnqueueRear(7)
    for i:=0; i<7; i++ {
        val, _ := q.DequeueFront()
        fmt.Println(val)
    }
}