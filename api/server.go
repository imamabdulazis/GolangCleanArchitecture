package api

import (
	"fmt"
	"log"
	"net/http"

	"tugasakhircoffe/TaCoffe/api/router"
	"tugasakhircoffe/TaCoffe/config"
)

func Init() {
	config.Load()
	// auto.Load()
}

//Run message
func Run() {
	Init()
	fmt.Printf("\n\tListening [::]%d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
