package practice

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello Play")
	port := 3000
	_, err := startServer(port, 2)
	fmt.Println(err)
}

//func startServer(port, numOfRetries int)  {  /* also can be used but considers all the params as single type of int */
func startServer(port int, numOfRetries int) (int, error) {
	fmt.Println("Starting Server")
	//do something
	fmt.Println("Server started in port:", port)
	fmt.Println("Server Started")

	return port, errors.New("Something went wrong;")
}
