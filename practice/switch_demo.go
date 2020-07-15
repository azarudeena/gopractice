package practice

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{"GET"}

	switch r.Method {
	case "GET":
		println("Get Request")
		fallthrough // break is implicit. kind of like continue.
	case "DELETE":
		println("DELETE REquest")
	default:
		println("Default case")

	}
}
