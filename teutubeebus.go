package main

import "fmt"

const NMAX = 9999

type Peserta struct {
	ID          int
	nama        string
	tanggal     string
	bidangMinat string
	status      bool
}

type Kursus struct {
	kodeKursus string
	namaKursus string
}

type DaftarPeserta struct {
	N    int
	Data [NMAX]Peserta
}

type KatalogKursus struct {
	N    int
	Data [NMAX]Kursus
}

func dataDummy(d *DaftarPeserta, k *KatalogKursus) {
	d.Data[0] = Peserta{ID: 103, nama: "Budi", tanggal: "17-10-2026", bidangMinat: "Struktur_Data_Lanjut", status: true}
	d.Data[1] = Peserta{ID: 101, nama: "Andi", tanggal: "03-10-2026", bidangMinat: "Algoritma_Pemrograman", status: true}
	d.Data[2] = Peserta{ID: 105, nama: "Citra", tanggal: "03-12-2026", bidangMinat: "Etika_AI", status: false}
	d.Data[3] = Peserta{ID: 102, nama: "Dewi", tanggal: "20-09-2026", bidangMinat: "Rekayasa_Perangkat_Lunak", status: false}
	d.Data[4] = Peserta{ID: 106, nama: "Eka", tanggal: "23-10-2026", bidangMinat: "Pemodelan_Basis_Data", status: true}
	d.Data[5] = Peserta{ID: 107, nama: "Fiona", tanggal: "21-10-2026", bidangMinat: "Pemodelan_Basis_Data", status: false}
	d.Data[6] = Peserta{ID: 108, nama: "Gio", tanggal: "21-09-2026", bidangMinat: "Pemodelan_Basis_Data", status: false}
	d.Data[7] = Peserta{ID: 109, nama: "Hani", tanggal: "14-12-2026", bidangMinat: "Etika_AI", status: true}
	d.Data[8] = Peserta{ID: 110, nama: "Cikuro", tanggal: "28-07-2026", bidangMinat: "Rekayasa_Perangkat_Lunak", status: true}
	d.Data[9] = Peserta{ID: 104, nama: "Naranta", tanggal: "08-09-2026", bidangMinat: "Algoritma_Pemrograman", status: false}
	d.N = 10

	k.Data[0] = Kursus{kodeKursus: "ALPRO", namaKursus: "Algoritma_Pemrograman"}
	k.Data[1] = Kursus{kodeKursus: "SD", namaKursus: "Struktur_Data_Lanjut"}
	k.Data[2] = Kursus{kodeKursus: "EA", namaKursus: "Etika_AI"}
	k.Data[3] = Kursus{kodeKursus: "RPL", namaKursus: "Rekayasa_Perangkat_Lunak"}
	k.Data[4] = Kursus{kodeKursus: "PBD", namaKursus: "Pemodelan_Basis_Data"}
	k.N = 5
}

func tambahPendaftar(d *DaftarPeserta) {
	var numID int

	if d.N >= NMAX { 
		fmt.Println("Kapasitas peserta sudah penuh")
		return
	}

	if d.N < NMAX { 
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&numID)

		
		for i := 0; i < d.N; i++ {
			if d.Data[i].ID == numID {
				fmt.Println("ID sudah terdaftar, silakan masukkan ID yang lain.")
				return
			}
		}

		d.Data[d.N].ID = numID
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&d.Data[d.N].nama)
		fmt.Print("Masukkan Tanggal, Bulan, Tahun: ")
		fmt.Scan(&d.Data[d.N].tanggal)
		fmt.Print("Masukkan Bidang Minat: ")
		fmt.Scan(&d.Data[d.N].bidangMinat)
		fmt.Print("Masukkan Status: ")
		fmt.Scan(&d.Data[d.N].status)
		d.N = d.N + 1 
		fmt.Println("Peserta berhasil ditambahkan.")
	}
}

func ubahPendaftar(d *DaftarPeserta) {
	var id int
	var ketemu bool = false
	var i int = 0

	fmt.Print("Masukkan ID yg Ingin diubah: ")
	fmt.Scan(&id)

	for i < d.N && !ketemu {
		if d.Data[i].ID == id {
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&d.Data[i].nama)
			fmt.Print("Masukkan tanggal baru: ")
			fmt.Scan(&d.Data[i].tanggal)
			fmt.Print("Masukkan bidang minat baru: ")
			fmt.Scan(&d.Data[i].bidangMinat)
			fmt.Print("Masukkan status baru: ")
			fmt.Scan(&d.Data[i].status)
			fmt.Println("Data Berhasil Diubah!")
			ketemu = true
		}
		i = i + 1
	}

	if !ketemu {
		fmt.Println("Data Tidak Ditemukan!")
	}
}

func hapusPendaftarBerdasarkanID(d *DaftarPeserta) {
	var ketemu int = -1
	var id int

	fmt.Print("Masukkan ID yg Ingin dihapus: ")
	fmt.Scan(&id)

	for i := 0; i < d.N; i++ {
		if d.Data[i].ID == id {
			ketemu = i
		}
	}

	if ketemu != -1 {

		for i := ketemu; i < d.N-1; i++ {
			d.Data[i] = d.Data[i+1]
		}
		d.N = d.N - 1 
		fmt.Println("Data Berhasil Dihapus!")
	} else {
		fmt.Println("Data Tidak Ditemukan!")
	}
}

func ururtBerdasarkanId(d *DaftarPeserta) {
	var pass, idx, i int
	var temp Peserta

	for pass = 1; pass <= d.N-1; pass++ {
		idx = pass - 1 

		for i = pass; i < d.N; i++ {
			if d.Data[i].ID < d.Data[idx].ID {
				idx = i
			}
		}

		temp = d.Data[pass-1]
		d.Data[pass-1] = d.Data[idx]
		d.Data[idx] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan ID.")
}


func ururtBerdasarkanNama(d *DaftarPeserta) {
	var pass int
	var temp Peserta

	for pass = 1; pass < d.N; pass++ {
		temp = d.Data[pass]
		i := pass - 1

		for i >= 0 && d.Data[i].nama > temp.nama {
			d.Data[i+1] = d.Data[i]
			i = i - 1
		}


		d.Data[i+1] = temp
	}
	fmt.Println("Data Berhasil Diurutkan Berdasarkan Nama.")
}

func cariBerdasarkanMinat(d DaftarPeserta) {
	var minat string
	var ketemu bool = false
	var k int = 0

	fmt.Print("Masukan Bidang Minat: ")
	fmt.Scan(&minat)

	for k < d.N {
		if d.Data[k].bidangMinat == minat {
			if !ketemu {
				fmt.Printf("%-31s\n", "=== Hasil Pencarian Berdasarkan Minat ===")
				fmt.Printf(" %-5s  %-15s  %-12s  %-25s  %-5s |\n",
					"ID", "NAMA", "TANGGAL", "MINAT", "STATUS")
				fmt.Println("-------------------------------------------------------------------------------------")
				ketemu = true
			}

			fmt.Printf(" %-5d  %-15s  %-12s  %-25s  %-5t |\n",
				d.Data[k].ID, d.Data[k].nama, d.Data[k].tanggal, d.Data[k].bidangMinat, d.Data[k].status)
		}
		k = k + 1
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan!")
	}
}

func cariBerdasarkanNama(d *DaftarPeserta) { 
	var namaNya string
	var kr, kn, teng int
	var ketemu int = -1
 
	fmt.Print("Masukkan Nama yang dicari: ")
	fmt.Scan(&namaNya)
 
	kr = 0
	kn = d.N - 1
 
	for kr <= kn && ketemu == -1 {
		teng = (kr + kn) / 2
 
		if namaNya < d.Data[teng].nama {
			kn = teng - 1
		} else if namaNya > d.Data[teng].nama {
			kr = teng + 1
		} else {
			ketemu = teng
		}
	}
 
	if ketemu != -1 {
		fmt.Printf("%-31s\n", "=== Hasil Pencarian Berdasarkan Nama ===")
		fmt.Printf(" %-5s  %-15s  %-12s  %-25s  %-5s |\n",
			"ID", "NAMA", "TANGGAL", "MINAT", "STATUS")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Printf(" %-5d  %-15s  %-12s  %-25s  %-5t |\n",
			d.Data[ketemu].ID, d.Data[ketemu].nama, d.Data[ketemu].tanggal, d.Data[ketemu].bidangMinat, d.Data[ketemu].status)
	} else {
		fmt.Println("Data tidak ditemukan!")
	}
}

func tambahKatalog(k *KatalogKursus) {
	var inputKode string

	if k.N < NMAX {
		fmt.Print("Masukkan Kode Kursus: ")
		fmt.Scan(&inputKode)

		for i := 0; i < k.N; i++ {
			if k.Data[i].kodeKursus == inputKode {
				fmt.Println("Kode Kursus sudah terdaftar, silakan masukkan kode yang lain.")
				return
			}
		}

		k.Data[k.N].kodeKursus = inputKode
		fmt.Print("Masukkan Nama Kursus: ")
		fmt.Scan(&k.Data[k.N].namaKursus)

		k.N = k.N + 1
		fmt.Println("Katalog Kursus berhasil ditambahkan.")
	} else {
		fmt.Println("Katalog Kursus penuh!")
	}
}

func ubahKatalog(k *KatalogKursus) {
	var kode string
	var ketemu bool = false
	var i int = 0

	fmt.Print("Masukkan Kode Kursus yg Ingin diubah: ")
	fmt.Scan(&kode)

	for i < k.N && !ketemu {
		if k.Data[i].kodeKursus == kode {
			fmt.Print("Masukkan Nama Kursus baru: ")
			fmt.Scan(&k.Data[i].namaKursus)

			fmt.Println("Data Katalog Berhasil Diubah!")
			ketemu = true
		}
		i = i + 1
	}

	if !ketemu {
		fmt.Println("Data Katalog Tidak Ditemukan!")
	}
}

func hapusKatalog(k *KatalogKursus) {
	var ketemu int = -1
	var kode string

	fmt.Print("Masukkan Kode Kursus yg Ingin dihapus: ")
	fmt.Scan(&kode)
	for i := 0; i < k.N; i++ {
		if k.Data[i].kodeKursus == kode {
			ketemu = i
		}
	}

	if ketemu != -1 {
		for i := ketemu; i < k.N-1; i++ {
			k.Data[i] = k.Data[i+1]
		}

		k.N = k.N - 1 
		fmt.Println("Data Katalog Berhasil Dihapus!")
	} else {
		fmt.Println("Data Katalog Tidak Ditemukan!")
	}
}

func tampilkanKatalog(k KatalogKursus) {
	if k.N == 0 {
		fmt.Println("Katalog Kursus masih kosong.")
	} else {
		fmt.Printf("%-31s\n", "=== Katalog Kursus ===")
		fmt.Printf(" %-10s  %-30s \n", "KODE", "NAMA KURSUS")
		fmt.Println("-----------------------------------------")

		for i := 0; i < k.N; i++ {
			fmt.Printf(" %-10s  %-30s \n", k.Data[i].kodeKursus, k.Data[i].namaKursus)
		}
	}
}

func tampilkanStatistik(d *DaftarPeserta, k *KatalogKursus) {
	var totalAktif int

	fmt.Printf("%-31s\n", "=== Statistik Peserta ===")

	for i := 0; i < k.N; i++ {
		var jumlahPeminat int = 0
		
		for j := 0; j < d.N; j++ {
			if d.Data[j].bidangMinat == k.Data[i].namaKursus {
				jumlahPeminat = jumlahPeminat + 1
			}
		}
		
		fmt.Printf("%-25s : %d\n", k.Data[i].kodeKursus, jumlahPeminat)
	}

	for i := 0; i < d.N; i++ {
		if d.Data[i].status == true {
			totalAktif = totalAktif + 1
		}
	}

	fmt.Printf("%-31s\n", "-------------------------------")
	fmt.Printf("%-25s : %d\n", "Total Peserta", d.N)
	fmt.Printf("%-25s : %d\n", "Total Peserta Aktif", totalAktif)
}

func tampilkanSemua(d DaftarPeserta) {
	if d.N == 0 {
		fmt.Println("Belum ada data peserta.")
	} else {
		fmt.Printf("%-31s\n", "=== Data Peserta ===")
		fmt.Printf("| %-5s | %-15s | %-12s | %-25s | %-10s |\n", "ID", "NAMA", "TANGGAL", "MINAT", "STATUS")
		fmt.Println("-----------------------------------------------------------------------------------")

		for i := 0; i < d.N; i++ {
			fmt.Printf("| %-5d | %-15s | %-12s | %-25s | %-10t |\n",
				d.Data[i].ID, d.Data[i].nama, d.Data[i].tanggal, d.Data[i].bidangMinat, d.Data[i].status)
		}
	}
}

func menu(daftar *DaftarPeserta, katalog *KatalogKursus) {
	var n int

	for n != 12 {
		fmt.Printf("%-31s\n", "=== KursusIn ===")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Ubah Data Peserta")
		fmt.Println("3. Hapus Data Peserta")
		fmt.Println("4. Urutkan Berdasarkan ID")
		fmt.Println("5. Urutkan Berdasarkan Nama")
		fmt.Println("6. Cari Berdasarkan Minat")
		fmt.Println("7. Cari Berdasarkan Nama")
		fmt.Println("8. Tampilkan Statistik")
		fmt.Println("9. Tampilkan Semua Peserta")
		fmt.Println("10. Kelola Katalog")
		fmt.Println("11. Tampilkan Katalog Kursus")
		fmt.Println("12. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&n)

		if n == 1 {
			tambahPendaftar(daftar)
		} else if n == 2 {
			ubahPendaftar(daftar)
		} else if n == 3 {
			hapusPendaftarBerdasarkanID(daftar)
		} else if n == 4 {
			ururtBerdasarkanId(daftar)
		} else if n == 5 {
			ururtBerdasarkanNama(daftar)
		} else if n == 6 {
			cariBerdasarkanMinat(*daftar)
		} else if n == 7 {
			cariBerdasarkanNama(daftar)
		} else if n == 8 {
			tampilkanStatistik(daftar, katalog)
		} else if n == 9 {
			tampilkanSemua(*daftar)
		} else if n == 10 {

			var subMenu int
			fmt.Print("1. Tambah | 2. Ubah | 3. Hapus (Pilih 1/2/3): ")
			fmt.Scan(&subMenu)

			if subMenu == 1 {
				tambahKatalog(katalog)
			} else if subMenu == 2 {
				ubahKatalog(katalog)
			} else if subMenu == 3 {
				hapusKatalog(katalog)
			}
		} else if n == 11 {
			tampilkanKatalog(*katalog)
		} else if n == 12 {
			fmt.Println("Terima kasih telah menggunakan Sistem KursusIn!")
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func main() {
	var daftar DaftarPeserta
	var katalog KatalogKursus

	dataDummy(&daftar, &katalog)
	menu(&daftar, &katalog)
}
