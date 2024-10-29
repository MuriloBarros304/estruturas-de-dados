//Deque
package main
import (
    "fmt"
    "errors"
)

// Deque
/*
A Deque é um tipo abstrato de dados que permite a inserção e remoção de elementos
em ambas as extremidades da fila, ou seja, é uma fila duplamente terminada.
Pode ser implementado de várias formas, como por exemplo, com um array ou com uma lista encadeada.
*/

type Deque interface { // tipo abstrato de dados
    Size() int
    EnqueueFront(e int) (error) // enfileirar na frente
    EnqueueRear(e int) (error)  // enfileirar na traseira
    DequeueFront() (int, error) // desenfileirar da frente
    DequeueRear() (int, error)  // desenfileirar da traseira
    Front() (int, error)        // obter frente
    Rear() (int, error)         // obter traseira
}

type ArrayDeque struct { // estrutura de dados, implementação de deque com array
    v []int       // array
    front int     // frente
    rear int      // traseira
    inserted int  // quantidade de elementos inseridos
}

func (q *ArrayDeque) Init(size int) error { // O(1) Omega(1)
    if size > 0 {
        q.v = make([]int, size) // cria um array com o tamanho especificado, fazendo a alocação de memória
            return nil
    } else {
        return errors.New("Size <= 0")
    }
}

func (q *ArrayDeque) Size() int { // O(1) Omega(1)
    return q.inserted
}

func (q *ArrayDeque) EnqueueRear(e int) error { // O(1) Omega(1)
    if q.inserted == len(q.v) {
        return errors.New("Deque está cheio")
    }
    if q.inserted == 0 {             // se a quantidade de elementos inseridos for 0, ou seja, deque vazio
        q.front++                    // incrementa a frente
        q.rear++                     // incrementa a traseira, que é a posição do único elemento
    } else {                         // se a deque não estiver vazia
        q.rear = (q.rear+1)%len(q.v) // circular caso chegue no final, volta pro início, não passando de len(q.v)
    }
    q.v[q.rear] = e                  // insere o elemento na traseira
    q.inserted++
    return nil
}

func (q *ArrayDeque) EnqueueFront(e int) error { // O(1) Omega(1)
    if q.inserted == len(q.v) {
        return errors.New("Deque está cheio")
    }
    if q.inserted == 0 {                        // se a quantidade de elementos inseridos for 0, ou seja, deque vazio
        q.front++                               // incrementa a frente
        q.rear++                                // incrementa a traseira, que é a posição do único elemento
    } else {                                    // se a deque não estiver vazia
        q.front = (q.front-1+len(q.v))%len(q.v) // circular caso chegue no início, volta pro final, não passando de len(q.v)
                                                // a soma com len(q.v) é para garantir que o resultado seja positivo, decrementando
                                                // a frente, ou seja, indo para a posição anterior
    }
    q.v[q.front] = e                            // insere o elemento na frente
    q.inserted++
    return nil
}

func (q *ArrayDeque) DequeueFront() (int, error) { // O(1) Omega(1)
    if q.inserted == 0 {
        return 0, errors.New("Deque está vazio")
    } else if q.inserted == 1 {        // se a quantidade de elementos inseridos for 1
        temp := q.v[q.front]           // guarda o valor do elemento da frente, o único elemento
        q.front = -1                   // a frente agora é -1, pois não há mais elementos
        q.rear = -1                    // a traseira agora é -1
        q.inserted--
        return temp, nil               // retorna o valor do elemento da frente
    } else {                           // se a quantidade de elementos inseridos for maior que 1
        temp := q.front                // guarda a posição da frente
        q.inserted-- 
        q.front = (q.front+1)%len(q.v) // incrementa a frente, circular caso chegue no final, volta pro início
        return q.v[temp], nil          // retorna o valor do elemento da frente
    }
}

func (q *ArrayDeque) DequeueRear() (int, error) { // O(1) Omega(1)
    if q.inserted == 0 {
        return 0, errors.New("Deque está vazio")
    } else if q.inserted == 1 {               // se a quantidade de elementos inseridos for 1
        temp := q.v[q.rear]                   // guarda o valor do elemento da traseira, o único elemento
        q.front = -1                          // a frente agora é -1, pois não há mais elementos
        q.rear = -1                           // a traseira agora é -1
        q.inserted--
        return temp, nil                      // retorna o valor do elemento da traseira
    } else {                                  // se a quantidade de elementos inseridos for maior que 1
        temp := q.rear                        // guarda a posição da traseira
        q.inserted--
        q.rear = (q.rear-1+len(q.v))%len(q.v) // decrementa a traseira, circular caso chegue no início, volta pro final
                                              // a soma com len(q.v) é para garantir que o resultado seja positivo, decrementando
                                              // a traseira, ou seja, indo para a posição anterior, se a traseira for 0, volta para len(q.v)-1
        return q.v[temp], nil                 // retorna o valor do elemento da traseira
    }
}

func (q *ArrayDeque) Front() (int, error) { // O(1) Omega(1)
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.v[q.front], nil // retorna o valor do elemento da frente
    }
}

func (q *ArrayDeque) Rear() (int, error) { // O(1) Omega(1)
    if q.inserted == 0 {
        return 0, errors.New("Impossível obter frente de fila vazia")
    } else {
        return q.v[q.rear], nil // retorna o valor do elemento da traseira
    }
}
