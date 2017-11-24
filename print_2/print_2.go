package main 

import "fmt"

func main(){
	nama := "Yudistira Nur Sasongko"
	email := "yudistira.nursasongko@gmail.com"
	
	/*
	Kedua kode di bawah ini akan menghasilkan output yang sama, 
	Dengan Printf kita bisa mendefinisikan format output yang kita inginkan
	Printf tidak menghasilkan baris baru di akhir output
	*/
	fmt.Printf("%s %3s \n",nama,email)
	fmt.Println(nama,email)	
}