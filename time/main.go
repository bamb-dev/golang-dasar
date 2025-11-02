package main

import (
	"fmt"
	"time"
)

func convertTimeZone(t time.Time, zona string) time.Time {
	loc, err := time.LoadLocation(zona)
	if err != nil {
		fmt.Printf("Zona waktu %v tidak ditemukan\n", zona)
		return t
	}
	return t.In(loc)
}

func main() {
	nowUTC := time.Now().UTC()
	fmt.Println("Waktu UTC : ", nowUTC)

	nowJakarta := convertTimeZone(nowUTC, "Asia/Jakarta")
	fmt.Println("Waktu Jakarta : ", nowJakarta)

	nowLondon := convertTimeZone(nowUTC, "Europe/London")
	fmt.Println("Waktu London : ", nowLondon)
	nowMakassar := convertTimeZone(nowUTC, "Asia/Makassar")
	fmt.Println("Waktu Makassar : ", nowMakassar)

}
