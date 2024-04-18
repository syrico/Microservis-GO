package main

import "fmt"

func main() {

	//conversi int ke int

	var var1 int32 = 32767
	var var2 int64 = int64(var1)
	var var3 int16 = int16(var2)
	fmt.Println(var1, var2, var3)

	//conversi byte ke string

	var (
		var4 string = "Eko"
		var5        = var4[0]
		var6 string = string(var5)
	)
	fmt.Println(var4, var5, var6)

}
