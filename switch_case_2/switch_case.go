package main

import "fmt"

func main(){
	nilai := 2
	switch{
	case nilai == 5 || nilai == 6:
		fmt.Println("Tingkatkan lagi")
	case nilai == 7:
		fmt.Println("Cukup")
	case nilai == 8:
		fmt.Println("Bagus")
	case nilai == 9:
		fmt.Println("Istimewa")
	case nilai == 10:
		fmt.Println("Excelent")
		fmt.Println("Amazing")
	default:
		fmt.Println("Belajar lebih keras bro")
		fmt.Println("Jangan menyerah")
	}
}

/*
outputnya adalah
Belajar lebih keras bro
Jangan menyerah
*/