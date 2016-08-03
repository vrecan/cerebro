package cerebro

import (
	log "github.com/cihub/seelog"
	// SPEW "github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vrecan/cerebro/graph"
	"testing"
)

func TestDepthFirst(t *testing.T) {
	defer log.Flush()
	Convey("Add items with edges validate edges", t, func() {
		//Don't add nodes at the same level for this test because
		//they are stored in maps and will rearange your results randomly at the same level
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

		r := DepthFirst(*g, Node("1"), nil)
		So(len(r), ShouldEqual, 13)

		Convey("Validate 1 order", func() {
			So(OrderValid(Node("1"), Node("10"), r), ShouldBeTrue)
			So(OrderValid(Node("1"), Node("11"), r), ShouldBeTrue)
			So(OrderValid(Node("1"), Node("12"), r), ShouldBeTrue)
			So(OrderValid(Node("1"), Node("2"), r), ShouldBeTrue)
			So(OrderValid(Node("1"), Node("3"), r), ShouldBeTrue)
			So(OrderValid(Node("1"), Node("4"), r), ShouldBeTrue)
		})
		Convey("Validate 2 order", func() {
			So(OrderValid(Node("2"), Node("20"), r), ShouldBeTrue)
			So(OrderValid(Node("2"), Node("21"), r), ShouldBeTrue)
			So(OrderValid(Node("2"), Node("22"), r), ShouldBeTrue)
			So(OrderValid(Node("2"), Node("3"), r), ShouldBeTrue)
			So(OrderValid(Node("2"), Node("4"), r), ShouldBeTrue)
		})

		Convey("Validate 3 order", func() {
			So(OrderValid(Node("3"), Node("30"), r), ShouldBeTrue)
			So(OrderValid(Node("3"), Node("31"), r), ShouldBeTrue)
			So(OrderValid(Node("3"), Node("32"), r), ShouldBeTrue)
			So(OrderValid(Node("3"), Node("4"), r), ShouldBeTrue)
		})

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
