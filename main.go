package main

import (
	"NotifyGoBot/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/", handlerMainPage)
	r.HandleFunc("/test", customResp)
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handlerMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the main page!")
}

func customResp(w http.ResponseWriter, r *http.Request) {

	m := model.Notify{From: "Ajeje", Id: 1, Testo: "ciao bello"}
	// non posso tornare oggeto diretto devo prima convertirlo in json
	resp, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}
