package main

import "fmt"
const NMAX = 9999
type Peserta struct {
	ID          int
	nama        string
	tanggal     string
	bidangMinat string
}
type Kursus struct {
	kodeKursus string
	namaKursus string
}
type DaftarPeserta struct {
	Data [NMAX]Peserta
	N    int
}
type KatalogKursus struct {
	Data [NMAX]Kursus
	N    int
}
func tambahPendaftar(d *DaftarPeserta){
	if d.N < NMAX{
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&d.Data[d.N].ID)
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&d.Data[d.N].nama)
		fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
		fmt.Scan(&d.Data[d.N].tanggal)
		fmt.Print("Masukkan bidang minat: ")
		fmt.Scan(&d.Data[d.N].bidangMinat)
		d.N = d.N + 1
		fmt.Println("Peserta berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas pendaftaran penuh!")
	}
}

func ubahPendaftar(d *DaftarPeserta){
	var id int
	var ketemu bool = false
	var i int = 0
	fmt.Print("Masukkan ID yg Ingin diubah: ")
	fmt.Scan(&id)
	for i < d.N && !ketemu{
		if d.Data[i].ID == id{
			fmt.Println("Masukkan nama baru:")
			fmt.Scan(&d.Data[i].nama)
			fmt.Println("Masukkan tanggal baru:")
			fmt.Scan(&d.Data[i].tanggal)
			fmt.Println("Masukkan bidang minat baru:")
			fmt.Scan(&d.Data[i].bidangMinat)
			ketemu = true
			fmt.Println("Data Berhasil Diubah!")
		}
		i = i + 1
	}
	if !ketemu{
		fmt.Println("Data Tidak Ditemukan!")
	}
}

func hapusPendaftar(d *DaftarPeserta){
	var id int
	var ketemu int = -1
	fmt.Print("Masukkan ID yg Ingin dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < d.N; i++{
		if d.Data[i].ID == id{
			ketemu = i
		}
	}
	if ketemu != -1{
		for i := ketemu; i < d.N-1; i++{
			d.Data[i] = d.Data[i+1]
		}
		d.N = d.N - 1
		fmt.Println("Data Berhasil Dihapus!")
	} else {
		fmt.Println("Data Tidak Ditemukan!")
	}
}

func cariBerdasarkanMinat(d DaftarPeserta){
	var minat string
	var ketemu bool = false
	fmt.Print("Masukan Bidang Minat: ")
	fmt.Scan(&minat)
	for i := 0; i < d.N; i++{
		if d.Data[i].bidangMinat == minat{
			fmt.Println("ID: ", d.Data[i].ID, "Nama: ", d.Data[i].nama, "Tanggal: ", d.Data[i].tanggal)
			ketemu = true
		}
	}
	if !ketemu{
		fmt.Println("Data idak ditemukan!")
	}
}

func tambahKatalog(k *KatalogKursus){ 
	if k.N < NMAX{
		fmt.Print("Masukkan Kode Kursus: ")
		fmt.Scan(&k.Data[k.N].kodeKursus)
		fmt.Print("Masukkan Nama Kursus: ")
		fmt.Scan(&k.Data[k.N].namaKursus)
		k.N = k.N + 1
		fmt.Println("Katalog Kursus berhasil ditambahkan.")
	} else{
		fmt.Println("Katalog Kursus penuh!")
	}
}

func tampilkanKatalog(k KatalogKursus){
	if k.N == 0{ 
		fmt.Println("Katalog Kursus masih kosong.")
	} else {
		fmt.Println("\n=== Katalog Kursus ===")
		for i := 0; i < k.N; i++ {
			fmt.Println(k.Data[i].kodeKursus, " ", k.Data[i].namaKursus)
		}
	}
}

func cariBerdasarkanNama(){ // binary search

}

func ururtBerdasarkanId(d *DaftarPeserta){
	var pass, idx_min, i int
	var temp Peserta
	for pass = 1; pass <= d.N-1; pass++{
		idx_min = pass - 1
		for i = pass; i < d.N; i++{
			if d.Data[i].ID < d.Data[idx_min].ID {
				idx_min = i
			}
		}
		temp = d.Data[pass-1]
		d.Data[pass-1] = d.Data[idx_min]
		d.Data[idx_min] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan ID (Selection Sort).")
}

func ururtBerdasarkanNama(){ //insertion sort

}

func tampilkanStatistik(){ //?

}

func tampilkanSemua(d DaftarPeserta){
	if d.N == 0{
		fmt.Println("Belum ada data peserta.")
	} else {
		fmt.Println("\n=== Semua Data Peserta ===")
		fmt.Printf("| %-5s | %-15s | %-12s | %-12s |\n", "ID", "NAMA", "TANGGAL", "MINAT")
		for i := 0; i < d.N; i++ {
			fmt.Printf("| %-5d | %-15s | %-12s | %-12s |\n", d.Data[i].ID, d.Data[i].nama, d.Data[i].tanggal, d.Data[i].bidangMinat)
		}
	}
}

func main(){
	var daftar DaftarPeserta
	var katalog KatalogKursus
	var n int
	for n != 12{
		fmt.Println("\n=== KursusIn ===")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Ubah Data Peserta")
		fmt.Println("3. Hapus Data Peserta")
		fmt.Println("4. Cari Berdasarkan Minat (Sequential)")
		fmt.Println("5. Cari Berdasarkan Nama (Binary)")
		fmt.Println("6. Urutkan Berdasarkan ID (Selection)")
		fmt.Println("7. Urutkan Berdasarkan Nama (Insertion)")
		fmt.Println("8. Tampilkan Statistik")
		fmt.Println("9. Tampilkan Semua Peserta")
		fmt.Println("10. Tambah Katalog Kursus")
		fmt.Println("11. Tampilkan Katalog Kursus")
		fmt.Println("12. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&n)

		if n == 1{
			tambahPendaftar(&daftar)
		} else if n == 2{
			ubahPendaftar(&daftar)
		} else if n == 3{
			hapusPendaftar(&daftar)
		} else if n == 4{
			cariBerdasarkanMinat(daftar) 
		} else if n == 5{
			cariBerdasarkanNama()
		} else if n == 6{
			ururtBerdasarkanId(&daftar)
		} else if n == 7{
			ururtBerdasarkanNama()
		} else if n == 8{
			tampilkanStatistik()
		} else if n == 9{
			tampilkanSemua(daftar)
		} else if n == 10{
			tambahKatalog(&katalog)
		} else if n == 11{
			tampilkanKatalog(katalog)
		} else if n == 12{
			fmt.Println("Terima kasih telah menggunakan Sistem KursusIn!")
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}