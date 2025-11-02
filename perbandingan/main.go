package main

import "fmt"

func main() {
	angka1 := 10
	angka2 := 20

	fmt.Print("Masukkan Angka Pertama : ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan Angka Kedua : ")
	fmt.Scanln(&angka2)

	// Latihan Perbandingan
	fmt.Println("\n=== Hasil Perbandingan ===")
	fmt.Printf("%d == %d : %v\n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d != %d : %v\n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d > %d : %v\n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d < %d : %v\n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d >= %d : %v\n", angka1, angka2, angka1 >= angka2)
	fmt.Printf("%d <= %d : %v\n", angka1, angka2, angka1 <= angka2)

	// ! Operator Logika
	/* and - or - not
	1.	(1>2) && (2<3)
	2.	(1>2) || (2<3)
	3.	!(1>2)
	*/
}
