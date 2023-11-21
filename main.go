package main

import "fmt"

//aggregate types (сукупні, сумарні типи) (array -масив, struct- )

//array

/*func main() {

	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("FIRST ELEMENT IN ARRAY IS", myStrings[0])
}*/

// struct

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        false,
		BucketSeats:   true,
		Make:          "BMW",
		Model:         "Q7",
		Year:          2025,
	}

	fmt.Printf("My car is %s %s %d year", myCar.Make, myCar.Model, myCar.Year)
}
