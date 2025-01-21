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
    if root == nil { // Se o nó atual for nulo, crie um novo nó
        return &BSTNode{val: val}
    }
    if val <= root.val { // Adiciona no lado esquerdo caso o valor seja menor ou igual
        if root.left != nil { // Se o nó esquerdo não for nulo, adicione o valor ao nó esquerdo
            root.left = root.left.Add(val)
        } else {
            root.left = root.createNode(val) // Se o nó esquerdo for nulo, crie um novo nó
        }
    } else { // Adiciona no lado direito
        if root.right != nil { // Se o nó direito não for nulo, adicione o valor ao nó direito
            root.right = root.right.Add(val)
        } else {
            root.right = root.createNode(val) // Se o nó direito for nulo, crie um novo nó
        }
    }
    root.UpdateProperties() // Atualiza altura e fator de balanceamento
    return root.Rebalance() // Rebalanceia o nó, se necessário
}

func (root *BSTNode) Remove(val int) *BSTNode { // O(log n)
    if root.val == val {                           // Caso o valor seja a raiz
        if root.left == nil && root.right == nil { // Se a raiz não tiver filhos
            //caso1
            return nil                             // Retorna nulo, pois removemos a raiz
        } else if root.left != nil && root.right == nil { // Se a raiz tiver um filho à esquerda
            //caso2: esq
            return root.left                       // Retorna o filho à esquerda, que será a nova raiz
        } else if root.left == nil && root.right != nil { // Se a raiz tiver um filho à direita
            //caso2: dir
            return root.right                      // Retorna o filho à direita, que será a nova raiz
        } else {                                   // Se a raiz tiver dois filhos
            maxLeft := root.left.max()             // Encontra o maior valor à esquerda
            root.val = maxLeft
            root.left.Remove(maxLeft)
        }
    } else if val < root.val {                     // Caso o valor seja menor que a raiz
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
    if root.bf == -2 { // desbalanceada à esquerda
        if root.left.bf == -1 || root.left.bf == 0 { // esq-esq ou esq-neutro
            return root.RebalanceLeftLeft()
        } else if root.left.bf == 1 { // esq-dir
            return root.RebalanceLeftRight()
        }
    } else if root.bf == 2 { // desbalanceada à direita
        if root.right.bf == 1 || root.right.bf == 0 { // dir-dir ou dir-neutro
            return root.RebalanceRightRight()
        } else if root.right.bf == -1 { // dir-esq
            return root.RebalanceRightLeft()
        }
    }
    return root // Caso a árvore já esteja balanceada
}


func (root *BSTNode) max() int {
    return root.right.max()
}

// Função auxiliar para imprimir a árvore em ordem (in-order traversal)
func (root *BSTNode) InOrderTraversal() {
    if root == nil {
        return
    }
    root.left.InOrderTraversal()
    fmt.Printf("%d ", root.val)
    root.right.InOrderTraversal()
}

func main() {
    fmt.Println("=== Construindo a Árvore AVL ===")
    root := (&BSTNode{}).createNode(10)
    
    // Inserindo valores na árvore
    valuesToAdd := []int{5, 15, 3, 7, 13, 17, 2, 4, 6, 8, 12, 14, 16, 18, 1, 9, 11, 19}
    fmt.Println("Adicionando valores:", valuesToAdd)
    for _, val := range valuesToAdd {
        root = root.Add(val)
    }

    fmt.Println("\nÁrvore após as inserções (em ordem):")
    root.InOrderTraversal()
    fmt.Println()

    // Removendo alguns valores
    valuesToRemove := []int{1, 2, 3, 4, 5}
    fmt.Println("\nRemovendo valores:", valuesToRemove)
    for _, val := range valuesToRemove {
        root = root.Remove(val)
    }

    fmt.Println("\nÁrvore após as remoções (em ordem):")
    root.InOrderTraversal()
    fmt.Println()

    // Finalizando
    fmt.Println("\n=== Programa Finalizado ===")
}

