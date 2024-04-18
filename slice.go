package main

import "fmt"

func main() {
	var arr1 [2]string

	arr1[0] = "EKO"
	arr1[1] = "LISTYO"

	fmt.Println(arr1)

	// cara2
	arr2 := [...]string{"EKO", "LISTYO", "SIGIT"}
	fmt.Println(arr2)

	//SLICE
	var slice1 []string = arr2[:1]
	fmt.Println(slice1)

	//cara 2 SLICE
	slice2 := arr2[1:2]
	fmt.Println(slice2)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daysSlice1 := days[5:]

	daysSlice1[0] = "sabtu baru"
	daysSlice1[1] = "minggu baru"

	daysSlice2 := append(daysSlice1, "hari libur")

	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
}
