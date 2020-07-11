package main

import (
	"fmt"
	"github.com/azarudeena/webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Azarudeen",
		LastName: "AKB",
	}
	fmt.Println(u)
}
