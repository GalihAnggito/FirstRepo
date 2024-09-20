package main

import (
	"fmt"
	"os"
)

type User struct {
	Username string
	Password string
}

type ParkingTicket struct {
	VehicleType string
	PlatNumber  string
	EntryTime   string
	ExitTime    string
	Amount      float64
}

type Admin struct {
	Username string
	Password string
}

var loggedInUser User
var admin Admin
var parkingTickets []ParkingTicket
var parkingAttendants []User

func main() {
	admin = Admin{"admin", "admin123"}

	fmt.Println("*** ------------------------------------------- ***")
	fmt.Println("***            Aplikasi Parkir                  ***")
	fmt.Println("***               Created by :                  ***")
	fmt.Println("***      - Galih Anggito Abimanyu               ***")
	fmt.Println("***      - Aryo Jati Pamungkas                  ***")
	fmt.Println("***    Tugas Besar Algoritma Pemrograman 2024   ***")
	fmt.Println("*** ------------------------------------------- ***")
	fmt.Println("*** Menu Utama ***")
	fmt.Println("1. Login")
	fmt.Println("2. Admin")
	fmt.Println("3. Keluar")
	fmt.Println("------------------")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		login()
	case 2:
		adminMenu()
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid.")
		main()
	}
}

func login() {
	var username, password string
	fmt.Println("*** Login ***")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if username == admin.Username && password == admin.Password {
		fmt.Println("Login berhasil sebagai admin.")
		adminMenu()
	} else {
		found := false
		i := 0
		for !found && i < len(parkingAttendants) {
			if parkingAttendants[i].Username == username && parkingAttendants[i].Password == password {
				found = true
				loggedInUser = parkingAttendants[i]
			}
			i++
		}
		if found {
			fmt.Println("Login berhasil sebagai petugas tiket parkir.")
			petugasMenu()
		} else {
			fmt.Println("Username atau Password salah.")
			var crn string
			fmt.Println("Kembali ke menu awal ?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Println("Pilih 1 / 2")
			fmt.Scan(&crn)
			if crn == "1" {
				main()
			}
		}
	}
}

func adminMenu() {
	fmt.Println("*** Menu Admin ***")
	fmt.Println("1. Tambah Petugas Parkir")
	fmt.Println("2. Lihat Daftar Petugas Parkir")
	fmt.Println("3. Hapus Petugas Parkir")
	fmt.Println("4. Lihat Daftar Kendaraan dan Pendapatan")
	fmt.Println("5. Kembali ke Menu Utama")
	fmt.Println("------------------")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addParkingAttendant()
	case 2:
		viewParkingAttendants()
	case 3:
		deleteParkingAttendant()
	case 4:
		viewParkingTickets()
	case 5:
		main()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func addParkingAttendant() {
	fmt.Println("*** Tambah Petugas Parkir ***")
	var username, password string

	fmt.Print("Masukkan username baru: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password baru: ")
	fmt.Scanln(&password)

	newAttendant := User{Username: username, Password: password}
	parkingAttendants = append(parkingAttendants, newAttendant)

	fmt.Println("Petugas parkir dengan username", username, "telah ditambahkan.")

	var crn string
	fmt.Println("Press x to continue...")
	fmt.Scan(&crn)
	if crn == "x" {
		main()
	}
}

func viewParkingAttendants() {
	fmt.Println("*** Daftar Petugas Parkir ***")
	if len(parkingAttendants) == 0 {
		fmt.Println("Belum ada petugas parkir terdaftar.")
	} else {
		for i, attendant := range parkingAttendants {
			fmt.Printf("%d. Username: %s\n", i+1, attendant.Username)
		}
	}
	fmt.Println("------------------")
	var crn string
	fmt.Println("Press x to continue...")
	fmt.Scan(&crn)
	if crn == "x" {
		main()
	}
}

func deleteParkingAttendant() {
	fmt.Println("*** Hapus Petugas Parkir ***")
	fmt.Print("Masukkan username petugas parkir yang ingin dihapus: ")
	var username string
	fmt.Scanln(&username)

	found := false
	for i, attendant := range parkingAttendants {
		if attendant.Username == username {
			found = true
			parkingAttendants = append(parkingAttendants[:i], parkingAttendants[i+1:]...)
			fmt.Println("Petugas parkir dengan username", username, "telah dihapus.")
			break
		}
	}

	if !found {
		fmt.Println("Petugas parkir dengan username", username, "tidak ditemukan.")
	}

	var crn string
	fmt.Println("Press x to continue...")
	fmt.Scan(&crn)
	if crn == "x" {
		main()
	}
}

func viewParkingTickets() {
	fmt.Println("*** Daftar Kendaraan dan Pendapatan ***")
	if len(parkingTickets) == 0 {
		fmt.Println("Belum ada transaksi parkir.")
	} else {
		totalPendapatan := 0.0
		for _, ticket := range parkingTickets {
			fmt.Printf("Jenis Kendaraan: %s, Nomor Plat: %s, Tarif: Rp.%0.2f\n", ticket.VehicleType, ticket.PlatNumber, ticket.Amount)
			totalPendapatan += ticket.Amount
		}
		fmt.Printf("Total Pendapatan Hari Ini: Rp.%0.2f\n", totalPendapatan)
	}
	fmt.Println("------------------")
	var crn string
	fmt.Println("Press x to continue...")
	fmt.Scan(&crn)
	if crn == "x" {
		main()
	}
}

func petugasMenu() {
	fmt.Println("*** Menu Petugas ***")
	fmt.Println("1. Input Transaksi Parkir")
	fmt.Println("2. Lihat Daftar Kendaraan dan Pendapatan")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println("------------------")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		inputParkingTransaction()
	case 2:
		viewParkingTickets()
	case 3:
		main()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func inputParkingTransaction() {
	var jenis, plat string
	var tarif, durasi int

	fmt.Print("Masukkan jenis kendaraan (motor / mobil) : ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan nomor kendaraan : ")
	fmt.Scan(&plat)

	if jenis == "mobil" {
		fmt.Println("60 menit pertama Rp.4000 ")
		fmt.Println("Setelahnya Rp.2000 per 60 menit berikutnya.")
		fmt.Print("Masukkan durasi parkir : ")
		fmt.Scan(&durasi)
		if durasi <= 60 {
			tarif = 4000
		} else if durasi > 60 {
			if durasi-60 <= 60 {
				tarif = (2000 * (60 / (durasi - 60))) + 4000
			} else {
				tarif = (2000 * ((durasi - 60) / 60)) + 4000
			}
		}
	} else if jenis == "motor" {
		fmt.Println("60 menit pertama Rp.2000 ")
		fmt.Println("Setelahnya Rp.1000 per 60 menit berikutnya.")
		fmt.Print("Masukkan durasi parkir : ")
		fmt.Scan(&durasi)
		if durasi <= 60 {
			tarif = 2000
		} else if durasi > 60 {
			if durasi-60 <= 60 {
				tarif = (1000 * (60 / (durasi - 60))) + 2000
			} else {
				tarif = (1000 * ((durasi - 60) / 60)) + 2000
			}
		}
	}

	parkingTicket := ParkingTicket{
		VehicleType: jenis,
		PlatNumber:  plat,
		Amount:      float64(tarif),
	}

	parkingTickets = append(parkingTickets, parkingTicket)

	fmt.Println("Total tarif parkir :", tarif)
	fmt.Println("Transaksi parkir berhasil dicatat.")
	var crn string
	fmt.Println("Press x to continue...")
	fmt.Scan(&crn)
	if crn == "x" {
		main()
	}
}
