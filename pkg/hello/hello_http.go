package hello

import (
	"net/http"

	sys "github.com/jcsw/go-rest-api-template/pkg/system"
)

// Handler - HTTP Handler by /
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			helloWorld(w, r)
		}
	default:
		{
			sys.HTTPResponseWithError(w, http.StatusBadRequest, "Invalid Method")
		}
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	sys.HTTPResponseWithJSON(w, 200, "Hello World!!")
}
