package main

import "fmt"

func main(){
	nilai := 5
	if nilai < 6 {
		fmt.Println("Tidak lulus")
		switch nilai{
		case 4,5:
			fmt.Println("Ayo coba lagi")
		case 0,1,2,3:
			fmt.Println("Belajar lebih keras")
		}
	}else{
		fmt.Println("Lulus")
		switch nilai {
		case 6,7:
			fmt.Println("Tingkatkan belajar")
		case 8:
			fmt.Println("Bagus, kalau bisa tingkatkan")
		case 9:
			fmt.Println("Istimewa")
		case 10:
			fmt.Println("Sempurna")
			fmt.Println("Pertahankan")
		}
	}
}

/*
outputnya adalah
Tidak lulus
Ayo coba lagi
*/