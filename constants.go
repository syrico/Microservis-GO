package main

import "fmt"

func main() {
	// constan diawali dengan cons
	// constan seprti tuple, tidak bisa diubah
	const firstName string = "Eko"

	const lastName = "Cahyadi"

	// firstName = "Eno"

	fmt.Println(firstName, lastName)

	// cara ke dua, create multiple constans

	const (
		alamat string = "Mranggen"
		detail string = "RT 05"
	)

	fmt.Println(alamat, detail)

}
