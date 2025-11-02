package main

import "fmt"

// ! Switch tidak perlu manual, mengetik break

func main() {
	angka := 10
	switch angka {
	case 10:
		fmt.Println("angka adalah 10")
	case 20:
		fmt.Println("angka adalah 20")
	default:
		fmt.Println("angka tidak dikenali")
	}

	// ! Contoh switch short statement
	// * Dengan membuat varialbe, untuk block switch biar kode makin rapi
	nama := "budi"
	switch panjang := len(nama); {
	case panjang > 10:
		fmt.Println("nama terlalu panjang")
	case panjang < 5:
		fmt.Println("nama terlalu pendek")
	default:
		fmt.Println("nama sudah benar")
	}
}
