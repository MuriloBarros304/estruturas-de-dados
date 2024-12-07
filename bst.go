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

func (node *BSTNode) InOrderNav() {
    if node.left != nil {
        node.left.InOrderNav()
    }
    if node.right != nil {
        node.right.InOrderNav()
    }
    fmt.Println(node.val)
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

func main() {
    bst := (&BSTNode{}).createNode(10)
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