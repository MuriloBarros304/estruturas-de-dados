package main
import (
    "fmt"
)

func LinearSearch(v []int, e int) int { // O(n) Omega(1)
    for i := 0; i < len(v); i++ { // percorre o vetor até encontrar o elemento
        if v[i] == e {            // se encontrar o elemento
            return i              // retorna a posição do elemento
        }
    }
    return -1                     // se não encontrar o elemento, retorna -1
}

func BinarySearch(v []int, e int) int { // O(log n) Omega(1)
    ini := 0                    // início do vetor
    fim := len(v) - 1           // fim do vetor
    meio := (ini+fim) / 2       // meio do vetor, inicialmente
    for ini <= fim {            // enquanto o início for menor ou igual ao fim
        if v[meio] == e {       // se o elemento do meio for igual ao elemento procurado
            return meio         // retorna a posição do elemento
        } else if e < v[meio] { // se o elemento procurado for menor que o elemento do meio
            fim = meio - 1      // o fim passa a ser o meio - 1, pois o elemento procurado está na primeira metade
        } else {                // se o elemento procurado for maior que o elemento do meio
            ini = meio + 1      // o início passa a ser o meio + 1, pois o elemento procurado está na segunda metade
        }
        meio = (ini+fim)/2      // atualiza o meio a cada iteração
    }
    return -1                   // se não encontrar o elemento, retorna -1
}

func BinarySearchRec(v []int, e int, ini int, fim int) int { // O(log n) Omega(1)
    if ini > fim {          // se o início for maior que o fim
        return -1           // retorna -1, pois o elemento não foi encontrado
    }
    meio := (ini+fim) / 2   // calcula o meio
    if v[meio] == e {       // se o elemento do meio for igual ao elemento procurado
        return meio         // retorna a posição do elemento
    } else if e < v[meio] { // se o elemento procurado for menor que o elemento do meio
        fim = meio - 1      // o fim passa a ser o meio - 1, pois o elemento procurado está na primeira metade
    } else {                // se o elemento procurado for maior que o elemento do meio
        ini = meio + 1      // o início passa a ser o meio + 1, pois o elemento procurado está na segunda metade
    }
    return BinarySearchRec(v, e, ini, fim) // chama a função recursivamente
}

func main() {
    l := make([]int, 20)

    for i := 0; i < 20; i++ {
        l[i] = i
    }
    val := BinarySearch(l, 15)
    fmt.Println(val)
}
