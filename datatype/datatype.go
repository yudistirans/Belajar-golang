package main

import "fmt"

func main(){
	//deklarasi string menggunakan tanda petik dua (digunakan untuk datu baris string)
	pesan := "Aku sedang belajar bahasa pemrograman Go"

	//deklarasi string menggunakan tanda backticks semua yang ada di dalamnya akan diperlakukan sebagai string
	deskripsi := `Golang adalah bahasa pemrograman yang dikembangkan oleh karyawan Google.
	Golang adalah bahasa yang cepat dan sangat disiplin
	Golang singkatan dari "Google Language".`
	
	fmt.Println(pesan)
	fmt.Println(deskripsi)
}