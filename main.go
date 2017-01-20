package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	. "github.com/larskluge/babl-server/utils"
)

const listen = ":7000"

func main() {

	pwd, err := os.Getwd()
	Check(err)
	dir := pwd + "/static"

	r := mux.NewRouter()
	// Static files and assets
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      r,
		Addr:         listen,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
