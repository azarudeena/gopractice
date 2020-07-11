package concepts

import "fmt"


func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Azar"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3,4)
	fmt.Println(c)

	r , im := real(c), imag(c)
	fmt.Println(r, im)

	//pointer

	var pointerName *string = new(string)
	*pointerName = "Azar"
	fmt.Println(*pointerName)

	// address of

	addressname := "Azar";
	fmt.Println(addressname)

	ptr := &addressname;

	fmt.Println(ptr,*ptr)
	firstName = "Deenu"
	fmt.Println(ptr,*ptr)

	//Constants

	const pi float32 = 3.14;
	fmt.Println(pi)

	const ci int = 3
	fmt.Println(ci + 3)

	fmt.Println(float32(ci) + 1.2)


	// iota
	fmt.Println(test1)
	fmt.Println(first, second, third)
}

const test1 = 939393;

const (
	first = iota // iota
	second // 2 << iota

)
