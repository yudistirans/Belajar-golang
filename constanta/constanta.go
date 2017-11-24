package main

import "fmt"

func main(){
	//deklarasi konstanta manifest typing atau dengan menyertakan tipe variabel
	const nilai_a int = 15

	//deklarasi konstanta type inference atau tidak menyertakan tipe data
	const nilai_b = 30
	
	fmt.Println("Nilai konstanta a : ", nilai_a)
	fmt.Println("Nilai konstanta b : ", nilai_b)
}