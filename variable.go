package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	//declarasi pertama menggunakan :=

	name_2 := "Wahyu"
	fmt.Println(name_2)

	name_2 = "Wahyudi"
	fmt.Println(name_2)

	//cara ke tiga

	var (
		firstName = "Eko"
		lastName  = "Kenedy"
	)

	fmt.Println(firstName, lastName)
}
