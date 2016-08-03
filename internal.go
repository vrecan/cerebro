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

func DepthFirst(g DirectedGraph, from graph.Node, visited map[string]struct{}) (s []graph.Node) {
	s = make([]graph.Node, 0)
	if visited == nil {
		visited = make(map[string]struct{}, 0)
	}
	s = append(s, from)
	visited[from.ID()] = struct{}{}
	deps := g.From(from)
	toDeps := g.To(from)
	for _, td := range toDeps {
		exists := false
		for _, fd := range deps {
			if td.ID() == fd.ID() {
				exists = true
			}
		}
		if !exists {
			deps = append(deps, td)
		}
	}
	// SPEW.Dump(deps)
	for _, n := range deps {
		_, ok := visited[n.ID()]
		if ok {
			continue
		}
		ns := DepthFirst(g, n, visited)
		for _, dfg := range ns {
			s = append(s, dfg)
		}
	}
	return s
}
