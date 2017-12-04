package main

import "fmt"

func main(){
	// Pengisian elemen array pada saat deklarasi variabel
	// Pembuatan variabel type inference / tanpa menuliskan tipe data
	var mapel = [3]string{"Matematika","Bahasa Indonesia","Bahasa Inggris"}
	
	// Pengisian elemen array setelah deklarasi variabel
	var nilai [3]int
	nilai = [3]int{10,9,9}

	fmt.Println("Elemen array mapel adalah : ",mapel)
	fmt.Println("Elemen array nilai adalah : ",nilai)
}