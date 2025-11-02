package main

import "fmt"

func main() {
	// ! Semua loop pakai for, gk ada while dan do while

	buah := []string{"jambu", "salak", "rambutan", "mangga", "durian"}

	for index, value := range buah {
		fmt.Println("Buah ke ", index+1, " adalah : ", value)
	}
}
