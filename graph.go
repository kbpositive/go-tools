package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key int
	adj []*Vertex
}

func (graph *Graph) AddVertex(val int) {
	if contains(graph.vertices, val) {
		err := fmt.Errorf("Vertex %v already exists.", val)
		fmt.Println(err.Error())
	} else {
		graph.vertices = append(graph.vertices, &Vertex{key: val})
	}
}

func (graph *Graph) AddEdge(from, to int) {
	f := graph.getVertex(from)
	t := graph.getVertex(to)

	if f == nil || t == nil {
		err := fmt.Errorf("Invalid edge %v-%v", from, to)
		fmt.Println(err.Error())
	} else if contains(f.adj, to) || contains(t.adj, from) {
		err := fmt.Errorf("Existing edge %v-%v", from, to)
		fmt.Println(err.Error())
	} else {
		f.adj = append(f.adj, t)
		t.adj = append(t.adj, f)
	}

}

func (graph *Graph) getVertex(n int) *Vertex {
	for i, v := range graph.vertices {
		if v.key == n {
			return graph.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, n int) bool {
	for _, i := range s {
		if i.key == n {
			return true
		}
	}
	return false
}

func (graph *Graph) Print() {
	for _, i := range graph.vertices {
		fmt.Println((i.key))
		for _, j := range i.adj {
			fmt.Println(" ", j.key)
		}
	}
}

func main() {
	g := &Graph{}

	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	g.Print()
}
