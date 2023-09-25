package main

import "fmt"

func main() {
	//variable()
	//aritmatika()
	//logika()
	perulangan()
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
