package main 

import "fmt"

func main(){
	// Deklarasi variabel dan alokasi elemen array menggunakan keyword make
	var mapel = make([]string,3)

	// Pengisian elemen array juga dapat dilakukan dengan mengakses elemen menggunakan indexnya
	mapel[0] = "Matematika"
	mapel[1] = "Bahasa Indonesia"
	mapel[2] = "Bahasa Inggris"

	fmt.Println("Elemen array mapel adalah : ",mapel)
}