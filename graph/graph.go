package graph

// Node is a graph node. It returns a unique string ID.
type Node interface {
	ID() string
}

//Edge defines the links of of the graph with weighted values
type Edge interface {
	From() Node
	To() Node
	Weight() float64
}
