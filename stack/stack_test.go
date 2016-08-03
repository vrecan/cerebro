package stack

import (
	log "github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStack(t *testing.T) {
	defer log.Flush()
	Convey("Validate order of items coming off the stack", t, func() {
		s := Stack{}
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)
		So(s.Pop(), ShouldEqual, 5)
		So(s.Len(), ShouldEqual, 4)
	})

	Convey("Use stack and reset, it should be empty", t, func() {
		s := Stack{}
		s.Push("1")
		s.Push("2")
		s.Push("3")
		So(s.Pop(), ShouldEqual, "3")
		So(s.Len(), ShouldEqual, 2)
		s.Reset()
		So(s.Len(), ShouldEqual, 0)
	})

}
