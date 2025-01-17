package main

import (
    "errors"
    "fmt"
)

// Map interface define as operações de um mapa
type Map interface {
    put(key int, value string)
    get(key int) (string, error)
    remove(key int)
    size() int
    loadFactor() float64
    Init()
}

// Tuple representa um par chave-valor
type Tuple struct {
    key   int
    value string
    next  *Tuple
}

// HashMap é a estrutura da tabela hash com encadeamento
type HashMap struct {
    buckets    []*Tuple
    numBuckets int
    numEntries int
}

// hashFunction calcula o índice para um dado key
func (hm *HashMap) hashFunction(key int) int {
    return key % hm.numBuckets
}

// Init inicializa a tabela hash
func (hm *HashMap) Init() {
    hm.numBuckets = 10
    hm.buckets = make([]*Tuple, hm.numBuckets)
    hm.numEntries = 0
}

// put adiciona ou atualiza um par chave-valor
func (hm *HashMap) put(key int, value string) {
    index := hm.hashFunction(key)
    head := hm.buckets[index]

    // Verifica se a chave já existe
    current := head
    for current != nil {
        if current.key == key {
            current.value = value // Atualiza o valor
            return
        }
        current = current.next
    }

    // Insere um novo par chave-valor no início da lista ligada
    newNode := &Tuple{key: key, value: value, next: head}
    hm.buckets[index] = newNode
    hm.numEntries++
}

// get recupera o valor associado a uma chave
func (hm *HashMap) get(key int) (string, error) {
    index := hm.hashFunction(key)
    current := hm.buckets[index]

    // Percorre a lista ligada no bucket
    for current != nil {
        if current.key == key {
            return current.value, nil
        }
        current = current.next
    }
    return "", errors.New("chave não encontrada")
}

// remove exclui uma chave da tabela hash
func (hm *HashMap) remove(key int) {
    index := hm.hashFunction(key)
    current := hm.buckets[index]
    var prev *Tuple

    // Procura a chave na lista ligada
    for current != nil {
        if current.key == key {
            if prev == nil {
                // Remove o nó do início da lista
                hm.buckets[index] = current.next
            } else {
                // Remove o nó intermediário ou final
                prev.next = current.next
            }
            hm.numEntries--
            return
        }
        prev = current
        current = current.next
    }
}

// size retorna o número de elementos na tabela
func (hm *HashMap) size() int {
    return hm.numEntries
}

// loadFactor calcula o fator de carga da tabela
func (hm *HashMap) loadFactor() float64 {
    return float64(hm.numEntries) / float64(hm.numBuckets)
}

// Exibe a tabela hash
func (hm *HashMap) display() {
    for i, bucket := range hm.buckets {
        fmt.Printf("Bucket %d: ", i)
        current := bucket
        for current != nil {
            fmt.Printf("[%d: %s] -> ", current.key, current.value)
            current = current.next
        }
        fmt.Println("nil")
    }
}

func main() {
    hm := &HashMap{}
    hm.Init()

    fmt.Println("Inserindo elementos:")
    hm.put(1, "um")
    hm.put(2, "dois")
    hm.put(12, "doze") // Colisão com 2
    hm.put(22, "vinte e dois") // Colisão com 2
    hm.display()

    fmt.Println("\nBuscando elementos:")
    value, err := hm.get(2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Chave 2: %s\n", value)
    }

    value, err = hm.get(22)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Chave 22: %s\n", value)
    }

    fmt.Println("\nRemovendo elemento com chave 2:")
    hm.remove(2)
    hm.display()

    fmt.Printf("\nTamanho da tabela: %d\n", hm.size())
    fmt.Printf("Fator de carga: %.2f\n", hm.loadFactor())
}
