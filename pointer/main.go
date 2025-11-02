package main

/*
! pointer – fitur penting di Go yang digunakan untuk menyimpan alamat memori dari sebuah nilai.

🔹 Menghemat memori dengan tidak menduplikasi data besar.
🔹 Mengubah nilai variabel langsung dari fungsi lain.
🔹 Mengoptimalkan performa program dan membuat kode lebih efisien.
*/
import "fmt"

func heal(hp *int) {
	*hp += 20
	fmt.Println("Hp Bertambah 20")
}
func attact(hp *int, damage int) {
	*hp -= damage
	if *hp <= 0 {
		fmt.Println("Game Over")
	}
	fmt.Println("Hp Di Kurangi", damage)
}

func main() {
	hp := 100
	heal(&hp)
	attact(&hp, 20)
	attact(&hp, 60)
	attact(&hp, 80)
}
