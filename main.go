package fiber_go

import "fmt"

// sintaks deklarasi variabel dengan tipe

var nama_lengkap string = "Rafi Mahrus"
var usia int = 22
var ipk float64 = 3.95
var pintar bool = true

// sintaks deklarasi tanpa value
var kota string
var tinggi_badan int
var sudah_makan bool
var pnl float64

// deklarasi variabel singkat * hanya bisa didalam fungsi

func DeklarasiVariabel(){
	kodepos := 546321
	berat_badan := 70.5
	pesan_pengguna := "Selamat datang di golang"

	println(kodepos)
	println(berat_badan)
	println(pesan_pengguna)
}


// deklarasi konstanta
const phi = 3.14
const nama_wali = "Rafi Mahrus"

// konversi

var a int = 10
var b int32 = 32

var c int = int(b)


// string di go adalah urutan byte
// string immutable (tidak bisa diubah)

func CobaString(){
	s1 := "Rafi"
	fmt.Println(s1[0])

	s2 := s1 + " Mahrus"
	fmt.Println(s2)
}

func getMessage(a bool) string {
	if a {
		return "Benar"
	}
	return  "Salah"
}

func KondisionalString() {
	a := false

	message := getMessage(a)
	fmt.Println(message)
}