package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transporte/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", HomeHandler)
	//rutas de autentificacion
	routes.RutasAuth(r)
	routes.RutasSeguridad(r)
	fmt.Println("Server on port 5000")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{"api": "apiexample", "version": 1.1}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
