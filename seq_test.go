package sequence

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func Test_SequenceInit(t *testing.T) {
	Convey("When a Sequence is initialized", t, func() {
		seq := New(0)
		Convey("the values should align", func() {
			for i := 1; i < 1000; i++ {
				So(seq.Next(), ShouldEqual, i)
			}
		})
	})
}
