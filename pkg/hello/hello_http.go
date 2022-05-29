package hello

import (
	"fmt"
	"net/http"

	sys "go-rest-api-template/pkg/system"
)

// Handler - HTTP Handler by /hello
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			helloWorld(w, r)
			break
		}
	default:
		{
			sys.HTTPResponseWithError(w, http.StatusBadRequest, "Invalid Method")
		}
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	sys.HTTPResponseWithJSON(w, 200, fmt.Sprintf("Hello %s", name))
}
