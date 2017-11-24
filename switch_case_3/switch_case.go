package main

import "fmt"

func main(){
	nilai := 7
	switch{
	case nilai == 7:
		fmt.Println("Tidak terlalu buruk")
		fallthrough
	case nilai >= 6 && nilai < 9:
		fmt.Println("Tingkatkan lagi, kamu pasti bisa")
	case nilai < 8:
		fmt.Println("Hahahahaha")
	case nilai > 8:
		fmt.Println("Mantap")
	}
}

/*
outputnya adalah :
Tidak terlalu buruk
Tingkatkan lagi, kamu pasti bisa
*/