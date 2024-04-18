package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Eko"
	names[1] = "Jaya"
	names[2] = "Kuladi"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(string(names[0][1]))

	//pembuatan array langsung

	var values = [3]string{
		"Eko",
		"Jaya",
	}

	fmt.Println(values[2])
}
