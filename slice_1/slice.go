package main

import "fmt"

func main(){
	// Deklarasi array disertai jumlah elemen
	var hari = [7]string{"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu"}

	// Deklarasi array tanpa jumlah elemen
	var bulan = [...]string{"Januari","Februari","Maret","April","Mei","Juni"}

	// Deklarasi slice
	var tahun = []string{"2017","2018","2019","2020"}	

	fmt.Println("Array hari : ",hari)
	fmt.Println("Array bulan : ",bulan)
	fmt.Println("Slice tahun : ",tahun)
}