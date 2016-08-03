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

	Convey("Add and remove nodes", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("3"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("4"), W: 1})
		g.RemoveNode(Node("2"))
		So(len(g.nodes), ShouldEqual, 3)
		So(g.HasEdgeBetween(Node("1"), Node("2")), ShouldBeFalse)
		So(g.HasEdgeBetween(Node("2"), Node("3")), ShouldBeFalse)
		So(g.HasEdgeBetween(Node("3"), Node("4")), ShouldBeTrue)
	})

}
