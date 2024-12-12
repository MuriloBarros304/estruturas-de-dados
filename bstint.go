package main

import (
    "fmt"
)

type BSTNode struct {
    left *BSTNode
    val int
    right *BSTNode
}

func (node *BSTNode) createNode(val int) *BSTNode {
    return &BSTNode{val: val}
}

func (node *BSTNode) Add(val int) {
    if val < node.val {
        if node.left != nil{
            node.left.Add(val) 
        } else {
            node.left = node.createNode(val)
        }
    } else {
        if node.right != nil{
            node.right.Add(val)
        } else {
            node.right = node.createNode(val)
        }
    }
}

func (node *BSTNode) Search(val int) bool {
    if val == node.val {
        return true
    } else if val < node.val {
        if node.left != nil{
            return node.left.Search(val) 
        } else {
            return false
        }
    } else {
        if node.right != nil{
            return node.right.Search(val)
        } else {
            return false
        }
    }
}

func (node *BSTNode) Min() int {
    if node.left != nil {
        return node.left.Min()
    } else {
        return node.val
    }
}

func (node *BSTNode) Max() int {
    if node.right != nil {
        return node.right.Max()
    } else {
        return node.val
    }
}

func (node *BSTNode) Height() int {
    hl := 0
    hr := 0
    if node.left == nil && node.right == nil {
        return 0
    }
    if node.left != nil {
        hl = node.left.Height()
    }
    if node.right != nil {
        hr = node.right.Height()
    }
    if hl > hr {
        return hl + 1
    }
    return hr + 1
}

func (raiz *BSTNode) PreOrderNav() {
    fmt.Println(raiz.val)
    if raiz.left != nil {
        raiz.left.PreOrderNav()
    }
    if raiz.right != nil {
        raiz.right.PreOrderNav()
    }
}

func (raiz *BSTNode) InOrderNav() {
    if raiz != nil {
        raiz.left.InOrderNav()
    }
    fmt.Println(raiz.val)
    if raiz.right != nil {
        raiz.right.InOrderNav()
    }
}

func (raiz *BSTNode) PostOrderNav() {
    if raiz.left != nil {
        raiz.left.PostOrderNav()
    }
    if raiz.right != nil {
        raiz.right.PostOrderNav()
    }
    fmt.Println(raiz.val)
}

func (node *BSTNode) Remove(val int) *BSTNode {
    if val < node.val {
        node.left = node.left.Remove(val)          // reatribui o nó esquerdo, para que o return nil seja passado para o nó pai
    } else if val > node.val {
        node.right = node.right.Remove(val)
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
    bst := createNode(10)
    bst.Add(5)
    bst.Add(20)
    bst.Add(15)
    bst.Add(16)
    bst.InOrderNav()
    fmt.Println(bst.Search(16))
    fmt.Println(bst.Search(1))
    fmt.Println(bst.Min())
    fmt.Println(bst.Max())
    fmt.Println(bst.Height())
}