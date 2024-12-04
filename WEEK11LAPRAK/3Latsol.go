package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bilangan)

	switch {
	case bilangan%10 == 0:
		fmt.Println("Kategori : Bilangan kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, bilangan/10)

	case bilangan%5 == 0 && bilangan != 5:
		fmt.Println("Kategori : Bilangan kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", bilangan, bilangan*bilangan)

	case bilangan%2 == 1:
		fmt.Println("Kategori : Bilangan ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, bilangan+(bilangan+1))

	case bilangan%2 == 0 && bilangan%5 != 0 && bilangan%10 != 0:
		fmt.Println("Kategori : Bilangan genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, bilangan*(bilangan+1))

	default:
		fmt.Println("Tidak ada bilangan yang cocok")
	}
}
