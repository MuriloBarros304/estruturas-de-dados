package main
import (
    "fmt"
    "math"
)
func BubbleSort(v []int) { // O(n^2) Omega(n)
    for varredura := 0; varredura < len(v)-1; varredura++ { // percorre o vetor
        trocou := false                       // flag para verificar se houve troca
        for i := 0; i < len(v)-1-varredura; i++ { 
            if v[i] > v[i+1] {
                v[i], v[i+1] = v[i+1], v[i]
                trocou = true
            }
        }
        if !trocou {
            return
        }
    }
}

func SelectionSortOP(v []int) []int { // O(n^2) Omega(n^2)
    ordenado := make([]int, len(v))                                                                                            
    for varredura := 1; varredura <= len(v); varredura++ {                       
        iMenor := varredura - 1                                                   
        for i:=0; i < len(v); i++{                                                    
            if v[i] < v[iMenor] {  
                iMenor = i
            }
        }
    ordenado[varredura-1] = v[iMenor]
    v[iMenor] = math.MaxInt
    }
    return ordenado
}


func SelectionSortIP(v []int) { // O(n^2) Omega(n^2)
    for varredura := 0; varredura < len(v) - 1; varredura++ { // percorre o vetor                 
        iMenor := varredura                                   // guarda a posição do menor elemento                
        for i:=varredura+1; i < len(v); i++{                  // percorre o vetor a partir da posição varredura                                  
            if v[i] < v[iMenor] {  
                iMenor = i
            }
        }
        v[iMenor],v[varredura] = v[varredura],v[iMenor]
    }
}


func InsertionSort(v []int) { // O(n^2) Omega(n)
    for insercao := 1; insercao < len(v); insercao++ { // percorre o vetor
        for i := insercao; i >= 1 && v[i] < v[i-1]; i-- { // compara o elemento atual com o anterior
            v[i], v[i-1] = v[i-1], v[i] // troca os elementos
        }
    }
}

func main() {
    l := []int{5, 3, 4, 1, 2}
    fmt.Println(l)
    sorted := SelectionSortOP(l)
    fmt.Println(sorted)
    //SelectionSortIP(l)
    //BubbleSort(l)
    //InsertionSort(l)
    //fmt.Println(l)
}