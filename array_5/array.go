package main

import "fmt"

func main(){
	var mapel = [5]string{"Matematika","Bahasa Indonesia","Bahasa Inggris","IPA","IPS"}
	
	// For biasa
	fmt.Println("Cara 1 : ")
	for i:=0;i<len(mapel);i++ {
		fmt.Printf("Elemen ke %d adalah %s \n",i,mapel[i])
	}

	// For range
	fmt.Println("Cara 2 : ")
	for i, ElemenMapel := range mapel{
		fmt.Printf("Elemen ke %d adalah %s \n",i,ElemenMapel)
	}

	// For range dan penggunaan underscore
	fmt.Println("Penggunaan underscore untuk variabel yang tidak digunakan")
	for _, ElemenMapel :=range mapel{
		fmt.Printf("Nama mapel : %s \n",ElemenMapel)
	}
}
