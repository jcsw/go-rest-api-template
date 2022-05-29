package monitor

import (
	http "net/http"

	mariadb "go-rest-api-template/pkg/infra/database/mariadb"
	mongodb "go-rest-api-template/pkg/infra/database/mongodb"
	sys "go-rest-api-template/pkg/infra/system"
)

type status struct {
	Component string `json:"component"`
	Status    string `json:"status"`
}

// Get function to handle GET /monitor
func Get(w http.ResponseWriter, r *http.Request) {
	allStatus := []status{}

	allStatus = append(allStatus, mongodbStatus())
	allStatus = append(allStatus, mariadbStatus())

	sys.HTTPResponseWithJSON(w, http.StatusOK, allStatus)
}

func mongodbStatus() status {
	if mongodb.IsAlive() {
		return status{Component: "Mongodb", Status: "OK"}
	}

	return status{Component: "Mongodb", Status: "ERROR"}
}
func mariadbStatus() status {
	if mariadb.IsAlive() {
		return status{Component: "Mariadb", Status: "OK"}
	}

	return status{Component: "Mariadb", Status: "ERROR"}
}
