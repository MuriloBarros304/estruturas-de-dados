package main

import (
    "fmt"
    "errors"
)
// Lista
/*
As listas são estruturas de dados que armazenam uma coleção de elementos,
o acesso aos elementos é feito através de índices, que são inteiros que
representam a posição do elemento na lista. As listas podem ser implementadas
de várias formas, como arrays, linked lists, doubly linked lists, etc.
*/
type List interface { // tipo abstrato de dados
    Size() int
    Get(index int) (int,error)
    Add(e int) 
    AddOnIndex(e int, index int) error
    Remove(index int) error
}

type ArrayList struct { // estrutura de dados
    v []int
    inserted int
}

func (l *ArrayList) Init(size int) error { // inicializa a lista, fazendo a alocação de memória
    if size > 0 {
        l.v = make([]int, size)
        return nil
    } else {
        return errors.New("Size <= 0") // caso o tamanho seja menor ou igual a zero
    }
}

func (l *ArrayList) doubleV() { // duplica o tamanho do vetor
    newSize := len(l.v)*2
    newV := make([]int, newSize)
    for i:=0; i < len(l.v); i++ {
        newV[i] = l.v[i]        // complexidade O(n), Omega(1)
    } 
    l.v = newV
}

func (l *ArrayList) Size() int{ // O(1) Omega(1)
    return l.inserted
}

func (l *ArrayList) Get(index int) (int,error){ // O(1) Omega(1)
    if index>=0 && index < l.inserted {         // proteção para index fora dos limites da lista
        return l.v[index], nil                  // retorna o elemento na posição index
    } else {
        return index, errors.New("Index fora dos limites da lista") 
    }
}

func (l *ArrayList) Add(e int) { // O(n) Omega(1)
    if l.inserted == len(l.v) {  // se o vetor está cheio, dobra o tamanho, no melhor caso o tamanho não é alterado
        l.doubleV()
    }
    l.v[l.inserted] = e
    l.inserted++
}

func (l *ArrayList) AddOnIndex(e int, index int) error { // O(n) Omega(1)
    if index>=0 && index <= l.inserted {
        if l.inserted == len(l.v) {                      // se o vetor está cheio
            l.doubleV()
        }
        for i:=l.inserted; i > index; i--{ // i vai do último elemento até o index, movendo os elementos para a direita
            l.v[i] = l.v[i-1]              // para abrir espaço para o novo elemento
        }

        l.v[index] = e
        l.inserted++
        return nil
    }
    return errors.New("Index invalido") 
}

func (l *ArrayList) Remove(index int)  error { // O(n) Omega(1)
    if index>=0 && index < l.inserted {
        for i:=index; i < l.inserted-1; i++{
            l.v[i] = l.v[i+1] // i vai do index até o penúltimo elemento, movendo os elementos para a esquerda
        }                     // para preencher o espaço do elemento removido, fechando o espaço
        l.inserted--
        return nil
    }
    return errors.New("Index invalido") 
}

// doubly Linked List -----------------------------------------------------------------------------------
type Node struct { // nó da lista duplamente encadeada
    v int          // elemento do nó
    next *Node     // ponteiro para o próximo nó
    prev *Node     // ponteiro para o nó anterior
}

type LinkedList struct { // estrutura de dados do tipo abstrato de dados: lista
    head *Node           // ponteiro para o primeiro nó
    tail *Node           // ponteiro para o último nó
    inserted int         // quantidade de elementos na lista
}

func (l *LinkedList) Size() int { // O(1) Omega(1)
    return l.inserted             // retorna a quantidade de elementos na lista
}

func (l *LinkedList) Get(index int) (int, error){ // O(n) Omega(1)
    if index>=0 && index < l.inserted { // proteção para index fora dos limites da lista
        aux := l.head                   // auxiliar para percorrer a lista
        for i := 0; i < index; i++ {    // percorre a lista até o index, por isso a complexidade O(n)
            aux = aux.next              // primeiro a cabeça recebe o próximo
        }
        return aux.v, nil
    } else {
        return index, errors.New("Index fora dos limites da lista") 
    }
}

func (l *LinkedList) Add(e int) { // O(1) Omega(1)
    newNode := &Node{v:e}     // instancia um novo nó
    if l.head == nil {        // verifica se a lista é vazia
        l.head = newNode      // se for, o novo nó é a cabeça
    } else {                  // se não for vazia
        l.tail.next = newNode // a cauda aponta para o novo nó, 
        newNode.prev = l.tail // ajustar prev de newnode para a antiga cauda
    }
    l.tail = newNode          // atualiza a cauda para o novo nó
    l.inserted++
}

func (l *LinkedList) AddOnIndex(e int, index int) error { // O(n) Omega(1)
    if index >= 0 && index <= l.inserted {  // proteção para index fora dos limites da lista
        newNode := &Node{v:e}               // instancia um novo nó
        if index == 0 {                     // se for uma inserção no início
            if l.head == nil {              // se a lista for vazia
                l.head = newNode            // o novo nó é a cabeça
                l.tail = newNode            // e a cauda, pois só tem um elemento
            } else {                        // se não for vazia
                newNode.next = l.head       // o próximo do novo nó é a cabeça
                l.head.prev = newNode       // o anterior da cabeça é o novo nó
                l.head = newNode            // a cabeça é o novo nó, já que a adição é no início
            }
        } else {                            // se não for no início
            aux := l.head                   // auxiliar para percorrer a lista
            for i := 0; i < index - 1; i++ {// percorre a lista até o index - 1, uma posição antes do index
                aux = aux.next              
            }
            newNode.next = aux.next         // o próximo do novo nó é o próximo do auxiliar
            newNode.prev = aux              // o anterior do novo nó é o auxiliar
            aux.next = newNode              // o próximo do auxiliar é o novo nó
            if newNode.next != nil {        // se não estiver no final da lista
                newNode.next.prev = newNode // o anterior do próximo vai ser o novo nó
            } else {                        // se estiver no final
                l.tail = newNode            // atualizar tail para o novo nó
            }
        }
    l.inserted++
    return nil
    } else {
        return errors.New("Index invalido") 
    }
}

func (l *LinkedList) Remove(index int) error { // O(n) Omega(1)
    if index >= 0 && index < l.inserted {
        if index == 0 {                   // se for a remoção do início
            l.head = l.head.next          // a cabeça recebe o próximo
            if l.head != nil {
                l.head.prev = nil         // se a lista não estiver vazia após a remoção, o prev da nova cabeça é nil
            } else {
                l.tail = nil              // se a lista estiver vazia, a cauda também é nil
            }
        } else {                          // se não for do início
            aux := l.head                 // auxiliar para percorrer a lista
            for i := 0; i < index-1; i++ {
                aux = aux.next
            }
            toRemove := aux.next
            aux.next = toRemove.next      // o próximo do auxiliar é o próximo do próximo, removendo o elemento
            if toRemove.next != nil {     // se não for o último elemento
                toRemove.next.prev = aux  // o anterior do próximo do próximo é o auxiliar
            } else {
                l.tail = aux              // se for o último elemento, atualiza a cauda
            }
        }
        l.inserted--
        return nil
    }
    return errors.New("Index invalido")
}

/*
// linked list 
type Node struct { // nó da lista encadeada (linked list)
    v int      // elemento do nó
    next *Node // ponteiro para o próximo nó
}

type LinkedList struct { // estrutura de dados do tipo abstrato de dados: lista
    head *Node
    inserted int
}

func (l *LinkedList) Size() int { // O(1) Omega(1)
    return l.inserted // retorna a quantidade de elementos na lista
}

func (l *LinkedList) Get(index int) (int, error){ // O(n) Omega(1)
    if index>=0 && index < l.inserted {    // proteção para index fora dos limites da lista
        aux := l.head                      // auxiliar para percorrer a lista, sempre começa pela cabeça e vai até o index

        for i := 0; i < index; i++ {       // percorre a lista até o index, por isso a complexidade O(n)
            aux = aux.next                 // recebe o próximo nó
        }
        return aux.v, nil

    } else {
        return index, errors.New("Index fora dos limites da lista") 
    }
}

func (l *LinkedList) AddOnIndex(e int, index int) error { // O(n) Omega(1)
    if index >= 0 && index <= l.inserted {
        newNode := &Node{v:e}          // instancia um novo nó
        if l.head == nil {             // se a lista for vazia
            l.head = newNode           // a cabeça recebe o novo nó
        } else {                       // se não for vazia
            if index == 0 {            // se for uma inserção no início
                newNode.next = l.head  // o próximo do novo nó é a cabeça
                l.head = newNode       // a cabeça é o novo nó
            } else {                   // se a inserção não for no início
                aux := l.head          // auxiliar para percorrer a lista
                for i := 0; i < index - 1; i++ { // percorre a lista até o index - 1, uma posição antes do index
                    aux = aux.next     // recebe o próximo nó
                }
                newNode.next = aux.next // o próximo do novo nó é o próximo do auxiliar, guardando a referência
                aux.next = newNode      // o próximo do auxiliar é o novo nó, inserindo o novo nó
            }
        }
        l.inserted++
        return nil
    } else {
        return errors.New("Index invalido") 
    }
}

func (l *LinkedList) Add(e int) { // O(n) Omega(1)
    l.AddOnIndex(e, l.inserted)   // adiciona o elemento no final da lista
}

func (l *LinkedList) Remove(index int) error { // O(n) Omega(1)
    if index >= 0 && index < l.inserted {
        if index == 0 {                       // se a remoção for no início
            l.head = l.head.next              // a cabeça recebe o próximo
        } else {                              // se não for no início
            aux := l.head                     // auxiliar para percorrer a lista
            for i := 0; i < index - 1; i++ {  // percorre a lista até o index - 1, uma posição antes do index
                aux = aux.next                // recebe o próximo nó
            }
            aux.next = aux.next.next          // o próximo do auxiliar é o próximo do próximo, removendo o elemento
        }

        l.inserted--
        return nil
    } else {
        return errors.New("Index invalido")
    } 
}
*/
// main
func main(){
    l := &LinkedList{}
    //l.Init(10)
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
    l.Remove(9)
    for i:=0; i<10; i++ {
        val, _ := l.Get(i)
        fmt.Println(val)
    }
    //l.AddOnIndex(-1,0)
}