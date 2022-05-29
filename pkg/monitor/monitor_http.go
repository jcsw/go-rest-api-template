package monitor

import (
	http "net/http"

	mariadb "go-rest-api-template/pkg/database/mariadb"
	mongodb "go-rest-api-template/pkg/database/mongodb"
	sys "go-rest-api-template/pkg/system"
)

type status struct {
	Component string `json:"component"`
	Status    string `json:"status"`
}

// Handler function to handle "/monitor"
func Handler(w http.ResponseWriter, r *http.Request) {
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
