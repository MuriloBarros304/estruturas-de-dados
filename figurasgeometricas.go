package main

import "fmt"

// Definindo a interface FiguraGeometrica com os atributos básicos
type FiguraGeometrica interface {
    Nome() string
    NumeroDeVertices() int
    NumeroDeArestas() int
    PrintInfo()
}

// Implementação da classe concreta Circulo
type Circulo struct {
    raio float64
    cor string
}

// Método que retorna o nome da figura
func (c Circulo) Nome() string {
    return "Círculo"
}

// Método que retorna o número de vértices do círculo (0, pois é uma curva)
func (c Circulo) NumeroDeVertices() int {
    return 0
}

// Método que retorna o número de arestas do círculo (0, pois é uma curva)
func (c Circulo) NumeroDeArestas() int {
    return 0
}

// Método para imprimir as informações do círculo
func (c Circulo) PrintInfo() {
    fmt.Println("Figura:", c.Nome())
    fmt.Println("Cor:", c.cor)
    fmt.Println("Raio:", c.raio)
    fmt.Println("Número de vértices:", c.NumeroDeVertices())
    fmt.Println("Número de arestas:", c.NumeroDeArestas())
}

// Implementação da classe concreta Triangulo
type Triangulo struct {
    cor string
}

// Método que retorna o nome da figura
func (t Triangulo) Nome() string {
    return "Triângulo"
}

// Método que retorna o número de vértices do triângulo (3)
func (t Triangulo) NumeroDeVertices() int {
    return 3
}

// Método que retorna o número de arestas do triângulo (3)
func (t Triangulo) NumeroDeArestas() int {
    return 3
}

// Método para imprimir as informações do triângulo
func (t Triangulo) PrintInfo() {
    fmt.Println("Figura:", t.Nome())
    fmt.Println("Cor:", t.cor)
    fmt.Println("Número de vértices:", t.NumeroDeVertices())
    fmt.Println("Número de arestas:", t.NumeroDeArestas())
}

// Implementação da classe concreta Quadrado
type Quadrado struct {
    cor string
}

// Método que retorna o nome da figura
func (q Quadrado) Nome() string {
    return "Quadrado"
}

// Método que retorna o número de vértices do quadrado (4)
func (q Quadrado) NumeroDeVertices() int {
    return 4
}

// Método que retorna o número de arestas do quadrado (4)
func (q Quadrado) NumeroDeArestas() int {
    return 4
}

// Método para imprimir as informações do quadrado
func (q Quadrado) PrintInfo() {
    fmt.Println("Figura:", q.Nome())
    fmt.Println("Cor:", q.cor)
    fmt.Println("Número de vértices:", q.NumeroDeVertices())
    fmt.Println("Número de arestas:", q.NumeroDeArestas())
}

func main() {
    // Criando uma lista de figuras geométricas
    figuras := []FiguraGeometrica{
        Circulo{raio: 5.0, cor: "Vermelho"},
        Triangulo{cor: "Azul"},
        Quadrado{cor: "Verde"},
    }

    // Iterando sobre a lista e printando as informações de cada figura
    for _, figura := range figuras {
        figura.PrintInfo()
        fmt.Println()
    }
}
