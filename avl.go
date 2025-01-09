package main
import (
    "fmt"
)

type BstNode struct {
    left *BstNode
    value int
    right *BstNode
    height int
    bf int
}

func (root *BstNode) RotRight() *BstNode {
    newRoot := root.left
    root.left = newRoot.right
    newRoot.right = root
    root.UpdateProperties()
    newRoot.UpdateProperties()
    return newRoot
}

func (bstNode *BstNode) RotLeft() *BstNode {
    newRoot := root.right
    root.right = newRoot.left
    newRoot.left = root
    root.UpdateProperties()
    newRoot.UpdateProperties()
    return newRoot
}

func (root *BstNode) UpdateProperties() {
    hl := 0
    hr := 0
    if root.left != nil {
            hl = root.left.height
    } 
    if root.right != nil {
            hr = root.right.height
    }
    root.bf = hr - hl
    if root.left == nil && root.right == nil {
            root.height = 0
    } else if hl > hr {
            root.height = hl + 1
    } else {
            root.height = hr + 1
    }
}

