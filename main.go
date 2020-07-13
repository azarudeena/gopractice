package main

import (
	"fmt"
	"github.com/azarudeena/webservice/controller"
	"net/http"
)

func main() {
	/*u := models.User{
		ID:        2,
		FirstName: "Azarudeen",
		LastName:  "AKB",
	}
	fmt.Println(u)*/
	fmt.Println("Server Starting")
	controller.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	fmt.Println("Server started")
}
