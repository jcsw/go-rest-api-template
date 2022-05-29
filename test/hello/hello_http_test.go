package hello_test

import (
	http "net/http"
	httptest "net/http/httptest"
	testing "testing"

	hello "go-rest-api-template/pkg/hello"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("Given a GET request", t, func() {

		req, _ := http.NewRequest("GET", "/hello?name=test", nil)
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(hello.Get)

		Convey("When receive request", func() {

			handler.ServeHTTP(resp, req)

			Convey("Then return status code 200", func() {
				So(resp.Code, ShouldEqual, http.StatusOK)
			})

			Convey("And text 'Hello test'", func() {
				So(resp.Body.String(), ShouldEqual, "\"Hello test\"")
			})
		})
	})
}
