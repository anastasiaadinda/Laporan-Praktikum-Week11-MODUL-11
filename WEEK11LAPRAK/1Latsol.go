package main

import "fmt"

func main() {
	var pH float64
	fmt.Print("Masukan Nilai pH air : ")
	fmt.Scan(&pH)

	switch {
	case pH < 0 || pH > 14:
		fmt.Println("Nilai ph tidak valid, masukan antara 0-14")
	case pH >= 6.5 && pH <= 8.6:
		fmt.Println("Air layak minum")
	default:
		fmt.Println("Air tidak layak minum")
	}
}
