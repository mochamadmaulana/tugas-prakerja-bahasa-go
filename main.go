package main

import (
	"fmt"
	"math"
)

func main() {
	//variable()
	//aritmatika()
	//logika()
	perulangan()

	// ------------------------ Bilangan Prima -------------------------------------
	//var num int
	//fmt.Print("Masukkan sebuah bilangan: ")
	//fmt.Scan(&num)
	//if isPrime(num) {
	//	fmt.Printf("%d adalah bilangan prima.\n", num)
	//} else {
	//	fmt.Printf("%d bukan bilangan prima.\n", num)
	//}
	// ------------------------ End Bilangan Prima -------------------------------------

	// ------------------------ Kelipatan 7 -------------------------------------
	//var num int
	//fmt.Print("Masukkan sebuah bilangan: ")
	//fmt.Scan(&num)

	//if kelipatan7(num) {
	//	fmt.Printf("%d adalah kelipatan 7.\n", num)
	//} else {
	//	fmt.Printf("%d bukan kelipatan 7.\n", num)
	//}
	// ------------------------ Kelipatan 7 -------------------------------------
}

func variable() {
	// Variable menggunakan kata kunci var
	var nama string
	var umur int8 = 24

	// Variable tanpa kata kunci harus berformat :=
	universitas := "Universitas Raharja"

	// Multi variable
	var (
		tempatLahir = "Sumedang"
		domisili    = "Tangerang"
	)

	nama = "Mochamad Maulana"
	fmt.Println("Nama :", nama)
	fmt.Println("Usia :", umur, "Tahun")
	fmt.Println("Sekolah :", universitas)
	fmt.Println("Tempat Lahir | Domisili :", tempatLahir, "|", domisili)
}

func aritmatika() {
	// Artimatika biasa +	-	*	/	%
	var a = 10
	var b = 10
	var c = a * b // 10 * 10
	fmt.Println(c)

	// Augmented Aritmatika +=	-=	*=	/+	 %=
	c *= a // c = 100 * 10
	fmt.Println(c)

	//	Unary Aritmatika ++	--	-+	!
	c++ // 1000 + 1
	fmt.Println(c)
}

func logika() {
	var (
		a     = 5
		b     = 10
		hasil = 0
	)
	if a < b {
		hasil = b - a
	} else {

	}
	fmt.Println(hasil)
}

func perulangan() {
	// Perulangan FOR
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	maxDivisor := int(math.Sqrt(float64(n)))

	for i := 3; i <= maxDivisor; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func kelipatan7(number int) bool {
	return number%7 == 0
}
