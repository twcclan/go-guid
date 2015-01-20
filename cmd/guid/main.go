package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/twcclan/go-guid"
)

func calculateGuid(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	defer r.Body.Close()

	if _, err := io.Copy(&body, r.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed reading request body: %s", err)
	} else {
		etkey := string(body.Bytes())
		if g, err := guid.Calculate(etkey); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Failed to calculate guid: %s", err)
		} else {
			fmt.Fprint(w, g)
		}
	}
}

func getMux() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", calculateGuid).Methods("POST").Headers("Content-Length", "18")

	return router
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	host := os.Getenv("HOST")

	log.Fatal(http.ListenAndServe(host+":"+port, getMux()))
}
