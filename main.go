package main

import (
	"log"
	"net/http"

	"github.com/VictorCometti/atendimento.git/src/router"
)

func main() {
	r := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
