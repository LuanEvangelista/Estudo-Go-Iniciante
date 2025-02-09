package main

import "fmt"

func main() {
	var variavel string = "Variavel 1" //Jeito 1
	variave2 := "Variavel 2"           //Jeito 2 (Implisito)
	fmt.Println(variavel)
	fmt.Println(variave2)

	var (
		variave3 string = "11"
		variave4 string = "22"
	)

	fmt.Println(variave3, variave4)
	variave5, variave6 := "33", "44"

	fmt.Println(variave5, variave6)
}
