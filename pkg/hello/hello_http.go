package hello

import (
	"fmt"
	"net/http"

	sys "go-rest-api-template/pkg/infra/system"
)

// Get function to handle GET /hello
func Get(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	sys.HTTPResponseWithJSON(w, 200, fmt.Sprintf("Hello %s", name))
}
