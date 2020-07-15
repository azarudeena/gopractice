package practice

import "fmt"

/*
	Arrays
	Slices
	Maps
	Struts - No class in go.
*/
func main() {
	// Array have to have same type.

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// Slice
	slice := arr1[:]
	//arr1[1] = 42
	//slice[2] = 27
	fmt.Println(arr1, slice)

	sliced := []int{1, 2, 3}
	fmt.Println(sliced)

	sliced = append(sliced, 4)
	fmt.Println(sliced)

	//
	s2 := sliced[1:]
	s3 := sliced[:2]
	s4 := sliced[1:2]

	fmt.Println(s2, s3, s4)

	// Maps
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m["foo"])
	delete(m, "foo")
	fmt.Println(m)

	// Struts
	var u user
	u.ID = 1
	u.FirstName = "Azar"
	u.LastName = "Deenu"
	fmt.Println(u)

	u2 := user{ID: 2,
		FirstName: "Ajju",
		LastName:  "Deenu",
	}
	fmt.Println(u2)

}

type user struct {
	ID        int
	FirstName string
	LastName  string
}
