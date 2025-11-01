package main

import "fmt"

func main() {
	fmt.Println(nama_lengkap)
	fmt.Println(usia)
	fmt.Println(ipk)
	fmt.Println(pintar)
	printVariabelTanpaValue()
	printKonstanta()
	printKonversi()
	CobaString()
	KondisionalString()
	pakaiStruct()
}

//  sintaks penggunaan struct

type Pegawai struct {
	id int
	name string
	age int
	gender bool
}

func pakaiStruct(){
	data_rafi := Pegawai{
		id: 1,
		name: "Rafi Mahrus",
		age: 22,
		gender: false,
	}
	fmt.Printf("Data Pegawai: %v \n =======\n", data_rafi)
	fmt.Printf("Saya adalah %v dengan usia %v\n", data_rafi.name, data_rafi.age)
}


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

func printVariabelTanpaValue(){
	fmt.Println("Variabel tanpa value:")
	fmt.Println(kota)
	fmt.Println(tinggi_badan)
	fmt.Println(sudah_makan)
	fmt.Println(pnl)
}

// deklarasi variabel singkat * hanya bisa didalam fungsi

func DeklarasiVariabel(){
	fmt.Println("Deklarasi variabel:")
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

func printKonstanta(){
	fmt.Println("Print Konstanta:")
	fmt.Println(phi)
	fmt.Println(nama_wali)
}

// konversi

var a int = 10
var b int32 = 32

var c int = int(b)

func printKonversi(){
	fmt.Println("Print konversi :")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}


// string di go adalah urutan byte
// string immutable (tidak bisa diubah)

func CobaString(){
	fmt.Println("Coba String:")
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
	fmt.Println("Kondisional String:")
	a := false

	message := getMessage(a)
	fmt.Println(message)
}