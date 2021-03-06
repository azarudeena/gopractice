package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {

	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc) // not the same /test and /test/

}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)

}
