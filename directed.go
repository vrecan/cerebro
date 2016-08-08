package cerebro

import (
	"fmt"
	"github.com/vrecan/cerebro/graph"
)

// DirectedGraph implements a generalized directed graph.
type DirectedGraph struct {
	nodes map[string]graph.Node
	from  map[string]map[string]graph.Edge
	to    map[string]map[string]graph.Edge

	weight, absent float64
}

//NewDirectedGraph creates a directional graph
func NewDirectedGraph(weight float64) *DirectedGraph {
	return &DirectedGraph{
		nodes: make(map[string]graph.Node),
		from:  make(map[string]map[string]graph.Edge),
		to:    make(map[string]map[string]graph.Edge),

		weight: weight,
		absent: 0,
	}
}

// AddNode adds n to the graph. It panics if the added node ID matches an existing node ID.
func (g *DirectedGraph) AddNode(n graph.Node) {
	if _, exists := g.nodes[n.ID()]; exists {
		return
	}
	g.nodes[n.ID()] = n
	g.from[n.ID()] = make(map[string]graph.Edge)
	g.to[n.ID()] = make(map[string]graph.Edge)

}

// RemoveNode removes n from the graph, as well as any edges attached to it
func (g *DirectedGraph) RemoveNode(n graph.Node) {
	if _, ok := g.nodes[n.ID()]; !ok {
		return
	}
	delete(g.nodes, n.ID())

	for from := range g.from[n.ID()] {
		delete(g.to[from], n.ID())
	}
	delete(g.from, n.ID())

	for to := range g.to[n.ID()] {
		delete(g.from[to], n.ID())
	}
	delete(g.to, n.ID())
}

// SetEdge add's an edge and their nodes if they don't already exist, panics if you set an edge
// of the same value as the edge it's working on.
func (g *DirectedGraph) SetEdge(e graph.Edge) {
	var (
		from = e.From()
		fid  = from.ID()
		to   = e.To()
		tid  = to.ID()
	)

	if fid == tid {
		panic(fmt.Sprintf("Adding edge to self: %v", fid))
	}

	if !g.Has(from) {
		g.AddNode(from)
	}
	if !g.Has(to) {
		g.AddNode(to)
	}
	g.from[fid][tid] = e
	g.to[tid][fid] = e
}

//RemoveEdge from the graph
func (g *DirectedGraph) RemoveEdge(e graph.Edge) {
	from, to := e.From(), e.To()
	if _, ok := g.nodes[from.ID()]; !ok {
		return
	}
	if _, ok := g.nodes[to.ID()]; !ok {
		return
	}

	delete(g.from[from.ID()], to.ID())
	delete(g.to[to.ID()], from.ID())
}

// Has returns if it exists or not.
func (g *DirectedGraph) Has(n graph.Node) bool {
	_, ok := g.nodes[n.ID()]

	return ok
}

// Node returns the node in the graph with the given ID.
func (g *DirectedGraph) Node(id string) graph.Node {
	return g.nodes[id]
}

// Nodes returns all the nodes.
func (g *DirectedGraph) Nodes() []graph.Node {
	nodes := make([]graph.Node, len(g.from))
	i := 0
	for _, n := range g.nodes {
		nodes[i] = n
		i++
	}

	return nodes
}

// Edges returns all the edges in the graph.
func (g *DirectedGraph) Edges() []graph.Edge {
	var edges []graph.Edge
	for _, u := range g.nodes {
		for _, e := range g.from[u.ID()] {
			edges = append(edges, e)
		}
	}
	return edges
}

// From returns all nodes in the graph that can be reached directly from n.
func (g *DirectedGraph) From(n graph.Node) []graph.Node {
	if _, ok := g.from[n.ID()]; !ok {
		return nil
	}

	from := make([]graph.Node, len(g.from[n.ID()]))
	i := 0
	for id := range g.from[n.ID()] {
		from[i] = g.nodes[id]
		i++
	}

	return from
}

// To returns all nodes in the graph that can reach directly to n.
func (g *DirectedGraph) To(n graph.Node) []graph.Node {
	if _, ok := g.from[n.ID()]; !ok {
		return nil
	}

	to := make([]graph.Node, len(g.to[n.ID()]))
	i := 0
	for id := range g.to[n.ID()] {
		to[i] = g.nodes[id]
		i++
	}

	return to
}

// HasEdgeBetween returns whether an edge exists between nodes x and y. It does not
// care which direction
func (g *DirectedGraph) HasEdgeBetween(x, y graph.Node) bool {
	xid := x.ID()
	yid := y.ID()
	if _, ok := g.nodes[xid]; !ok {
		return false
	}
	if _, ok := g.nodes[yid]; !ok {
		return false
	}
	if _, ok := g.from[xid][yid]; ok {
		return true
	}
	_, ok := g.from[yid][xid]
	return ok
}

// Edge returns the edge from u to v. OK returns fals if no edge exists
func (g *DirectedGraph) Edge(u, v graph.Node) (edge graph.Edge, ok bool) {
	if _, ok = g.nodes[u.ID()]; !ok {
		return edge, ok
	}
	if _, ok = g.nodes[v.ID()]; !ok {
		return edge, ok
	}
	edge, ok = g.from[u.ID()][v.ID()]
	if !ok {
		return edge, ok
	}
	return edge, true
}

// HasEdgeFromTo returns whether an edge exists in the graph from u to v.
func (g *DirectedGraph) HasEdgeFromTo(u, v graph.Node) bool {
	if _, ok := g.nodes[u.ID()]; !ok {
		return false
	}
	if _, ok := g.nodes[v.ID()]; !ok {
		return false
	}
	if _, ok := g.from[u.ID()][v.ID()]; !ok {
		return false
	}
	return true
}

// Weight returns the weight of two nodes in a graph if they are not linked
// ok will be false
func (g *DirectedGraph) Weight(x, y graph.Node) (w float64, ok bool) {
	xid := x.ID()
	yid := y.ID()
	if xid == yid {
		return g.weight, true
	}
	if to, ok := g.from[xid]; ok {
		if e, ok := to[yid]; ok {
			return e.Weight(), true
		}
	}
	return g.absent, false
}

// Degree returns the in+out degree of n in g.
func (g *DirectedGraph) Degree(n graph.Node) int {
	if _, ok := g.nodes[n.ID()]; !ok {
		return 0
	}
	return len(g.from[n.ID()]) + len(g.to[n.ID()])
}
