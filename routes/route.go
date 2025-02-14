package routes

import (
	"TaxCalcPoints/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/tax", controllers.TaxHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
