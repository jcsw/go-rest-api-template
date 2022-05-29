package mongodb_test

import (
	testing "testing"

	. "github.com/smartystreets/goconvey/convey"

	mongodb "go-rest-api-template/pkg/database/mongodb"
	sys "go-rest-api-template/pkg/system"
)

func TestSpec(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	Convey("Given a valid Mongodb URI", t, func() {

		sys.Properties = sys.AppProperties{
			Mongodb: "mongodb://gorest:gorest_pdw@localhost:27017/admin?connectTimeoutMS=1000&serverSelectionTimeoutMS=1000&socketTimeoutMS=1500",
		}

		Convey("When connect on Mongodb", func() {
			mongodb.Connect()
			defer mongodb.Disconnect()

			Convey("Then Mongodb session it's alive", func() {
				So(mongodb.IsAlive(), ShouldEqual, true)
			})
		})
	})
}
