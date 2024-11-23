package main
import (
    "fmt"
    "math"
)

// BubbleSort ----------------------------------------------
// As comparações são feitas em pares, aos poucos os maiores
// elementos vão para o final do vetor. A função para quando
// não houver mais trocas. O melhor caso é quando o vetor já
// está ordenado, pois a função para em uma varredura. O pior
// caso é quando o vetor está ordenado de forma decrescente,
// pois a função faz n-1 varreduras e n-1 comparações em cada
// varredura. 
func BubbleSort(v []int) { // O(n^2) Omega(n)
    for varredura := 0; varredura < len(v)-1; varredura++ { // percorre o vetor
        trocou := false                                     // flag para verificar se houve troca
        for i := 0; i < len(v)-1-varredura; i++ {           // percorre a partição não ordenada
            if v[i] > v[i+1] {                              // compara com o seguinte
                v[i], v[i+1] = v[i+1], v[i]                 // troca os elementos
                trocou = true                               // sinaliza que houve troca
            }
        }
        if !trocou {                                        // se não houve troca, o vetor está ordenado
            return
        }
    }
}

// SelectionSort ------------------------------------------
// A cada varredura, o menor elemento é colocado na posição
// correta, ou seja, na primeira posição da partição ordenada.
// No out of place, o vetor ordenado é criado e o menor 
func SelectionSortOP(v []int) []int { // O(n^2) Omega(n^2)
    ordenado := make([]int, len(v))                        // vetor ordenado
    for varredura := 1; varredura <= len(v); varredura++ { // percorre o vetor até a última posição
        iMenor := varredura - 1                            // guarda a posição do menor elemento, inicialmente a varredura        
        for i:=0; i < len(v); i++{                         // percorre o vetor 
            if v[i] < v[iMenor] {                          // se o elemento atual for menor que o menor já armazenado
                iMenor = i                                 // guarda a posição do menor elemento
            }
        }
        ordenado[varredura-1] = v[iMenor]                  // coloca o menor elemento na posição correta
        v[iMenor] = math.MaxInt                            // marca o menor elemento como o maior inteiro, para não ser selecionado novamente
    }
    return ordenado                                        // retorna o vetor ordenado
}

// Na versão in place, o menor elemento é colocado na primeira
// posição da partição ordenada, sem a necessidade de um vetor
// auxiliar. A partição desordenada é reduzida a cada varredura.
// Este algoritmo de ordenação é um dos piores, pois faz muitas
// varreduras e trocas.
func SelectionSortIP(v []int) { // O(n^2) Omega(n^2) a versão in place é mais eficiente
    for varredura := 0; varredura < len(v) - 1; varredura++ { // percorre o vetor até a penúltima posição              
        iMenor := varredura                                   // guarda a posição do menor elemento                
        for i:=varredura+1; i < len(v); i++{                  // percorre o vetor a partir da posição varredura                                  
            if v[i] < v[iMenor] {                             // se o elemento atual for menor que o menor já armazenado
                iMenor = i                                    // guarda a posição do menor elemento
            }
        }
        v[iMenor],v[varredura] = v[varredura],v[iMenor]       // coloca o menor elemento na posição correta
    }
}

// InsertionSort ------------------------------------------
// A cada varredura, o elemento é comparado com o anterior
// e trocado se necessário. O melhor caso é quando o vetor
// já está ordenado, pois a função para em uma varredura.
// O pior caso é quando o vetor está ordenado de forma
// decrescente, pois a função faz n-1 varreduras e n-1
// comparações em cada varredura.
func InsertionSort(v []int) { // O(n^2) Omega(n)
    for insercao := 1; insercao < len(v); insercao++ {    // percorre o vetor
        for i := insercao; i >= 1 && v[i] < v[i-1]; i-- { // compara o elemento atual com o anterior
            v[i], v[i-1] = v[i-1], v[i]                   // troca os elementos
        }
    }
}

func Merge(v []int , e []int, d []int) { // O(n) Omega(n)
    indexE, indexD, indexV := 0, 0, 0
    for indexE < len(e) && indexD < len(d) {
        if e[indexE] < d[indexD] {
            v[indexV] = e[indexE]
            indexE++
        } else {
            v[indexV] = d[indexD]
            indexD++
        }
        indexV++
    }
    for indexD < len(d) {
        v[indexV] = d[indexD]
        indexD++
        indexV++
    }
    for indexE < len(e) {
        v[indexV] = e[indexE]
        indexE++
        indexV++
    }
}

func MergeSort(v []int) { // O(n log n) Omega(n log n), n para cada nível da recursão, log n para a quantidade de níveis
    if len(v) <= 1 {
        tamE := len(v)/2
        tamD := len(v) - tamE // restante do vetor
        e := make([]int, tamE)
        for i := 0; i < tamE; i++ {
            e[i] = v[i]
        }
        d := make([]int, tamD)
        for j := tamE; j < len(v); j++ {
            d[j-tamE] = v[j]    // preenche o vetor da direita
        }
        MergeSort(e)
        MergeSort(d)
        Merge(v, e, d)
    }
}

func QuickSort(v []int, ini int, fim int) { // O(n^2) Omega(n log n)
    if ini < fim {
        iPivot := partition(v, ini, fim)
        QuickSort(v, ini, iPivot-1)
        QuickSort(v, iPivot+1, fim)
    }
}

func main() {
    //v := make([]int, 100)
    l := []int{5, 3, 4, 1, 2}
    /* fmt.Println(l)
    MergeSort(l)
    fmt.Println(l) */
    QuickSort(l)
    //SelectionSortIP(l)
    //MergeSort(l)
    //BubbleSort(l)
    //InsertionSort(l)
    fmt.Println(l)
}