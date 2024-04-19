package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{}

	person["name"] = "Eko"
	person["address"] = "Subang"

	// cara 2

	person_2 := map[string]string{
		"name":    "eko",
		"address": "Subang",
	}

	fmt.Println(person["name"], person_2["address"])
}
