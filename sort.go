package main
import (
    "fmt"
    "math"
    "math/rand"
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

// MergeSort ----------------------------------------------
// Função para unir dois vetores ordenados em um vetor
// ordenado.
func merge(v []int , e []int, d []int) { // O(n) Omega(n)
    indexE, indexD, indexV := 0, 0, 0        // índices para percorrer os vetores esquerdo, direito e o vetor final
    for indexE < len(e) && indexD < len(d) { // enquanto estiverem dentro dos limites
        if e[indexE] < d[indexD] {           // compara os elementos dos vetores da esquerda e direita (se o da esquerda for menor)
            v[indexV] = e[indexE]            // coloca o elemento da esquerda no vetor v
            indexE++                         // incrementa o índice da esquerda
        } else {                             // se o elemento do vetor da direita for menor
            v[indexV] = d[indexD]            // coloca o elemento da direita no vetor v
            indexD++                         // incrementa o índice da direita
        }
        indexV++                             // incrementa o índice do vetor final, isso se repete até um dos vetores acabar
    }
    for indexD < len(d) {                    // para o caso de um dos vetores ter elementos restantes
        v[indexV] = d[indexD]                // coloca os elementos restantes no vetor v
        indexD++
        indexV++
    }
    for indexE < len(e) {
        v[indexV] = e[indexE]
        indexE++
        indexV++
    }
}

// Função principal do MergeSort, que divide o vetor em
// dois vetores e chama a função Merge para unir os vetores
// ordenados. Geralmente é implementado de forma recursiva.
func MergeSort(v []int) { // O(n log n) Omega(n log n), n para cada nível da recursão, log n para a quantidade de níveis
    if len(v) > 1 {             // se o vetor tiver mais de um elemento
        tamE := len(v) / 2      // vetor da esquerda é a primeira metade
        tamD := len(v) - tamE   // restante do vetor
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
        merge(v, e, d)
    }
}

// QuickSort ----------------------------------------------
// Função para particionar o vetor em relação a um pivô,
// colocando os elementos menores à esquerda e os maiores
// à direita. A função retorna a posição do pivô. O pivô
// é escolhido aleatoriamente.
func partition(v []int, ini int, fim int) int { // O(n^2) Omega(n log n)
    random := rand.Intn(fim-ini) + ini        // sorteia valor entre ini e fim
    v[fim], v[random] = v[random], v[fim]     // trocar rand de posição com fim
    indexPivot := fim                         // pega o último elemento como pivô, que foi sorteado
    pIndex := ini                             // índice de troca
    for i := ini; i < indexPivot; i++ {       // percorre o vetor
        if v[i] <= v[indexPivot] {            // se o elemento for menor ou igual ao pivô
            v[i], v[pIndex] = v[pIndex], v[i] // troca os elementos, colocando os menores à esquerda do pivô e os maiores à direita
            pIndex++                          // incrementa o índice de troca
        }
    }
    v[pIndex], v[indexPivot] = v[indexPivot], v[pIndex] // coloca o pivô na posição correta
    return pIndex                                       // retorna a posição do pivô
}

// Função principal do QuickSort, que chama a função
// partition para particionar o vetor e chama a função
// recursivamente para as partições da esquerda e da direita.
func QuickSort(v []int, ini int, fim int) { // O(n^2) Omega(n log n)
    if ini < fim {
        indexPivot := partition(v, ini, fim) // chama a função partition
        QuickSort(v, ini, indexPivot-1)      // chama a função recursivamente para a partição da esquerda
        QuickSort(v, indexPivot+1, fim)      // chama a função recursivamente para a partição da direita
    }
}

// CountingSort -------------------------------------------
// Algoritmo para ordenar um vetor de inteiros não negativos
// em ordem crescente. A função conta a quantidade de
// elementos e faz uma soma cumulativa. É estável,
// ou seja, elementos iguais mantêm a ordem relativa.
// Geralmente é utilizado para ordenar inteiros pequenos.
func CountingSort(v []int) { // O(n+k) Omega(n), k = max - min
    n := len(v)
    max := 0
    for i := 0; i < n; i++ {    // encontra o maior elemento
        if v[i] > max {
            max = v[i]
        }
    }
    count := make([]int, max+1) // vetor para contar a quantidade de elementos
    for i := 0; i < n; i++ {
        count[v[i]]++
    }
    for i := 1; i <= max; i++ { // soma cumulativa
        count[i] += count[i-1]
    }
    sorted := make([]int, n)    // vetor ordenado
    for i := n-1; i >= 0; i-- {
        sorted[count[v[i]]-1] = v[i] // shift para a esquerda
        count[v[i]]--
    }
    for i := 0; i < n; i++ {
        v[i] = sorted[i]
    }
}

func main() {
    v := []int{8, 4, 9, 5, 3, 2, 6, 1}
    fmt.Println("Desordenado: ", v)

    v = SelectionSortOP(v)
    fmt.Println("SelectionSortOP: ", v)

    v = []int{8, 4, 9, 5, 3, 2, 6, 1}
    SelectionSortIP(v)
    fmt.Println("SelectionSortIP: ", v)

    v = []int{8, 4, 9, 5, 3, 2, 6, 1}
    BubbleSort(v)
    fmt.Println("BubbleSort: ", v)

    v = []int{8, 4, 9, 5, 3, 2, 6, 1}
    InsertionSort(v)
    fmt.Println("InsertionSort: ", v)

    v = []int{8, 4, 9, 5, 3, 2, 6, 1}
    MergeSort(v)
    fmt.Println("MergeSort: ", v)

    v = []int{8, 4, 9, 5, 3, 2, 6, 1}
    QuickSort(v, 0, len(v)-1)
    fmt.Println("QuickSort: ", v)

    v = []int{8, 4, 9, 5, 3, 2, 6, 1}
    CountingSort(v)
    fmt.Println("CountingSort: ", v)
}