package main

import "fmt"

func main() {
	//declare new type with type

	type typeKTP string

	var dataKTP string = "1111"

	var convertKTP typeKTP = typeKTP(dataKTP)

	fmt.Println(dataKTP, convertKTP)

}
