package main
import (
    "fmt"
)

func LinearSearch(v []int, e int) int {
    for i := 0; i < len(v); i++ {
        if v[i] == e {
            return i
        }
    }
    return -1
}

func BinarySearch(v []int, e int) int {
    ini := 0
    fim := len(v) - 1
    meio := (ini+fim) / 2
    for ini <= fim {
        if v[meio] == e {
            return meio
        } else if e < v[meio] {
            fim = meio - 1
        } else {
            ini = meio + 1
        }
        meio = (ini+fim)/2
    }
    return -1
}

func BinarySearchRec(v []int, e int, ini int, fim int) int {
    if ini > fim {
        return -1
    }
    meio := (ini+fim) / 2
    if v[meio] == e {
        return meio
    } else if e < v[meio] {
        fim = meio - 1
    } else {
        ini = meio + 1
    }
    return BinarySearchRec(v, e, ini, fim)
}

func main() {
    l := make([]int, 20)

    for i := 0; i < 20; i++ {
        l[i] = i
    }
    val := BinarySearch(l, 15)
    fmt.Println(val)
}
