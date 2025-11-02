package main

import "fmt"

func main() {
	// ! Semua loop pakai for, gk ada while dan do while

	buah := []string{"jambu", "salak", "rambutan", "mangga", "durian"}

	for index, value := range buah {
		fmt.Println("Buah ke ", index+1, " adalah : ", value)
	}

	buah1 := buah
	fmt.Printf("Awal - Alamat buah : %p, isi : %v\n", buah, buah)
	fmt.Printf("Awal - Alamat buah1 : %p, isi : %v\n", buah1, buah1)
	
	buah = append(buah, "Manggis")
	fmt.Println("\n-- Setelah Append --")
	fmt.Printf("Baru - Alamat Bru buah : %p, isi : %v\n", buah, buah)
	fmt.Printf("Lama - Alamat buah1 : %p, isi : %v\n", buah1, buah1)
	
	buah[0] = "Kelapa"
	fmt.Println(buah)
	fmt.Println(buah1)
	
}
