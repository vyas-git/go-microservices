package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vyas-git/go-microservices/details"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

}
func apiHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"status": "200", "msg": "Success"}
	json.NewEncoder(w).Encode(resp)
}
func hostDetailsHandler(w http.ResponseWriter, r *http.Request) {
	hostName, err := details.GetHostName()
	if err != nil {
		log.Println(err)
	}
	resp := map[string]string{"status": "200", "msg": "Success", "hostname": hostName}
	json.NewEncoder(w).Encode(resp)

}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", rootHandler)
	r.HandleFunc("/api", apiHandler)
	r.HandleFunc("/details", hostDetailsHandler)

	http.ListenAndServe(":80", r)
}
