package main

import "fmt"

func main() {

	// slice := []string{"Eko", "Kunadi"}

	for i := 0; i < 10; i++ {
		fmt.Println("perualangan ke : ", i+1)
		if i == 5 {
			break
		}
	}

	fmt.Println("continue..")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("perualangan ke : ", i+1)
	}
	// fmt.Println(slice[1])
}
