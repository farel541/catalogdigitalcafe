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

type Pelanggan struct {
	username string
	saldo    float64
}

type DaftarPelanggan struct {
	data            [NMAX]Pelanggan
	BanyakPelanggan int
}

func main() {
	var TabMenu DaftarMenu
	TabMenu.BanyakMenu = 0

	var TabPelanggan DaftarPelanggan
	TabPelanggan.BanyakPelanggan = 0

	var username string
	var pilihan int
	var running bool = true

	// Data Awal Menu
	TabMenu.data[0] = menu{nama: "Espresso", harga: 20000, kategori: "Coffee"}
	TabMenu.data[1] = menu{nama: "Cafe_Latte", harga: 28000, kategori: "Coffee"}
	TabMenu.data[2] = menu{nama: "Matcha_Latte", harga: 25000, kategori: "Non-Coffee"}
	TabMenu.data[3] = menu{nama: "Ice_Lemon_Tea", harga: 15000, kategori: "Non-Coffee"}
	TabMenu.data[4] = menu{nama: "Croissant", harga: 22000, kategori: "Food"}
	TabMenu.data[5] = menu{nama: "French_Fries", harga: 18000, kategori: "Food"}
	TabMenu.BanyakMenu = 6

	// Data Awal Pelanggan
	TabPelanggan.data[0] = Pelanggan{username: "budi", saldo: 75000}
	TabPelanggan.data[1] = Pelanggan{username: "andi", saldo: 15000}
	TabPelanggan.BanyakPelanggan = 2

	// Loop utama tanpa break/continue
	for running {
		fmt.Println("===========================================")
		fmt.Println("=== SELAMAT DATANG DI APLIKASI KATALOG DIGITAL CAFE ===")
		fmt.Println("===========================================")
		fmt.Print("Masukkan Username Anda (ketik 'exit' untuk selesai): ")
		fmt.Scan(&username)

		if username == "exit" {
			running = false
		} else if username == "admin" {
			fmt.Printf("\nHalo Admin! Anda masuk ke Mode Manajemen Menu.\n")
			loopAdmin := true
			for loopAdmin {
				fmt.Println("=== MENU ADMIN ===")
				fmt.Println("1. Tambah Menu")
				fmt.Println("2. Hapus Menu (Binary Search)")
				fmt.Println("3. Tampilkan Menu Terurut Harga (Selection Sort)")
				fmt.Println("4. Keluar (Log Out)")
				fmt.Print("Pilih opsi (1-4): ")
				fmt.Scan(&pilihan)

				if pilihan == 1 {
					tambahMenu(&TabMenu)
				} else if pilihan == 2 {
					hapusMenuBinary(&TabMenu)
				} else if pilihan == 3 {
					menuSortingTampilan(TabMenu)
				} else if pilihan == 4 {
					fmt.Println("\nKeluar dari sistem admin...\n")
					loopAdmin = false // Mengganti break
				} else {
					fmt.Println("Pilihan tidak valid!\n")
				}
			}
		} else {
			idxUser := cariPelangganSequential(&TabPelanggan, username)
			fmt.Printf("\nHalo %s! Selamat datang di CATALOG DIGITAL CAFE kami.\n", username)
			loopPelanggan := true
			for loopPelanggan {
				fmt.Println("=== MENU PELANGGAN ===")
				fmt.Println("1. Cek Saldo")
				fmt.Println("2. Top Up Saldo")
				fmt.Println("3. Tampilkan Menu Berdasarkan Kategori")
				fmt.Println("4. Pesan Menu")
				fmt.Println("5. Keluar (Log Out)")
				fmt.Print("Pilih opsi (1-5): ")
				fmt.Scan(&pilihan)

				if pilihan == 1 {
					fmt.Printf("\nSaldo Anda (%s) saat ini: Rp %.2f\n\n", username, TabPelanggan.data[idxUser].saldo)
				} else if pilihan == 2 {
					var topUp float64
					fmt.Print("Masukkan nominal Top Up: Rp ")
					fmt.Scan(&topUp)
					if topUp > 0 {
						TabPelanggan.data[idxUser].saldo += topUp
						fmt.Println("Top Up berhasil!\n")
					} else {
						fmt.Println("Nominal tidak valid!\n")
					}
				} else if pilihan == 3 {
					TampilkanMenuSesuaiKategoriPelanggan(TabMenu)
				} else if pilihan == 4 {
					pesanMenu(&TabMenu, &TabPelanggan.data[idxUser].saldo)
				} else if pilihan == 5 {
					fmt.Println("\nKeluar dari akun pelanggan...\n")
					loopPelanggan = false // Mengganti break
				} else {
					fmt.Println("Pilihan tidak valid!\n")
				}
			}
		}
	}
}

// ================= SUBPROGRAM DENGAN SPESIFIKASI BERPARAMETER =================

/*
Spesifikasi Subprogram:
IS: Terdefinisi pointer TabPelanggan berkapasitas NMAX dan string nama.
FS: Mengembalikan indeks posisi nama pelanggan jika ditemukan via Sequential Search. Jika tidak ditemukan, data baru didaftarkan otomatis dan mengembalikan indeks baru.
*/
func cariPelangganSequential(TabPelanggan *DaftarPelanggan, nama string) int {
	var i int = 0
	var ketemu int = -1

	for i < TabPelanggan.BanyakPelanggan && ketemu == -1 {
		if TabPelanggan.data[i].username == nama {
			ketemu = i
		}
		i++
	}

	if ketemu != -1 {
		return ketemu
	}

	if TabPelanggan.BanyakPelanggan < NMAX {
		idxBaru := TabPelanggan.BanyakPelanggan
		TabPelanggan.data[idxBaru].username = nama
		TabPelanggan.data[idxBaru].saldo = 0
		TabPelanggan.BanyakPelanggan++
		return idxBaru
	}
	return 0
}

/*
Spesifikasi Subprogram:
IS: Terdefinisi pointer TabMenu berisi sejumlah BanyakMenu data.
FS: Menambahkan satu elemen menu baru di akhir array jika memori masih cukup.
*/
func tambahMenu(TabMenu *DaftarMenu) {
	if TabMenu.BanyakMenu < NMAX {
		fmt.Print("Masukkan Nama Menu (tanpa spasi): ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].nama)
		fmt.Print("Masukkan Harga Menu: ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].harga)
		fmt.Print("Masukkan Kategori Menu: ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].kategori)

		TabMenu.BanyakMenu++
		fmt.Println("SUKSES: MENU BERHASIL DITAMBAHKAN\n")
	} else {
		fmt.Println("GAGAL: KAPASITAS KATALOG SUDAH PENUH\n")
	}
}

/*
Spesifikasi Subprogram:
IS: Terdefinisi data TabMenu terisi.
FS: Mengurutkan menu internal secara alfabetis (A-Z) khusus sebagai syarat jalan Binary Search.
*/
func urutkanMenuBerdasarNamaAsc(TabMenu *DaftarMenu) {
	var i, j, idxMin int
	var tukar menu
	for i = 0; i < TabMenu.BanyakMenu-1; i++ {
		idxMin = i
		for j = i + 1; j < TabMenu.BanyakMenu; j++ {
			if TabMenu.data[j].nama < TabMenu.data[idxMin].nama {
				idxMin = j
			}
		}
		tukar = TabMenu.data[i]
		TabMenu.data[i] = TabMenu.data[idxMin]
		TabMenu.data[idxMin] = tukar
	}
}

/*
Spesifikasi Subprogram (MEMENUHI ATURAN BINARY SEARCH UNTUK PENGHAPUSAN):
IS: Terdefinisi pointer TabMenu terisi data acak/urut.
FS: Mencari data berdasarkan nama menu menggunakan Binary Search, jika ketemu, elemen dihapus dengan cara menggeser array.
*/
func hapusMenuBinary(TabMenu *DaftarMenu) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("Tidak ada menu yang bisa dihapus.\n")
		return
	}

	urutkanMenuBerdasarNamaAsc(TabMenu) // Wajib diurutkan sebelum Binary Search

	fmt.Println("\n=== DAFTAR MENU SAAT INI (URUT ABJAD) ===")
	TampilkanMenu(*TabMenu)

	var namaCari string
	fmt.Print("Masukkan NAMA menu yang ingin dihapus: ")
	fmt.Scan(&namaCari)

	// LOGIKA BINARY SEARCH
	kr := 0
	kn := TabMenu.BanyakMenu - 1
	indexHapus := -1

	for kr <= kn && indexHapus == -1 {
		med := (kr + kn) / 2
		if TabMenu.data[med].nama == namaCari {
			indexHapus = med
		} else if TabMenu.data[med].nama < namaCari {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}

	if indexHapus == -1 {
		fmt.Println("GAGAL: Nama menu tidak ditemukan!\n")
		return
	}

	// Penghapusan dengan pergeseran array
	for i := indexHapus; i < TabMenu.BanyakMenu-1; i++ {
		TabMenu.data[i] = TabMenu.data[i+1]
	}
	TabMenu.BanyakMenu--
	fmt.Println("SUKSES: MENU BERHASIL DIHAPUS VIA BINARY SEARCH\n")
}

/*
Spesifikasi Subprogram:
IS: Terdefinisi data TabMenu terisi.
FS: Menampilkan daftar seluruh data menu ke layar komputer secara tabular.
*/
func TampilkanMenu(TabMenu DaftarMenu) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("\n--- Menu Kosong ---\n")
		return
	}
	fmt.Printf("\n%-3s | %-20s | %-12s | %-10s\n", "No", "Nama Menu", "Kategori", "Harga")
	fmt.Println("---------------------------------------------------------")
	for i := 0; i < TabMenu.BanyakMenu; i++ {
		fmt.Printf("%-3d | %-20s | %-12s | Rp %-8.2f\n", i+1, TabMenu.data[i].nama, TabMenu.data[i].kategori, TabMenu.data[i].harga)
	}
	fmt.Println()
}

/*
Spesifikasi Subprogram (MEMENUHI ATURAN SORTING SELECTION & INSERTION DENGAN ASC/DESC):
IS: Terdefinisi salinan data TabMenu.
FS: Mengurutkan menu berdasarkan harga menggunakan Selection Sort atau Insertion Sort, baik secara Ascending atau Descending sesuai input user.
*/
func menuSortingTampilan(TabMenu DaftarMenu) {
	var algoritma, jenisUrutan int
	fmt.Println("\nPilih Algoritma Pengurutan Harga:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilihan (1-2): ")
	fmt.Scan(&algoritma)

	fmt.Println("Pilih Urutan:")
	fmt.Println("1. Naik (Ascending)")
	fmt.Println("2. Turun (Descending)")
	fmt.Print("Pilihan (1-2): ")
	fmt.Scan(&jenisUrutan)

	if algoritma == 1 {
		// SELECTION SORT
		for i := 0; i < TabMenu.BanyakMenu-1; i++ {
			idxTarget := i
			for j := i + 1; j < TabMenu.BanyakMenu; j++ {
				if jenisUrutan == 1 && TabMenu.data[j].harga < TabMenu.data[idxTarget].harga {
					idxTarget = j
				} else if jenisUrutan == 2 && TabMenu.data[j].harga > TabMenu.data[idxTarget].harga {
					idxTarget = j
				}
			}
			tukar := TabMenu.data[i]
			TabMenu.data[i] = TabMenu.data[idxTarget]
			TabMenu.data[idxTarget] = tukar
		}
	} else {
		// INSERTION SORT
		for i := 1; i < TabMenu.BanyakMenu; i++ {
			key := TabMenu.data[i]
			j := i - 1
			if jenisUrutan == 1 {
				for j >= 0 && TabMenu.data[j].harga > key.harga {
					TabMenu.data[j+1] = TabMenu.data[j]
					j--
				}
			} else {
				for j >= 0 && TabMenu.data[j].harga < key.harga {
					TabMenu.data[j+1] = TabMenu.data[j]
					j--
				}
			}
			TabMenu.data[j+1] = key
		}
	}
	TampilkanMenu(TabMenu)
}

func TampilkanMenuPerKategori(TabMenu DaftarMenu, kategoriPilihan string) [NMAX]int {
	var indeksAsli [NMAX]int
	jumlahFilter := 0

	fmt.Printf("\n--- Daftar Menu Kategori: %s ---\n", kategoriPilihan)
	fmt.Printf("%-3s | %-20s | %-10s\n", "No", "Nama Menu", "Harga")
	fmt.Println("----------------------------------------")

	for i := 0; i < TabMenu.BanyakMenu; i++ {
		if TabMenu.data[i].kategori == kategoriPilihan {
			indeksAsli[jumlahFilter] = i
			jumlahFilter++
			fmt.Printf("%-3d | %-20s | Rp %-8.2f\n", jumlahFilter, TabMenu.data[i].nama, TabMenu.data[i].harga)
		}
	}
	fmt.Println()
	indeksAsli[NMAX-1] = jumlahFilter
	return indeksAsli
}

func TampilkanMenuSesuaiKategoriPelanggan(TabMenu DaftarMenu) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("Belum ada menu yang tersedia.\n")
		return
	}

	var daftarKategori [NMAX]string
	banyakKategori := 0
	for i := 0; i < TabMenu.BanyakMenu; i++ {
		sudahAda := false
		for j := 0; j < banyakKategori; j++ {
			if TabMenu.data[i].kategori == daftarKategori[j] {
				sudahAda = true
			}
		}
		if !sudahAda {
			daftarKategori[banyakKategori] = TabMenu.data[i].kategori
			banyakKategori++
		}
	}

	fmt.Println("\n=== PILIH KATEGORI YANG INGIN DITAMPILKAN ===")
	for i := 0; i < banyakKategori; i++ {
		fmt.Printf("%d. %s\n", i+1, daftarKategori[i])
	}

	var pilihKat int
	fmt.Print("Pilih nomor kategori: ")
	fmt.Scan(&pilihKat)

	if pilihKat >= 1 && pilihKat <= banyakKategori {
		TampilkanMenuPerKategori(TabMenu, daftarKategori[pilihKat-1])
	} else {
		fmt.Println("Pilihan kategori tidak valid!\n")
	}
}

func pesanMenu(TabMenu *DaftarMenu, saldo *float64) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("Belum ada menu yang tersedia.\n")
		return
	}

	var daftarKategori [NMAX]string
	banyakKategori := 0

	for i := 0; i < TabMenu.BanyakMenu; i++ {
		sudahAda := false
		for j := 0; j < banyakKategori; j++ {
			if TabMenu.data[i].kategori == daftarKategori[j] {
				sudahAda = true
			}
		}
		if !sudahAda {
			daftarKategori[banyakKategori] = TabMenu.data[i].kategori
			banyakKategori++
		}
	}

	fmt.Println("\n=== PILIH KATEGORI MENU ===")
	for i := 0; i < banyakKategori; i++ {
		fmt.Printf("%d. %s\n", i+1, daftarKategori[i])
	}

	var pilihKat int
	fmt.Print("Pilih nomor kategori: ")
	fmt.Scan(&pilihKat)

	if pilihKat < 1 || pilihKat > banyakKategori {
		fmt.Println("Pilihan kategori tidak valid!\n")
		return
	}

	kategoriTerpilih := daftarKategori[pilihKat-1]
	petaIndeks := TampilkanMenuPerKategori(*TabMenu, kategoriTerpilih)
	jumlahMenuFilter := petaIndeks[NMAX-1]

	if jumlahMenuFilter == 0 {
		fmt.Println("Tidak ada menu untuk kategori ini.\n")
		return
	}

	var nomor int
	fmt.Print("Pilih nomor menu yang ingin dipesan: ")
	fmt.Scan(&nomor)

	if nomor >= 1 && nomor <= jumlahMenuFilter {
		indeksAsliMenu := petaIndeks[nomor-1]
		menuDipilih := TabMenu.data[indeksAsliMenu]

		if *saldo >= menuDipilih.harga {
			*saldo -= menuDipilih.harga
			fmt.Printf("Berhasil memesan %s! Saldo Anda terpotong Rp %.2f\n\n", menuDipilih.nama, menuDipilih.harga)
		} else {
			fmt.Println("Gagal memesan! Saldo Anda tidak mencukupi. Silakan Top Up.\n")
		}
	} else {
		fmt.Println("Pilihan nomor menu tidak valid!\n")
	}
}
