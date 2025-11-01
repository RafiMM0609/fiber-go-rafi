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
	pakaiTagStruct()
	pakaiSlice()
	dataMap()
	pemainBola := SepakBola{}
	ManchesterUnited(pemainBola, true)
	pemainBasket := Basket{}
	ManchesterUnited(pemainBasket, true)
}

// penggunaan interface
type Pemain interface {
	Tendang(status bool) (string, error) 
}

type SepakBola struct {}
func (s SepakBola) Tendang(status bool) (string, error) {
	if status {
		return "Tendangan Berhasil", nil
	}
	return "Tendangan Gagal", nil
}

type Basket struct {}
func (b Basket) Tendang(status bool) (string, error){
	if status {return  "Pelanggaran, basket pakai tangan", nil}
	return "Tidak ada pelanggaran", nil
}

func ManchesterUnited(p Pemain, status bool){
	fmt.Println("Eksekusi manchester united")
	hasil, err := p.Tendang(status)
	if err != nil {fmt.Println("Manchester united error:", err)}
	fmt.Println("Ini hasil eksekusi :",hasil)
}

// penggunaan map dengan type any
func dataMap(){
	data_mahasiswa := map[string]any{
		"name": "Rafi Mahrus",
		"age": 22,
		"ipk": 3.95,
		"isGraduated": false,
	}
	fmt.Printf("Ini data mahasiswa : \n %v \n======\n", data_mahasiswa)
	fmt.Printf("Nama : %v \n", data_mahasiswa["name"])
}

// penggunaan slice
 var makanan = []string{"roti", "agar-agar", "nasi goreng", "sate"}

 func pakaiSlice(){
	// tambah item ke slice
	makanan = append(makanan, "bakso")
	fmt.Println("Pakai Slice:")
	fmt.Println("Item terakhir : ",makanan[len(makanan)-1])
}

//  sintaks penggunaan struct
type Pegawai struct {
	id int
	name string
	age int
	gender bool
}

// pakai tag di struct
type Mahasiswa struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Gender bool `json:"-"`
}

func pakaiTagStruct(){
	data_rafi := Mahasiswa{
		Id: 1,
		Name: "Rafi Mahrus",
		Age: 22,
		Gender: false,
	}
	fmt.Printf("Data Mahasiswa: %v \n =======\n", data_rafi)
	fmt.Printf("Saya adalah %v dengan usia %v\n", data_rafi.Name, data_rafi.Age)
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