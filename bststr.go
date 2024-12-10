package main

import (
    "fmt"
)

type BSTNode struct {
    left *BSTNode
    val string
    right *BSTNode
}

func (node *BSTNode) createNode(val string) *BSTNode {
    return &BSTNode{val: val}
}

func (node *BSTNode) Add(val string) {
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

func (node *BSTNode) Search(val string) bool {
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

func (node *BSTNode) Min() string {
    if node.left != nil {
        return node.left.Min()
    } else {
        return node.val
    }
}

func (node *BSTNode) Max() string {
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

func main() {
    raiz := createNode("L")
    raiz.Add("D")
    raiz.Add("Q")
    raiz.Add("B")
    raiz.Add("G")
    raiz.Add("N")
    raiz.Add("G")
    fmt.Println("PreOrderNav")
    raiz.PreOrderNav()
    fmt.Println("InOrderNav")
    raiz.InOrderNav()
    fmt.Println("PostOrderNav")
    raiz.PostOrderNav()
}