package main

import "fmt"

func main(){
	//Hasil output Print tidak memiliki pemisah(spasi) antar variabel
	fmt.Print("Golang","singkatan","dari","Google","Language\n")

	//Hasil output Println memiliki pemisah(spasi) antar variabel
	fmt.Println("Golang","singkatan","dari","Google","Language")

	//Print tidak menghasilkan baris baru di akhir outputnya
	fmt.Print("Aku suka Golang")
	fmt.Print("Golang itu keren\n")

	//Println menghasilkan baris baru di akhir outputnya
	fmt.Println("Aku suka Golang")
	fmt.Println("Golang itu keren")
}