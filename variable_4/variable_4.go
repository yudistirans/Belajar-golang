package main

import "fmt"

func main(){
	//deklarasi banyak variabel dengan menuliskan tipe data
	var nama, alamat, email string = "Yudistira", "Purbalingga", "yudistira.nursasongko@gmail.com"
	
	//deklarasi banyak variabel tanpa menuliskan tipe data
	mapel, nilai, dosen := "Bahasa Pemrograman", 80, "Jhon Due"

	fmt.Println(nama, alamat, email)
	fmt.Println(mapel, nilai, dosen)
}