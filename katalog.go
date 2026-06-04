package main

import "fmt"

const NMAX int = 100

type menu struct {
	nama     string
	harga    float64
	kategori string
}

type DaftarMenu struct {
	data       [NMAX]menu
	BanyakMenu int
}

func main() {
	var TabMenu DaftarMenu
	TabMenu.BanyakMenu = 0

	var pilihan int
	for {
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tambahMenu(&TabMenu)
			fmt.Println()
		} else if pilihan == 2 {
			fmt.Println("\n=== SEMUA DAFTAR MENU ===")
			TampilkanMenu(TabMenu)
		}
	}

}

func tambahMenu(TabMenu *DaftarMenu) {
	if TabMenu.BanyakMenu < NMAX {
		fmt.Println("Masukkan Nama Menu: ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].nama)
		fmt.Println("Masukkan Harga Menu: ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].harga)
		fmt.Println("Masukkan Kategori Menu: ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].kategori)

		TabMenu.BanyakMenu++
		fmt.Println("SUKSES MENU BERHASIL DITAMBAHKAN")
	} else {
		fmt.Println("GAGAL KAPASITAS KATALOG SUDAH PENUH")
	}
}

func TampilkanMenu(TabMenu DaftarMenu) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("--- Menu Kosong ---")
		return
	}
	fmt.Printf("\n%-3s | %-20s | %-12s | %-10s\n", "No", "Nama Menu", "Kategori", "Harga")
	fmt.Println("---------------------------------------------------------")
	for i := 0; i < TabMenu.BanyakMenu; i++ {
		fmt.Printf("%-3d | %-20s | %-12s | Rp %.2f\n", i+1, TabMenu.data[i].nama, TabMenu.data[i].kategori, TabMenu.data[i].harga)
	}
	fmt.Println()
}
