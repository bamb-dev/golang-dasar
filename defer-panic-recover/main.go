package main

import "fmt"

/*
  🔹 defer – Menjadwalkan fungsi agar dijalankan paling terakhir, bahkan ketika program keluar dari sebuah fungsi lebih awal.
  🔹 recover – Menangkap dan memulihkan program dari kondisi panic sehingga aplikasi tidak langsung berhenti.
  🔹 panic – Menghentikan eksekusi program ketika terjadi error kritis yang tidak bisa ditangani.
*/

func cleanup() {
	fmt.Println("CleanUP: Menutup resource...")
}

func bacaConfig(namaFile string) {
	defer cleanup()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Error:", r)
		}
	}()
	if namaFile == "" {
		panic("Nama file tidak boleh kosong")
	}
	fmt.Println("Membaca file configurasi dari : ", namaFile)
}

func main() {
	bacaConfig("")
	fmt.Println("dieksekusi setelah recover")
}
