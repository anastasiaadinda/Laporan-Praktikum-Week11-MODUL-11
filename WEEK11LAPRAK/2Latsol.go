package main

import "fmt"

func main() {
	var jenisKendaraan string
	var tarif float64
	var durasi int

	fmt.Print("Masukan jenis kendaraan : ")
	fmt.Scan(&jenisKendaraan)

	fmt.Print("Masukan durasi : ")
	fmt.Scan(&durasi)

	switch jenisKendaraan {
	case "motor":
		tarif = float64(durasi) * 2000
	case "mobil":
		tarif = float64(durasi) * 5000
	case "truk":
		tarif = float64(durasi) * 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}
	fmt.Println("Total biaya Rp", tarif)
}
