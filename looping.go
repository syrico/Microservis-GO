package main

import "fmt"

func main() {
	var counter int
	counter = 1

	for counter <= 10 {
		fmt.Println("Perulangan ke : ", counter)
		counter = counter + 1
	}

	counter = 1

	println("percabangan 2")

	for counter := counter; counter <= 5; counter = counter + 1 {
		fmt.Println("Perulangan ke : ", counter)
	}
}
