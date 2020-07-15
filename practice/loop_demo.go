package practice

type User struct {
	ID int
	FN string
	LN string
}

func main() {
	for i := 0; i < 5; i++ {
		println(i)

		if i == 3 {
			continue
		}
		println("continuing")
	}

	//infinite loop
	j := 0
	for {

		if j == 5 {
			break
		}
		j++
		println(j)
	}

	// iterate over collections

	slice := []int{1, 2, 3}
	//for k := 0 ; k < len(slice); k++ {
	for i, V := range slice {
		println(i, V)
	}

	// map iteration
	wellknownports := map[string]int{"http": 80, "https": 443}
	for _, V := range wellknownports {
		println(V)

	}

	//panic("I am panicking")
	u1 := User{
		ID: 1,
		FN: "Azar",
		LN: "Deen",
	}

	u2 := User{
		ID: 2,
		FN: "Test",
		LN: "Azar",
	}

	if u1.ID == u2.ID {
		println(" Same User")
	} else if u1.FN == u2.FN {
		println("Similar User")
	} else {
		println("Different User")
	}
}
