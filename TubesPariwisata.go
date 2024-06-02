package main

import "fmt"

const NMAX int = 100

type atributRekreasi struct {
	nama                       string
	jarak, biaya, jumFasilitas int
}
type Rekreasi [NMAX]atributRekreasi

var dataRekreasi Rekreasi
var nData int

func menuRole() {
	fmt.Println("Pilih Role:")
	fmt.Println("1. Admin")
	fmt.Println("2. User")
	fmt.Println("3. Keluar App")
}

func menuAdmin() {
	fmt.Println("Apa yang ingin anda lakukan?")
	fmt.Println("1. Menambah Data")
	fmt.Println("2. Mengubah Data")
	fmt.Println("3. Menghapus Data")
	fmt.Println("4. Exit")
}

func menuUser() {
	fmt.Println("Pilih kategori pariwisata yang sesuai kriteria anda")
	fmt.Println("1. Jarak")
	fmt.Println("2. Biaya")
	fmt.Println("3. Jumlah Fasilitasnya")
	fmt.Println("4. Ketikkan Nama Tempat Rekreasi Yang ingin Anda Cari ")
	fmt.Println("5. Exit")
}

func switchRole(choiceRole int) {
	var choiceAdmin, choiceUser int
	switch choiceRole {
	case 1:
		for choiceAdmin != 4 {
			menuAdmin()
			fmt.Scan(&choiceAdmin)
			switchAdmin(choiceAdmin)
		}
	case 2:
		for choiceUser != 5 {
			menuUser()
			fmt.Scan(&choiceUser)
			switchUser(choiceUser)
		}
	}
}

func switchAdmin(choiceAdminTask int) {
	var adminEditChoices int
	var namaDataInginDihapus string
	switch choiceAdminTask {
	case 1:
		tambahData()
	case 2:
		fmt.Println("Data apa yang ingin anda edit?")
		fmt.Println("1. Jarak")
		fmt.Println("2. Biaya")
		fmt.Println("3. Jumlah Fasilitasnya")
		fmt.Scan(&adminEditChoices)
		switchDataEdit(adminEditChoices)
	case 3:
		fmt.Println("Input nama tempat wisata yang akan dihapus dari database")
		fmt.Scan(&namaDataInginDihapus)
		idxAkun := searchDatainginDihapus(namaDataInginDihapus)
		if idxAkun == -1 {
			fmt.Println("Data tidak ditemukan")
		} else {
			hapusData(idxAkun)
			fmt.Println("Data berhasil dihapus")
			fmt.Println()
			printDatabase()
			fmt.Println()
		}
	}
}

func switchUser(choiceUser int) {
	var namaWisataInginDicari string
	switch choiceUser {
	case 1:
		fmt.Println("Berikut adalah tempat wisata berdasarkan jarak terdekat:")
		sortJarak()
		fmt.Println()
		printDatabase()
		fmt.Println()
	case 2:
		fmt.Println("Berikut adalah tempat wisata berdasarkan biaya termurah:")
		sortBiaya()
		fmt.Println()
		printDatabase()
		fmt.Println()
	case 3:
		fmt.Println("Berikut adalah tempat wisata berdasarkan jumlah fasilitas terbanyak:")
		sortjumFasilitas()
		fmt.Println()
		printDatabase()
		fmt.Println()
	case 4:
		fmt.Println("Apa tempat wisata yang ingin anda cari?")
		fmt.Scan(&namaWisataInginDicari)
		searchDataRequest(namaWisataInginDicari)
	}
}

func switchDataEdit(choicecDataEdit int) {
	var jarakBaru, biayaBaru, jumFasilitasBaru int
	var lokasiWisata string
	switch choicecDataEdit {
	case 1:
		fmt.Println("Input lokasi wisata beserta data jarak yang baru")
		fmt.Scan(&lokasiWisata, &jarakBaru)
		editDataJarak(lokasiWisata, jarakBaru)
		fmt.Println()
		printDatabase()
		fmt.Println()
	case 2:
		fmt.Println("Input lokasi wisata beserta data biaya baru")
		fmt.Scan(&lokasiWisata, &biayaBaru)
		editDataBiaya(lokasiWisata, biayaBaru)
		fmt.Println()
		printDatabase()
		fmt.Println()
	case 3:
		fmt.Println("Input lokasi wisata beserta data jumlah fasilitas baru")
		fmt.Scan(&lokasiWisata, &jumFasilitasBaru)
		editDataJumFasilitas(lokasiWisata, jumFasilitasBaru)
		fmt.Println()
		printDatabase()
		fmt.Println()
	}
}

func tambahData() {
	var name string
	var price, distance, fasilitas, n int

	fmt.Println("Berapa banyak data yang diinginkan?")
	fmt.Scan(&n)
	n += nData
	if n > NMAX {
		n = NMAX
	}

	fmt.Println("Input data-datanya")
	for i := nData; i < n; i++ {
		fmt.Scan(&name, &price, &distance, &fasilitas)
		dataRekreasi[i].nama = name
		dataRekreasi[i].biaya = price
		dataRekreasi[i].jarak = distance
		dataRekreasi[i].jumFasilitas = fasilitas
		nData++
	}
	fmt.Println()
	printDatabase()
	fmt.Println()
}

func editDataJarak(nama string, jarak int) {
	var hasil int = 0
	var i int = 0
	for i < nData && hasil == 0 {
		if dataRekreasi[i].nama == nama {
			dataRekreasi[i].jarak = jarak
			hasil = 1
		}
		i++
	}
}

func editDataBiaya(nama string, biaya int) {
	var hasil int = 0
	var i int = 0
	for i < nData && hasil == 0 {
		if dataRekreasi[i].nama == nama {
			dataRekreasi[i].biaya = biaya
			hasil = 1
		}
		i++
	}
}

func editDataJumFasilitas(nama string, jumFasilitas int) {
	var hasil int = 0
	var i int = 0
	for i < nData && hasil == 0 {
		if dataRekreasi[i].nama == nama {
			dataRekreasi[i].jumFasilitas = jumFasilitas
			hasil = 1
		}
		i++
	}
}

func searchDataRequest(nama string) {
	var hasil int = 0
	var i int = 0
	for i < nData && hasil == 0 {
		if dataRekreasi[i].nama == nama {
			fmt.Println()
			fmt.Println("Berikut adalah tempat wisata yang anda cari:")
			fmt.Println(dataRekreasi[i].nama, dataRekreasi[i].biaya, dataRekreasi[i].jarak, dataRekreasi[i].jumFasilitas)
			fmt.Println()
			hasil = 1
		}
		i++
	}
	if hasil == 0 {
		fmt.Println("Data tidak ditemukan")
	}
}

func searchDatainginDihapus(nama string) int {
	for i := 0; i < nData; i++ {
		if dataRekreasi[i].nama == nama {
			return i
		}
	}
	return -1
}

func hapusData(idx int) {
	if idx < 0 || idx >= nData {
		fmt.Println("Data tidak ditemukan")
		return
	}
	for i := idx; i < nData-1; i++ {
		dataRekreasi[i] = dataRekreasi[i+1]
	}
	nData--
}

func sortJarak() {
	var pass int
	var temp atributRekreasi
	pass = nData - 1
	for i := 0; i < pass; i++ {
		j := i
		temp = dataRekreasi[j+1]
		for j >= 0 && dataRekreasi[j].jarak > temp.jarak {
			dataRekreasi[j+1] = dataRekreasi[j]
			j--
		}
		dataRekreasi[j+1] = temp
	}
}

func sortBiaya() {
	var pass int
	var temp atributRekreasi
	pass = nData - 1
	for i := 0; i < pass; i++ {
		j := i
		temp = dataRekreasi[j+1]
		for j >= 0 && dataRekreasi[j].biaya > temp.biaya {
			dataRekreasi[j+1] = dataRekreasi[j]
			j--
		}
		dataRekreasi[j+1] = temp
	}
}

func sortjumFasilitas() {
	var pass, idx int
	var temp atributRekreasi
	pass = 1
	for pass <= nData-1 {
		idx = pass - 1
		i := pass
		for i < nData {
			if dataRekreasi[idx].jumFasilitas < dataRekreasi[i].jumFasilitas {
				idx = i
			}
			i++
		}
		temp = dataRekreasi[pass-1]
		dataRekreasi[pass-1] = dataRekreasi[idx]
		dataRekreasi[idx] = temp
		pass++
	}
}

func printDatabase() {
	fmt.Println("===============================================================")
	fmt.Printf("| %-20s | %-10s | %-10s | %-10s |\n", "Nama", "Biaya", "Jarak", "Fasilitas")
	fmt.Println("===============================================================")

	for i := 0; i < nData; i++ {
		name := dataRekreasi[i].nama
		price := dataRekreasi[i].biaya
		distance := dataRekreasi[i].jarak
		fasilitas := dataRekreasi[i].jumFasilitas
		fmt.Printf("| %-20s | %-10d | %-10d | %-10d |\n", name, price, distance, fasilitas)
	}

	fmt.Println("===============================================================")
}

func main() {
	var choiceAkun int
	for choiceAkun != 3 {
		menuRole()
		fmt.Scan(&choiceAkun)
		switchRole(choiceAkun)
	}
}
