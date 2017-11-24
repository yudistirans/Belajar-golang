package main

import "fmt"

func main(){
	total_nilai := 220
	jumlah_mapel := 3

	if rata2 := float64(total_nilai)/float64(jumlah_mapel); rata2 <=40 {
		fmt.Printf("Nilai rata2 kamu %.1f, Belajar lebih keras\n", rata2)
	}else if rata2 < 60 {
		fmt.Printf("Nilai rata2 kamu %.1f, Tingkatkan lagi\n", rata2)
	}else if rata2 < 80 {
		fmt.Printf("Nilai rata2 kamu %.1f, Sudah cukup bagus\n", rata2)
	}else if rata2 == 80 {
		fmt.Printf("Nilai rata2 kamu %.1f, Bagus\n", rata2)
	}else if rata2 > 80{
		fmt.Printf("Nilai rata2 kamu %.1f, Istimewa\n", rata2)
	}
}

// outputnya adalah
// Nilai rata2 kamu 73.3, Sudah cukup bagus