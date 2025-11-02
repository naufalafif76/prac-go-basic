package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// slice1[0] = "Mei Lagi" // mengubah slice akan mengubah array
	// fmt.Println(months)

	// slice2 := months[1:5]
	slice2 := months[10:]
	fmt.Println(slice2)

	slices3 := append(slice2, "bulanBaru")
	fmt.Println(slices3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Astral"
	newSlice[1] = "Express"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice)) // panjang slice baru akan terpotong apabila panjang slice lama tidak cukup
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniSlice := []int{1, 2, 3, 4, 5}
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniArray2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(iniSlice)
	fmt.Println(iniArray)
	fmt.Println(iniArray2)

}
