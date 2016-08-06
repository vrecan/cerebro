package cerebro

import (
	// SPEW "github.com/davecgh/go-spew/spew"
	// "fmt"
	"github.com/vrecan/cerebro/graph"
	"math"
)

// Node of the graph
type Node string

// ID returns the string value of the node
func (n Node) ID() string {
	return string(n)
}

// Edge is the links of the graph
type Edge struct {
	F, T graph.Node
	W    float64
}

// From returns start position of the edge
func (e Edge) From() graph.Node { return e.F }

// To returns the end position of the edge
func (e Edge) To() graph.Node { return e.T }

// Weight returns the weight of the edge.
func (e Edge) Weight() float64 { return e.W }

// isSame returns whether two float64 values are the same where NaN values
// are equalable.
func isSame(a, b float64) bool {
	return a == b || (math.IsNaN(a) && math.IsNaN(b))
}

func Dependencies(g DirectedGraph, vtx graph.Node) []graph.Node {
	deps := make([]graph.Node, 0)
	TopologicalSort(g, vtx, nil, &deps)
	return deps
}

func TopologicalSort(g DirectedGraph, vtx graph.Node, visited map[string]struct{}, s *[]graph.Node) {
	if s == nil {
		return
	}
	if visited == nil {
		visited = make(map[string]struct{}, 0)
	}
	_, ok := visited[vtx.ID()]
	if !ok {
		//ignore yourself in your graph
		visited[vtx.ID()] = struct{}{}
	}

	fromDeps := g.From(vtx)
	for _, n := range fromDeps {
		if _, ok := visited[n.ID()]; !ok {
			TopologicalSort(g, n, visited, s)
		}
	}
	*s = append(*s, vtx)
}
