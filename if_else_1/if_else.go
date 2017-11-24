package main

import "fmt"

func main(){
	usia := 15
	if usia <= 3 {
		fmt.Println("Tarif penumpang : Rp. 0")
	}else if usia < 17 {
		fmt.Println("Tarif penumpang : Rp. 5.000")
	}else if usia < 50 {
		fmt.Println("Tarif penumpang : Rp. 10.000")
	}else {
		fmt.Println("Tarif penumpang : Rp. 6.000")
	}
}

// outputnya adalah
// Tarif penumpang : Rp. 5.000