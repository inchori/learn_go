package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func main() {
	var amount, total float64
	amount, err := paintNeeded(4.2, -3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	fmt.Println(err)
	total += amount
	amount, err = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f needed\n", total)

	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
}
