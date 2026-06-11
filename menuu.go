package main

import "fmt"

const NMAX int = 100

type menu struct {
	nama     string
	harga    float64
	kategori string
	stok     int // Tambahan properti stok
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

	// Data Awal Menu (ditambah stok awal)
	TabMenu.data[0] = menu{nama: "Espresso", harga: 20000, kategori: "Coffee", stok: 10}
	TabMenu.data[1] = menu{nama: "Cafe_Latte", harga: 28000, kategori: "Coffee", stok: 15}
	TabMenu.data[2] = menu{nama: "Matcha_Latte", harga: 25000, kategori: "Non-Coffee", stok: 8}
	TabMenu.data[3] = menu{nama: "Ice_Lemon_Tea", harga: 15000, kategori: "Non-Coffee", stok: 20}
	TabMenu.data[4] = menu{nama: "Croissant", harga: 22000, kategori: "Food", stok: 5}
	TabMenu.data[5] = menu{nama: "French_Fries", harga: 18000, kategori: "Food", stok: 7}
	TabMenu.BanyakMenu = 6

	// Data Awal Pelanggan
	TabPelanggan.data[0] = Pelanggan{username: "budi", saldo: 75000}
	TabPelanggan.data[1] = Pelanggan{username: "andi", saldo: 15000}
	TabPelanggan.BanyakPelanggan = 2

	// Loop utama
	for running {
		fmt.Println("=======================================================")
		fmt.Println("=== SELAMAT DATANG DI APLIKASI KATALOG DIGITAL CAFE ===")
		fmt.Println("=======================================================")
		fmt.Print("Masukkan Username Anda (ketik 'exit' untuk selesai): ")
		fmt.Scan(&username)

		if username == "exit" {
			running = false
		} else if username == "admin" {
			fmt.Printf("\nHalo Admin!\n")
			loopAdmin := true
			for loopAdmin {
				fmt.Println("=== MENU ADMIN ===")
				fmt.Println("1. Tambah Menu")
				fmt.Println("2. Update Stok Menu")
				fmt.Println("3. Hapus Menu")
				fmt.Println("4. Tampilkan Menu Termahal")
				fmt.Println("5. Tampilkan Menu Termurah")
				fmt.Println("6. Keluar (Log Out)")
				fmt.Print("Pilih opsi (1-6): ")
				fmt.Scan(&pilihan)

				if pilihan == 1 {
					tambahMenu(&TabMenu)
				} else if pilihan == 2 {
					updateStokMenu(&TabMenu)
				} else if pilihan == 3 {
					hapusMenuBinary(&TabMenu)
				} else if pilihan == 4 {
					menuSortingTampilan(TabMenu, 1)
				} else if pilihan == 5 {
					menuSortingTampilan(TabMenu, 2)
				} else if pilihan == 6 {
					fmt.Println("\nKeluar dari sistem admin...\n")
					loopAdmin = false
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
				fmt.Println("4. Tampilkan Menu Termahal")
				fmt.Println("5. Tampilkan Menu Termurah")
				fmt.Println("6. Pesan Menu")
				fmt.Println("7. Keluar (Log Out)")
				fmt.Print("Pilih opsi (1-7): ")
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
					menuSortingTampilan(TabMenu, 1)
				} else if pilihan == 5 {
					menuSortingTampilan(TabMenu, 2)
				} else if pilihan == 6 {
					pesanMenu(&TabMenu, &TabPelanggan.data[idxUser].saldo)
				} else if pilihan == 7 {
					fmt.Println("\nKeluar dari akun pelanggan...\n")
					loopPelanggan = false
				} else {
					fmt.Println("Pilihan tidak valid!\n")
				}
			}
		}
	}
}

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
FS: Menambahkan satu elemen menu baru di akhir array.
    Ditambahkan opsi pembatalan menggunakan input teks "1".
*/
func tambahMenu(TabMenu *DaftarMenu) {
	if TabMenu.BanyakMenu < NMAX {
		var inputNama string
		fmt.Println("\n--- TAMBAH MENU BARU ---")
		fmt.Print("Masukkan Nama Menu (Gunakan underscore '_' untuk spasi, ketik '1' untuk BATAL): ")
		fmt.Scan(&inputNama)

		if inputNama == "1" {
			fmt.Println("Aksi dibatalkan. Kembali ke menu admin...\n")
			return
		}

		TabMenu.data[TabMenu.BanyakMenu].nama = inputNama

		fmt.Print("Masukkan Harga Menu: Rp ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].harga)
		
		fmt.Print("Masukkan Kategori Menu (Gunakan underscore '_' untuk spasi): ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].kategori)
		
		fmt.Print("Masukkan Jumlah Stok Awal: ")
		fmt.Scan(&TabMenu.data[TabMenu.BanyakMenu].stok)

		TabMenu.BanyakMenu++
		fmt.Println("SUKSES: MENU BERHASIL DITAMBAHKAN\n")
	} else {
		fmt.Println("GAGAL: KAPASITAS KATALOG SUDAH PENUH\n")
	}
}

/*
Fungsi Baru: updateStokMenu
FS: Memperbarui stok menu berdasarkan nama menu yang dicari menggunakan Binary Search.
    Bisa membatalkan proses dengan mengetikkan "1" saat pencarian nama.
*/
func updateStokMenu(TabMenu *DaftarMenu) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("Tidak ada menu untuk diupdate stoknya.\n")
		return
	}

	urutkanMenuBerdasarNamaAsc(TabMenu)

	fmt.Println("\n=== DAFTAR MENU SAAT INI (URUT ABJAD) ===")
	TampilkanMenu(*TabMenu)

	var namaCari string
	fmt.Print("Masukkan NAMA menu yang ingin diupdate stoknya (ketik '1' untuk BATAL): ")
	fmt.Scan(&namaCari)

	if namaCari == "1" {
		fmt.Println("Aksi dibatalkan. Kembali ke menu admin...\n")
		return
	}

	kr := 0
	kn := TabMenu.BanyakMenu - 1
	indexKetemu := -1

	for kr <= kn && indexKetemu == -1 {
		med := (kr + kn) / 2
		if TabMenu.data[med].nama == namaCari {
			indexKetemu = med
		} else if TabMenu.data[med].nama < namaCari {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}

	if indexKetemu == -1 {
		fmt.Println("GAGAL: Nama menu tidak ditemukan!\n")
		return
	}

	var stokBaru int
	fmt.Printf("Stok saat ini untuk %s: %d\n", TabMenu.data[indexKetemu].nama, TabMenu.data[indexKetemu].stok)
	fmt.Print("Masukkan jumlah STOK BARU (ketik '-1' untuk BATAL): ")
	fmt.Scan(&stokBaru)

	if stokBaru == -1 {
		fmt.Println("Aksi dibatalkan. Kembali ke menu admin...\n")
		return
	}

	if stokBaru >= 0 {
		TabMenu.data[indexKetemu].stok = stokBaru
		fmt.Println("SUKSES: STOK BERHASIL DIPERBARUI\n")
	} else {
		fmt.Println("GAGAL: Jumlah stok tidak valid!\n")
	}
}

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

func hapusMenuBinary(TabMenu *DaftarMenu) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("Tidak ada menu yang bisa dihapus.\n")
		return
	}

	urutkanMenuBerdasarNamaAsc(TabMenu)

	fmt.Println("\n=== DAFTAR MENU SAAT INI (URUT ABJAD) ===")
	TampilkanMenu(*TabMenu)

	var namaCari string
	fmt.Print("Masukkan NAMA menu yang ingin dihapus (ketik '1' untuk BATAL): ")
	fmt.Scan(&namaCari)

	if namaCari == "1" {
		fmt.Println("Aksi dibatalkan. Kembali ke menu admin...\n")
		return
	}

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

	for i := indexHapus; i < TabMenu.BanyakMenu-1; i++ {
		TabMenu.data[i] = TabMenu.data[i+1]
	}
	TabMenu.BanyakMenu--
	fmt.Println("SUKSES: MENU BERHASIL DIHAPUS\n")
}

/*
FS: Menampilkan daftar seluruh data menu ke layar komputer (ditambah kolom Stok).
*/
func TampilkanMenu(TabMenu DaftarMenu) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("\n--- Menu Kosong ---\n")
		return
	}
	fmt.Printf("\n%-3s | %-20s | %-12s | %-10s | %-5s\n", "No", "Nama Menu", "Kategori", "Harga", "Stok")
	fmt.Println("-----------------------------------------------------------------")
	for i := 0; i < TabMenu.BanyakMenu; i++ {
		fmt.Printf("%-3d | %-20s | %-12s | Rp %-8.2f | %-5d\n", i+1, TabMenu.data[i].nama, TabMenu.data[i].kategori, TabMenu.data[i].harga, TabMenu.data[i].stok)
	}
	fmt.Println()
}

func menuSortingTampilan(TabMenu DaftarMenu, opsi int) {
	if TabMenu.BanyakMenu == 0 {
		fmt.Println("\n--- Menu Kosong ---\n")
		return
	}

	if opsi == 1 {
		fmt.Println("\n=== DAFTAR MENU TERMAHAL ===")
		for i := 1; i < TabMenu.BanyakMenu; i++ {
			key := TabMenu.data[i]
			j := i - 1
			for j >= 0 && TabMenu.data[j].harga < key.harga {
				TabMenu.data[j+1] = TabMenu.data[j]
				j--
			}
			TabMenu.data[j+1] = key
		}
	} else if opsi == 2 {
		fmt.Println("\n=== DAFTAR MENU TERMURAH ===")
		for i := 0; i < TabMenu.BanyakMenu-1; i++ {
			idxTarget := i
			for j := i + 1; j < TabMenu.BanyakMenu; j++ {
				if TabMenu.data[j].harga < TabMenu.data[idxTarget].harga {
					idxTarget = j
				}
			}
			tukar := TabMenu.data[i]
			TabMenu.data[i] = TabMenu.data[idxTarget]
			TabMenu.data[idxTarget] = tukar
		}
	}

	TampilkanMenu(TabMenu)
}

func TampilkanMenuPerKategori(TabMenu DaftarMenu, kategoriPilihan string) [NMAX]int {
	var indeksAsli [NMAX]int
	jumlahFilter := 0

	fmt.Printf("\n--- Daftar Menu Kategori: %s ---\n", kategoriPilihan)
	fmt.Printf("%-3s | %-20s | %-10s | %-5s\n", "No", "Nama Menu", "Harga", "Stok")
	fmt.Println("-------------------------------------------------")

	for i := 0; i < TabMenu.BanyakMenu; i++ {
		if TabMenu.data[i].kategori == kategoriPilihan {
			indeksAsli[jumlahFilter] = i
			jumlahFilter++
			fmt.Printf("%-3d | %-20s | Rp %-8.2f | %-5d\n", jumlahFilter, TabMenu.data[i].nama, TabMenu.data[i].harga, TabMenu.data[i].stok)
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
	fmt.Print("Pilih nomor kategori (ketik '-1' untuk BATAL): ")
	fmt.Scan(&pilihKat)

	if pilihKat == -1 {
		fmt.Println("Aksi dibatalkan.\n")
		return
	}

	if pilihKat >= 1 && pilihKat <= banyakKategori {
		TampilkanMenuPerKategori(TabMenu, daftarKategori[pilihKat-1])
	} else {
		fmt.Println("Pilihan kategori tidak valid!\n")
	}
}

/*
FS: Memproses pesanan pelanggan, mengurangi saldo dan mengurangi stok item jika berhasil.
*/
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
	fmt.Print("Pilih nomor kategori (ketik '-1' untuk BATAL): ")
	fmt.Scan(&pilihKat)

	if pilihKat == -1 {
		fmt.Println("Pemesanan dibatalkan.\n")
		return
	}

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
	fmt.Print("Pilih nomor menu yang ingin dipesan (ketik '-1' untuk BATAL): ")
	fmt.Scan(&nomor)

	if nomor == -1 {
		fmt.Println("Pemesanan dibatalkan.\n")
		return
	}

	if nomor >= 1 && nomor <= jumlahMenuFilter {
		indeksAsliMenu := petaIndeks[nomor-1]
		
		// Validasi Stok dan Saldo
		if TabMenu.data[indeksAsliMenu].stok <= 0 {
			fmt.Printf("Gagal memesan! Stok %s sedang habis.\n\n", TabMenu.data[indeksAsliMenu].nama)
		} else if *saldo >= TabMenu.data[indeksAsliMenu].harga {
			*saldo -= TabMenu.data[indeksAsliMenu].harga
			TabMenu.data[indeksAsliMenu].stok-- // Kurangi stok
			fmt.Printf("Berhasil memesan %s! Saldo Anda terpotong Rp %.2f\n\n", TabMenu.data[indeksAsliMenu].nama, TabMenu.data[indeksAsliMenu].harga)
		} else {
			fmt.Println("Gagal memesan! Saldo Anda tidak mencukupi. Silakan Top Up.\n")
		}
	} else {
		fmt.Println("Pilihan nomor menu tidak valid!\n")
	}
}