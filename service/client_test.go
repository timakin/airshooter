package service

import (
	. "github.com/smartystreets/goconvey/convey"
	db "github.com/timakin/airshooter/datasource"
	"testing"
)

func TestIntegerStuff(t *testing.T) {
	db.Init()

	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})

	Reset(func() {
		db.Close()
	})
}
