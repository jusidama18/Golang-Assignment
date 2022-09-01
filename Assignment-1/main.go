package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	opt, _ := strconv.Atoi(os.Args[1]) // Convert String to Integer
	opt = int(opt)
	if opt <= 0 {
		fmt.Println("Args Mulai dari 1")
		return
	}
	resp, err := search(opt)
	if err != nil {
		fmt.Println(err) // Jika error tidak nil, maka akan print error
	} else {
		// Jika error nil, maka jalankan method
		fmt.Println(resp.parse())
	}
}

// Struct Member yang berisi Nama, Alamat, Pekerjaan, dan Alasan dengan tipe data string
type Member struct {
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}

func (mem *Member) parse() string {
	// Menambah Method untuk membuat tampilan data lebih enak dilihat.
	result := "[ Biodata ]\n" + strings.Repeat("=", 50) + "\n"
	result += fmt.Sprintln(`Nama		: 	`, mem.Nama)
	result += fmt.Sprintln(`Alamat 		: 	`, mem.Alamat)
	result += fmt.Sprintln(`Pekerjaan 	: 	`, mem.Pekerjaan)
	result += fmt.Sprintln(`Alasan 		:	`, mem.Alasan)
	result += strings.Repeat("=", 50)
	return result
}

func search(opt int) (Member, error) {
	// List Data Member dalam bentuk Struct
	data := []Member{
		{
			Nama: "Juan",
			Alamat: "Jl. Kapten Satu",
			Pekerjaan: "Mahasiswa",
			Alasan: "Saya ingin belajar hal baru",
		},
		{
			Nama: "John",
			Alamat: "Jl. Kapten Dua",
			Pekerjaan: "Penulis",
			Alasan: "Saya ingin belajar Golang",

		},
		{
			Nama: "Pablo",
			Alamat: "Jl. Kapten Tiga",
			Pekerjaan: "Marketing",
			Alasan: "Saya ingin belajar bahasa baru",

		},
		{
			Nama: "Albert",
			Alamat: "Jl. Kapten Empat",
			Pekerjaan: "Dokter",
			Alasan: "Saya merasa Golang menarik",

		},
		{
			Nama: "Toni",
			Alamat: "Jl. Kapten Lima",
			Pekerjaan: "Pegawai Negeri",
			Alasan: "Saya ingin belajar pemrograman",

		},
		{
			Nama: "Bobi",
			Alamat: "Jl. Kapten Enam",
			Pekerjaan: "Dosen",
			Alasan: "Saya melihat golang menarik",

		},
		{
			Nama: "Boris",
			Alamat: "Jl. Kapten Tujuh",
			Pekerjaan: "Musisi",
			Alasan: "Saya mudah menggunakan Golang",

		},
		{
			Nama: "Jacky",
			Alamat: "Jl. Kapten Delapan",
			Pekerjaan: "Berkebun",
			Alasan: "Saya belajar Golang sebagai hal",

		},
	}
	
	if opt > len(data) {
		// Return Data Kosong Karena Error
		mem := Member{
			Nama:      "",
			Alamat:    "",
			Pekerjaan: "",
			Alasan:    "",
		}
		return mem, fmt.Errorf("index : %d , not found ! max data only : %d", opt, len(data))
	}
	// Kurang satu karena Index mulai dari 0
	return data[opt - 1], nil
}