package cerebro

import (
	log "github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDirectedGraph(t *testing.T) {
	defer log.Flush()
	Convey("Add items with edges validate edges", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.AddNode(Node("4"))
		So(len(g.nodes), ShouldEqual, 4)

		So(g.HasEdgeBetween(Node("1"), Node("2")), ShouldBeTrue)
		So(g.HasEdgeBetween(Node("2"), Node("3")), ShouldBeTrue)
		So(g.HasEdgeBetween(Node("3"), Node("4")), ShouldBeFalse)
	})

	Convey("Add same node multiple times doesn't panic", t, func() {
		g := NewDirectedGraph(1)
		g.AddNode(Node("1"))
		g.AddNode(Node("1"))
		g.AddNode(Node("1"))
		g.AddNode(Node("1"))
		So(len(g.nodes), ShouldEqual, 1)
	})

	Convey("Add and remove nodes validate edges", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		g.RemoveNode(Node("2"))
		So(len(g.nodes), ShouldEqual, 3)
		So(g.HasEdgeBetween(Node("1"), Node("2")), ShouldBeFalse)
		So(g.HasEdgeFromTo(Node("1"), Node("2")), ShouldBeFalse)
		So(g.HasEdgeBetween(Node("2"), Node("3")), ShouldBeFalse)
		So(g.HasEdgeFromTo(Node("2"), Node("3")), ShouldBeFalse)
		So(g.HasEdgeBetween(Node("3"), Node("4")), ShouldBeTrue)
		So(g.HasEdgeFromTo(Node("3"), Node("4")), ShouldBeTrue)
		//validate each way works
		So(g.HasEdgeBetween(Node("4"), Node("3")), ShouldBeTrue)

	})

	Convey("Remove node that doesn't exist", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		g.RemoveNode(Node("not_here"))
		So(len(g.Nodes()), ShouldEqual, 4)

	})

	Convey("Add and remove edges validate edges", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		So(len(g.Edges()), ShouldEqual, 3)
		g.RemoveEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		So(len(g.Edges()), ShouldEqual, 2)
		So(g.HasEdgeBetween(Node("1"), Node("2")), ShouldBeTrue)
		So(g.HasEdgeBetween(Node("2"), Node("3")), ShouldBeFalse)
		So(g.HasEdgeBetween(Node("3"), Node("4")), ShouldBeTrue)

	})

	Convey("Add edges grab individual node", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		So(len(g.Edges()), ShouldEqual, 3)
		n := g.Node("1")
		So(n, ShouldEqual, Node("1"))
		invalid := g.Node("5")
		So(invalid, ShouldBeEmpty)
		So(g.Has(Node("1")), ShouldBeTrue)
		So(g.Has(Node("5")), ShouldBeFalse)

	})

	Convey("Add edges grab single edge", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		So(len(g.Edges()), ShouldEqual, 3)
		e, ok := g.Edge(Node("2"), Node("3"))
		So(ok, ShouldBeTrue)
		So(e, ShouldResemble, Edge{F: Node("2"), T: Node("3"), W: 1})
	})

	Convey("Add edges grab invalid edge", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		So(len(g.Edges()), ShouldEqual, 3)
		_, ok := g.Edge(Node("not"), Node("here"))
		So(ok, ShouldBeFalse)
	})
	Convey("get weight of a node in the graph", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		So(len(g.Edges()), ShouldEqual, 3)
		w, ok := g.Weight(Node("1"), Node("2"))
		So(ok, ShouldBeTrue)
		So(w, ShouldEqual, 1)
	})

	Convey("get weight of invalid edge", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		So(len(g.Edges()), ShouldEqual, 3)
		_, ok := g.Weight(Node("not"), Node("here"))
		So(ok, ShouldBeFalse)
	})

	Convey("get degree of node relationships", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("4"), W: 1})

		g.SetEdge(Edge{F: Node("5"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("6"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("7"), T: Node("1"), W: 1})
		So(len(g.Edges()), ShouldEqual, 6)
		So(g.Degree(Node("1")), ShouldEqual, 6)
	})

	Convey("get degree of node relationships for node that doesn't exist", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("4"), W: 1})

		g.SetEdge(Edge{F: Node("5"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("6"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("7"), T: Node("1"), W: 1})
		So(len(g.Edges()), ShouldEqual, 6)
		So(g.Degree(Node("nothere")), ShouldEqual, 0)
	})

	Convey("get All the nodes in the graph", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("4"), W: 1})

		g.SetEdge(Edge{F: Node("5"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("6"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("7"), T: Node("1"), W: 1})
		n := g.Nodes()
		So(len(n), ShouldEqual, 7)
	})

	Convey("get To from Node", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("4"), W: 1})

		g.SetEdge(Edge{F: Node("5"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("6"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("7"), T: Node("1"), W: 1})
		n := g.To(Node("1"))
		So(len(n), ShouldEqual, 3)
	})

}
