package main

/*
Struct – tipe data di Go yang mengelompokkan beberapa properti menjadi satu kesatuan, seperti blueprint untuk objek.
🔹 Menggabungkan berbagai properti yang saling berhubungan dalam satu tempat.
🔹 Membuat cetak biru (blueprint) untuk membuat banyak objek dengan pola yang sama.
🔹 Menulis kode yang lebih terstruktur dan mudah dipelihara.
*/
import "fmt"

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func main() {
	pp := PersegiPanjang{Panjang: 10, Lebar: 5}
	fmt.Println(pp.Luas())
	fmt.Println(pp.Keliling())
}
