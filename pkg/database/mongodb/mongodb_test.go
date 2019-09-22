package mongodb_test

import (
	testing "testing"

	. "github.com/smartystreets/goconvey/convey"

	mongodb "github.com/jcsw/go-rest-api-template/pkg/database/mongodb"
	sys "github.com/jcsw/go-rest-api-template/pkg/system"
)

func TestSpec(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	Convey("Given an valid Mongodb URI", t, func() {

		sys.Properties = sys.AppProperties{
			Mongodb: "mongodb://latanton:latanton_pdw@localhost:27017/admin?connectTimeoutMS=1000&serverSelectionTimeoutMS=1000&socketTimeoutMS=1500",
		}

		Convey("When connect an Mongodb", func() {

			mongodb.Connect()
			defer mongodb.Disconnect()

			Convey("The Mongodb session it's alive", func() {
				So(mongodb.IsAlive(), ShouldEqual, true)
			})

		})

	})

}
