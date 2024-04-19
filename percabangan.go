package main

import "fmt"

func main() {
	nilai := 80

	if nilai >= 90 {
		fmt.Println("A")
	} else if nilai >= 80 && nilai < 90 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
