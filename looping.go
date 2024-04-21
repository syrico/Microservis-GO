package main

import "fmt"

func main() {
	var counter int
	array := [...]string{"Eko", "Kunadi"}

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

	//looping untuk mengakses tiap elemen didalam array
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// cara 2

	for index, names := range array {
		fmt.Println(index, names)
	}

	// cara 2 hanya menampilkan nama tanpa index

	for _, names := range array {
		fmt.Println(names)
	}
}
