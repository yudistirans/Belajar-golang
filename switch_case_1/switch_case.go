package main

import "fmt"

func main(){
	nilai := 10
	switch nilai{
	case 5,6:
		fmt.Println("Tingkatkan lagi") 
	case 7:
		fmt.Println("Cukup")
	case 8:
		fmt.Println("Bagus")
	case 9:
		fmt.Println("Istimewa")
	case 10:
			fmt.Println("Excelent")
			fmt.Println("Amazing")
	default:
		{
			fmt.Println("Belajar lebih keras bro")
			fmt.Println("Jangan menyerah")
		}
	}
}


/*
outputnya adalah 
Excelent
Amazing
*/