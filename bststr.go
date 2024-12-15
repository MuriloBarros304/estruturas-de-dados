package main

import (
    "fmt"
)

type BSTNode struct {
    left *BSTNode
    val string
    right *BSTNode
}

func (node *BSTNode) createNode(val string) *BSTNode { // O(1)
    return &BSTNode{val: val}
}

func (node *BSTNode) Add(val string) { // O(n), onde n é a altura da árvore
    if val < node.val {        // se o valor que será adicionado for menor que o valor do nó atual
        if node.left != nil{   // se o nó à esquerda não for nulo
            node.left.Add(val) // chama a função Add para o nó à esquerda
        } else {               // se o nó à esquerda for nulo
            node.left = node.createNode(val) // cria um novo nó à esquerda
        }
    } else {                    // se o valor que será adicionado for maior que o valor do nó atual
        if node.right != nil{   // se o nó à direita não for nulo
            node.right.Add(val) // chama a função Add para o nó à direita
        } else {                // se o nó à direita for nulo
            node.right = node.createNode(val) // cria um novo nó à direita
        }
    }
}

func (node *BSTNode) Search(val string) bool { // O(n), onde n é a altura da árvore
    if val == node.val {                  // se o valor que está sendo procurado for igual ao valor do nó atual
        return true
    } else if val < node.val {            // se o valor que está sendo procurado for menor que o valor do nó atual
        if node.left != nil{              // se o nó à esquerda não for nulo
            return node.left.Search(val)  // chama a função Search para o nó à esquerda
        } else {                          // não encontrou o valor
            return false
        }
    } else {                              // se o valor que está sendo procurado for maior que o valor do nó atual
        if node.right != nil{             // se o nó à direita não for nulo
            return node.right.Search(val) // chama a função Search para o nó à direita
        } else {                          // não encontrou o valor
            return false
        }
    }
}

func (node *BSTNode) Min() string { // O(n), onde n é a altura da árvore
    if node.left != nil {      // se o nó à esquerda não for nulo
        return node.left.Min() // chama a função Min para o nó à esquerda
    } else {                   // se o nó à esquerda for nulo
        return node.val        // retorna o valor do nó atual, que é o mínimo
    }
}

func (node *BSTNode) Max() string { // O(n), onde n é a altura da árvore
    if node.right != nil {      // se o nó à direita não for nulo
        return node.right.Max() // chama a função Max para o nó à direita
    } else {                    // se o nó à direita for nulo
        return node.val         // retorna o valor do nó atual, que é o máximo
    }
}

func (node *BSTNode) Height() int { // O(n), onde n é a altura da árvore
    hl := 0 // altura do nó à esquerda
    hr := 0 // altura do nó à direita
    if node.left == nil && node.right == nil { // se o nó for folha, a altura é 0
        return 0
    }
    if node.left != nil {              // se o nó à esquerda não for nulo
        hl = node.left.Height()        // chama a função Height para o nó à esquerda recursivamente
    }
    if node.right != nil {             // se o nó à direita não for nulo
        hr = node.right.Height()       // chama a função Height para o nó à direita recursivamente
    }
    if hl > hr {                       // retorna a maior altura entre o nó à esquerda e o nó à direita
        return hl + 1                  // adiciona 1 para contar o nó atual, a altura incrementa
    }                                  // recursivamente, retornando a maior entre os dois.
    return hr + 1
}

func (raiz *BSTNode) PreOrderNav() { // O(n)
    fmt.Println(raiz.val)       // pré-ordem: raiz, esquerda, direita
    if raiz.left != nil {       // se tiver nó à esquerda
        raiz.left.PreOrderNav()
    }
    if raiz.right != nil {      // se tiver nó à direita
        raiz.right.PreOrderNav()
    }
}

func (raiz *BSTNode) InOrderNav() { // O(n)
    if raiz.left != nil {             // em ordem: esquerda, raiz, direita
        raiz.left.InOrderNav()   // se tiver nó à esquerda
    }
    fmt.Println(raiz.val)        // imprime o valor do nó atual
    if raiz.right != nil {       // se tiver nó à direita
        raiz.right.InOrderNav()
    }
}

func (raiz *BSTNode) PostOrderNav() { // O(n)
    if raiz.left != nil {        // pós-ordem: esquerda, direita, raiz
        raiz.left.PostOrderNav() // se tiver nó à esquerda
    }
    if raiz.right != nil {       // se tiver nó à direita
        raiz.right.PostOrderNav()
    }
    fmt.Println(raiz.val)        // imprime o valor do nó atual
}

func (node *BSTNode) Remove(val string) *BSTNode {
    if val < node.val {                            // se o valor que será removido for menor que o valor do nó atual
        node.left = node.left.Remove(val)          // reatribui o nó esquerdo, para que o return nil seja passado para o nó pai
    } else if val > node.val {                     // se o valor que será removido for maior que o valor do nó atual
        node.right = node.right.Remove(val)        // reatribui o nó direito, para que o return nil seja passado para o nó pai
    } else {
        if node.left == nil && node.right == nil { // caso 1: nó folha
            return nil
        } else if node.left == nil {               // caso 2: nó com um filho à direita
            return node.right
        } else if node.right == nil {              // caso 2: nó com um filho à esquerda
            return node.left
        } else {                                   // caso 3: nó com dois filhos
            min := node.right.Min()                // abordagem 1: encontrar o menor valor do nó à direita
            node.val = min
            node.right = node.right.Remove(min)    // remove o nó com o menor valor
        }
    }
    return node
}

func main() {

    raiz := BSTNode{nil, "L", nil}
    raiz.Add("D")
    raiz.Add("Q")
    raiz.Add("B")
    raiz.Add("G")
    raiz.Add("N")
    fmt.Println("PreOrderNav")
    raiz.PreOrderNav()
    fmt.Println("InOrderNav")
    raiz.InOrderNav()
    fmt.Println("PostOrderNav")
    raiz.PostOrderNav()

    /*
            L
          /    \  
        D        Q
      /   \    /  
    B      G   N
    */
}