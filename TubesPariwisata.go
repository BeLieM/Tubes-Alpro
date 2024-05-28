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
}

func switchRole(a int) {
	var choiceAdmin, choiceUser int
	switch a {
	case 1:
		for choiceAdmin != 4 {
			menuAdmin()
			fmt.Scan(&choiceAdmin)
			switchAdmin(choiceAdmin)
		}
	case 2:
		for choiceUser != 4 {
			menuUser()
			fmt.Scan(&choiceUser)
			switchUser(choiceUser)
		}
	}
}

func switchAdmin(a int) {
	var b, jarak, idxAkun int
	var nama string
	switch a {
	case 1:
		fmt.Println("Berapa banyak data yang diinginkan?")
		fmt.Scan(&nData)
		fmt.Println("Input data-datanya")
		tambahData()
	case 2:
		fmt.Println("Data apa yang ingin anda edit?")
		fmt.Println("1. Jarak")
		fmt.Println("2. Biaya")
		fmt.Println("3. Jumlah Fasilitasnya")
		fmt.Scan(&b)
		switchDataEdit(b)
	case 3:
		fmt.Println("Input data mana yang akan dihapus dari database")
		fmt.Scan(&nama, &jarak)
		idxAkun = searchData(nama, jarak)
		hapusData(idxAkun)
	}
}

func switchUser(a int) {
	var nama string
	switch a {
	case 1:
		fmt.Println("Berikut adalah tempat wisata berdasarkan jarak terdekat:")
		sortJarak()
		printDatabase()
	case 2:
		fmt.Println("Berikut adalah tempat wisata berdasarkan biaya termurah:")
		sortBiaya()
		printDatabase()
	case 3:
		fmt.Println("Berikut adalah tempat wisata berdasarkan jumlah fasilitas terbanyak:")
		sortjumFasilitas()
		printDatabase()
	case 4:
		fmt.Println("Apa tempat wisata yang ingin anda cari?")
		fmt.Scan(&nama)
		searchDataRequest(nama)
	}
}

func switchDataEdit(a int) {
	var jarakBaru, biayaBaru, jumFasilitasBaru int
	var lokasi string
	switch a {
	case 1:
		fmt.Println("Input data jarak yang baru")
		fmt.Scan(&lokasi, &jarakBaru)
		editDataJarak(lokasi, jarakBaru)
	case 2:
		fmt.Println("Input data biaya baru")
		fmt.Scan(&lokasi, &biayaBaru)
		editDataBiaya(lokasi, biayaBaru)
	case 3:
		fmt.Println("Input data jumlah fasilitas baru")
		fmt.Scan(&lokasi, &jumFasilitasBaru)
		editDataJumFasilitas(lokasi, jumFasilitasBaru)
	}
}

func tambahData() {
	var name string
	var price, distance, fasilitas int
	if nData > NMAX {
		nData = NMAX
	}
	for i := 0; i < *&nData; i++ {
		fmt.Scan(&name, &price, &distance, &fasilitas)
		dataRekreasi[i].nama = name
		dataRekreasi[i].biaya = price
		dataRekreasi[i].jarak = distance
		dataRekreasi[i].jumFasilitas = fasilitas
	}
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
			fmt.Println(dataRekreasi[i].nama, dataRekreasi[i].biaya, dataRekreasi[i].jarak, dataRekreasi[i].jumFasilitas)
			hasil = 1
		}
		i++
	}
	if hasil == 0 {
		fmt.Println("Data tidak ditemukan")
	}
}

func searchData(nama string, jarak int) int {
	sortJarak()
	var le, ri, mid int
	le = 0
	ri = nData - 1
	for jarak != mid && le <= ri {
		mid = (le + ri) / 2
		if dataRekreasi[mid].jarak > jarak {
			ri = mid - 1
		} else if dataRekreasi[mid].jarak < jarak {
			le = mid + 1
		} else {
			mid = jarak
		}
	}
	return mid
}

func hapusData(idx int) {
	var i int
	for i = idx; i < nData-1; i++ {
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
	var name string
	var price, distance, fasilitas int
	for i := 0; i < nData; i++ {
		name = dataRekreasi[i].nama
		price = dataRekreasi[i].biaya
		distance = dataRekreasi[i].jarak
		fasilitas = dataRekreasi[i].jumFasilitas
		fmt.Println(name, price, distance, fasilitas)
	}
}

func main() {
	var choice int
	for choice != 3 {
		menuRole()
		fmt.Scan(&choice)
		switchRole(choice)
	}
}
