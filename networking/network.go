package networking

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/meant4/computer"
	"net/http"
)

/**
calculate endpoint custom middleware for validation
*/
func (api *Api) CalculateMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		api.response = Response{}           // reset struct values back to their defaults
		api.calculateBody = CalculateBody{} // reset struct values back to their defaults
		err := json.NewDecoder(r.Body).Decode(&api.calculateBody)
		if err != nil || !api.isValidCalculateRequest() {
			api.response.Error = "Incorrect input"
			respond(w, http.StatusBadRequest, api.response)
		} else {
			next(w, r, p)
		}
	}
}

/**
handler for calculate endpoint
*/
func (api *Api) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	aCh := make(chan uint)
	bCh := make(chan uint)
	go computer.CalculateUsingGoRoutines(*api.calculateBody.A, aCh) // spin off a go routine to compute A
	go computer.CalculateUsingGoRoutines(*api.calculateBody.B, bCh) // spin off a go routine to compute B
	fA, fB := <-aCh, <-bCh
	api.response.Product = fA + fB // compute and assign product from channels accordingly
	// then close the channels
	close(aCh)
	close(bCh)
	respond(w, http.StatusOK, api.response)
}

/**
check if the value of a or b is an invalid input. uint already will handle and prevent this
from happening during json Decoding of request body. However, this function only needs to check
for null values
*/
func (api *Api) isValidCalculateRequest() bool {
	return api.calculateBody.A != nil && api.calculateBody.B != nil
}

/**
helper function to pack response and write accordingly
*/
func respond(w http.ResponseWriter, code int, response Response) {
	responseJson, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(responseJson)
}
