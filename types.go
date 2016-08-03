package cerebro

import (
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
