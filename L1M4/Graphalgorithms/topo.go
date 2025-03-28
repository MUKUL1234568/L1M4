package main

import "fmt"

type Graph struct {
    ver map[int][]int
}

func NewGraph() *Graph {
    return &Graph{
        ver: make(map[int][]int),
    }
}

func (g *Graph) AddEdge(v int, u int) {
    g.ver[v] = append(g.ver[v], u)
}

func (g *Graph) RemoveEdge(v int, u int) {
    g.ver[v] = removeFromSlice(g.ver[v], u)
}

func removeFromSlice(slice []int, u int) []int {
    for i, val := range slice {
        if val == u {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}

func (g *Graph) print() {
    for ver, edge := range g.ver {
        fmt.Printf("%d: %v\n", ver, edge)
    }
}

func (g *Graph) topo() {
    visited := make(map[int]bool)
    var stack []int

    for v := range g.ver {
        if !visited[v] {
            g.dfs(v, visited, &stack)
        }
    }

    for i := len(stack) - 1; i >= 0; i-- {
        fmt.Println(stack[i])
    }
}

func (g *Graph) dfs(v int, visited map[int]bool, stack *[]int) {
    visited[v] = true

    for _, vt := range g.ver[v] {
        if !visited[vt] {
            g.dfs(vt, visited, stack)
        }
    }

    *stack = append(*stack, v)
}

func main() {
    g := NewGraph()
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)
    g.AddEdge(3, 4)
    
    g.print()
    fmt.Println("Topological Sort:")
    g.topo()
}
