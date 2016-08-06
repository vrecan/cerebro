package cerebro

import (
	log "github.com/cihub/seelog"
	// SPEW "github.com/davecgh/go-spew/spew"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vrecan/cerebro/graph"
	"testing"
)

func TestDepthFirst(t *testing.T) {
	defer log.Flush()
	Convey("Add items with edges validate edges, multiple layers of relations", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("10"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("11"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("12"), W: 1})

		g.SetEdge(Edge{F: Node("12"), T: Node("2"), W: 1})

		g.SetEdge(Edge{F: Node("2"), T: Node("20"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("21"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("22"), W: 1})

		g.SetEdge(Edge{F: Node("22"), T: Node("3"), W: 1})

		g.SetEdge(Edge{F: Node("3"), T: Node("30"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("31"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("32"), W: 1})

		g.SetEdge(Edge{F: Node("4"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("4"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("4"), T: Node("3"), W: 1})
		r := Dependencies(*g, Node("1"))
		So(len(r), ShouldEqual, 12)
		for _, n := range r {
			fmt.Println(n)
		}

		So(OrderValid(Node("10"), Node("1"), r), ShouldBeTrue)
		So(OrderValid(Node("11"), Node("1"), r), ShouldBeTrue)
		So(OrderValid(Node("12"), Node("1"), r), ShouldBeTrue)
		So(OrderValid(Node("2"), Node("12"), r), ShouldBeTrue)

		So(OrderValid(Node("20"), Node("2"), r), ShouldBeTrue)
		So(OrderValid(Node("21"), Node("2"), r), ShouldBeTrue)
		So(OrderValid(Node("22"), Node("2"), r), ShouldBeTrue)
		So(OrderValid(Node("3"), Node("22"), r), ShouldBeTrue)

		So(OrderValid(Node("30"), Node("3"), r), ShouldBeTrue)
		So(OrderValid(Node("31"), Node("3"), r), ShouldBeTrue)
		So(OrderValid(Node("32"), Node("3"), r), ShouldBeTrue)
	})

	Convey("Same graph but pluck out node 4", t, func() {
		g := NewDirectedGraph(1)
		g.SetEdge(Edge{F: Node("1"), T: Node("10"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("11"), W: 1})
		g.SetEdge(Edge{F: Node("1"), T: Node("12"), W: 1})

		g.SetEdge(Edge{F: Node("12"), T: Node("2"), W: 1})

		g.SetEdge(Edge{F: Node("2"), T: Node("20"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("21"), W: 1})
		g.SetEdge(Edge{F: Node("2"), T: Node("22"), W: 1})

		g.SetEdge(Edge{F: Node("22"), T: Node("3"), W: 1})

		g.SetEdge(Edge{F: Node("3"), T: Node("30"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("31"), W: 1})
		g.SetEdge(Edge{F: Node("3"), T: Node("32"), W: 1})

		g.SetEdge(Edge{F: Node("4"), T: Node("1"), W: 1})
		g.SetEdge(Edge{F: Node("4"), T: Node("2"), W: 1})
		g.SetEdge(Edge{F: Node("4"), T: Node("3"), W: 1})
		r := Dependencies(*g, Node("4"))
		So(len(r), ShouldEqual, 13)
		for _, n := range r {
			fmt.Println(n)
		}

		So(OrderValid(Node("10"), Node("1"), r), ShouldBeTrue)
		So(OrderValid(Node("11"), Node("1"), r), ShouldBeTrue)
		So(OrderValid(Node("12"), Node("1"), r), ShouldBeTrue)
		So(OrderValid(Node("2"), Node("12"), r), ShouldBeTrue)

		So(OrderValid(Node("20"), Node("2"), r), ShouldBeTrue)
		So(OrderValid(Node("21"), Node("2"), r), ShouldBeTrue)
		So(OrderValid(Node("22"), Node("2"), r), ShouldBeTrue)
		So(OrderValid(Node("3"), Node("22"), r), ShouldBeTrue)

		So(OrderValid(Node("30"), Node("3"), r), ShouldBeTrue)
		So(OrderValid(Node("31"), Node("3"), r), ShouldBeTrue)
		So(OrderValid(Node("32"), Node("3"), r), ShouldBeTrue)

		So(OrderValid(Node("1"), Node("4"), r), ShouldBeTrue)
		So(OrderValid(Node("2"), Node("4"), r), ShouldBeTrue)
		So(OrderValid(Node("3"), Node("4"), r), ShouldBeTrue)

	})

}

func OrderValid(first graph.Node, second graph.Node, r []graph.Node) bool {
	for _, a := range r {
		if a.ID() == first.ID() {
			return true
		}
		if a.ID() == second.ID() {
			return false
		}
	}
	return false //doesn't exist
}
