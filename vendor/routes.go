package vendor

import (
	"net/http"

	"controller"
	"github.com/gorilla/mux"
)

// APIRoutes :
func APIRoutes(r *mux.Router) {

	r.HandleFunc("/", testFunc)

	r.HandleFunc("/login", controller.UserLogin)
	r.HandleFunc("/list-users", controller.GetUsers)
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<div><h1>Simple API</h1></div>"))
}
