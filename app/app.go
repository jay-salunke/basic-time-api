package app

import (
	"fmt"
	"github.com/jay-salunke/TimeAPI/Routers"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("Time API in golang")
	r := Routers.Routers()
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
