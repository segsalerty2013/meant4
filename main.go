package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/meant4/networking"
	"log"
	"net/http"
)

func main() {
	api := &networking.Api{} //initialize null value of API
	router := httprouter.New()
	router.POST("/calculate", api.CalculateMiddleware(api.Calculate)) // middleware and handler for the endpoint
	fmt.Printf("Application Ready.\n")
	log.Fatal(http.ListenAndServe(":8989", router))
}
