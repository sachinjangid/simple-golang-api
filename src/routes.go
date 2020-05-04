package src

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/simple-golang-api/src/controller"
)

// APIRoutes :
func APIRoutes(r *mux.Router) {

	r.HandleFunc("/", testFunc)

	r.HandleFunc("/list-users", controller.GetUsers)
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<div><h1>Simple API</h1></div>"))
}
