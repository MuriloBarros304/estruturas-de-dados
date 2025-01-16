package main
import (
    "fmt"
)

type BSTNode struct {
    left *BSTNode
    val int
    right *BSTNode
    height int
    bf int
}

func (root *BSTNode) createNode(val int) *BSTNode { // O(1)
    return &BSTNode{val: val}
}

func (root *BSTNode) Add(val int) *BSTNode {
    if val <= root.val {
        if root.left != nil {
            root.left = root.left.Add(val)
            root.left = root.createNode(val)
        }
    } else {
        if root.right != nil {
            root.right = root.right.Add(val)
            root.right = root.createNode(val)
        }
    }
    root.UpdateProperties()
    return root.Rebalance()
}

func (root *BSTNode) Remove(val int) *BSTNode {
    if root.val == val {
        if root.left == nil && root.right == nil {
            //caso1
            return nil
        } else if root.left != nil && root.right == nil {
            //caso2: esq
            return root.left
        } else if root.left == nil && root.right != nil {
            //caso2: dir
            return root.right
        } else {
            maxLeft := root.left.max()
            root.val = maxLeft
            root.left.Remove(maxLeft)
        }
    } else if val < root.val {
        if root.left != nil {
            root.left = root.left.Remove(val)
        }
    } else {
        if root.right != nil {
            root.right = root.right.Remove(val)
        }
    }
    root.UpdateProperties()
    return root.Rebalance()
}


func (root *BSTNode) RotRight() *BSTNode {
    newRoot := root.left
    root.left = newRoot.right
    newRoot.right = root
    root.UpdateProperties()
    newRoot.UpdateProperties()
    return newRoot
}
func (root *BSTNode) RotLeft() *BSTNode {
    newRoot := root.right
    root.right = newRoot.left
    newRoot.left = root
    root.UpdateProperties()
    newRoot.UpdateProperties()
    return newRoot
}

func (root *BSTNode) UpdateProperties() {
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
    } else if hl > hr {           // altura do filho esquerdo é maior
            root.height = hl + 1
    } else {
            root.height = hr + 1
    }
}

//funcao para rebalancear caso 1
func (root *BSTNode) RebalanceLeftLeft() *BSTNode {
    return root.RotRight()
}

//funcao para rebalancear caso 2
func (root *BSTNode) RebalanceLeftRight() *BSTNode {
    root.left = root.left.RotLeft()
    return root.RotRight()
}

//funcao para rebalancear caso 3
func (root *BSTNode) RebalanceRightRight() *BSTNode {
    return root.RotLeft()
}

//funcao para rebalancear caso 4
func (root *BSTNode) RebalanceRightLeft() *BSTNode {
    root.right = root.right.RotRight()
    return root.RotLeft()
}

//funcao para rebalancear um no
func (root *BSTNode) Rebalance() *BSTNode {
    if root.bf == -2 { //desbalanceada à esquerda
            if root.left.bf == -1 || root.left.bf == 0  { // esq-esq ou esq-neutro
                    return root.RebalanceLeftLeft()
            } else if root.left.bf == 1 { // esq-dir
                    return root.RebalanceLeftRight()
            }
    } else if root.bf == 2 { //desbalanceada à direita
            if root.right.bf == -1 || root.right.bf == 0  { // dir-dir ou dir-neutro
                    return root.RebalanceRightRight()
            } else if root.right.bf == 1 { //dir-esq
                    return root.RebalanceRightLeft()
            }
    }
    //return root.Rebalance()
}

func (root *BSTNode) max() int {
    return root.right.max()
}

func main() {
    root := (&BSTNode{}).createNode(10)
    //root := createNode(10)
    root = root.Add(5)
    root = root.Add(15)
    root = root.Add(3)
    root = root.Add(7)
    root = root.Add(13)
    root = root.Add(17)
    root = root.Add(2)
    root = root.Add(4)
    root = root.Add(6)
    root = root.Add(8)
    root = root.Add(12)
    root = root.Add(14)
    root = root.Add(16)
    root = root.Add(18)
    root = root.Add(1)
    root = root.Add(9)
    root = root.Add(11)
    root = root.Add(19)
/*     root = root.Remove(1)
    root = root.Remove(2)
    root = root.Remove(3)
    root = root.Remove(4)
    root = root.Remove(5)
    root = root.Remove(6)
    root = root.Remove(7)
    root = root.Remove(8)
    root = root.Remove(9)
    root = root.Remove(10)
    root = root.Remove(11)
    root = root.Remove(12)
    root = root.Remove(13)
    root = root.Remove(14)
    root = root.Remove(15)
    root = root.Remove(16)
    root = root.Remove(17)
    root = root.Remove(18)
    root = root.Remove(19) */
    fmt.Println(root)
}

