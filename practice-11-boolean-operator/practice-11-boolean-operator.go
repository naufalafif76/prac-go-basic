package main

import "fmt"

func main() {
	var nilaiUjian = 80
	var absen = 80

	var lulusAkhir = nilaiUjian >= 60
	var lulusAbsensi = absen >= 80

	var lulus = lulusAkhir && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(nilaiUjian > 60 && absen > 80)

	// || (atau)
	// && (dan)
	// ! (tidak)
}
