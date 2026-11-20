package main

import "fmt"

const NMAX int = 100

type Pengeluaran struct {
	kategori string
	jumlah   int
}

type tabData [NMAX]Pengeluaran

// ================= TAMBAH DATA =================
func tambahData(A *tabData, N *int) {
	if *N < NMAX {
		fmt.Print("Kategori: ")
		fmt.Scan(&A[*N].kategori)

		fmt.Print("Jumlah: ")
		fmt.Scan(&A[*N].jumlah)

		*N = *N + 1
	}
}

// ================= TAMPIL DATA =================
func cetakData(A tabData, N int) {
	fmt.Println("\n===== DATA PENGELUARAN =====")

	if N == 0 {
		fmt.Println("Belum ada data")
		return
	}

	for i := 0; i < N; i = i + 1 {
		fmt.Println(i, A[i].kategori, "-", A[i].jumlah)
	}
}

// ================= SEQ SEARCH =================
func SeqSearch(A tabData, N int, X string) int {
	var idx, i int

	idx = -1
	i = 0

	for idx == -1 && i < N {
		if A[i].kategori == X {
			idx = i
		}
		i = i + 1
	}

	return idx
}

// ================= BINARY SEARCH =================
func BinSearch(A tabData, N int, X string) int {
	var left, right, mid, idx int

	left = 0
	right = N - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2

		if X < A[mid].kategori {
			right = mid - 1
		} else if X > A[mid].kategori {
			left = mid + 1
		} else {
			idx = mid
		}
	}

	return idx
}

// ================= UBAH DATA =================
func ubahData(A *tabData, N int) {
	var cari string
	var idx int

	fmt.Print("Masukkan kategori yang ingin diubah: ")
	fmt.Scan(&cari)

	idx = SeqSearch(*A, N, cari)

	if idx != -1 {
		fmt.Print("Kategori baru: ")
		fmt.Scan(&A[idx].kategori)

		fmt.Print("Jumlah baru: ")
		fmt.Scan(&A[idx].jumlah)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

// ================= HAPUS DATA =================
func hapusData(A *tabData, N *int) {
	var cari string
	var idx int

	fmt.Print("Masukkan kategori yang ingin dihapus: ")
	fmt.Scan(&cari)

	idx = SeqSearch(*A, *N, cari)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	for i := idx; i < *N-1; i = i + 1 {
		A[i] = A[i+1]
	}

	*N = *N - 1
	fmt.Println("Data berhasil dihapus")
}

// ================= TOTAL =================
func total(A tabData, N int) int {
	var sum int

	for i := 0; i < N; i = i + 1 {
		sum = sum + A[i].jumlah
	}

	return sum
}

// ================= SELECTION SORT (JUMLAH) =================
func selectionSort(A *tabData, N int) {
	var pass, i, idxMin int
	var temp Pengeluaran

	for pass = 0; pass < N-1; pass = pass + 1 {
		idxMin = pass

		for i = pass + 1; i < N; i = i + 1 {
			if A[i].jumlah < A[idxMin].jumlah {
				idxMin = i
			}
		}

		temp = A[pass]
		A[pass] = A[idxMin]
		A[idxMin] = temp
	}
}

// ================= INSERTION SORT (KATEGORI) =================
func insertionSort(A *tabData, N int) {
	var pass, i int
	var temp Pengeluaran

	for pass = 1; pass < N; pass = pass + 1 {
		temp = A[pass]
		i = pass

		for i > 0 && temp.kategori < A[i-1].kategori {
			A[i] = A[i-1]
			i = i - 1
		}

		A[i] = temp
	}
}

// ================= LAPORAN =================
func laporan(A tabData, N int, budget int) {
	var totalPeng int
	var sisa int

	totalPeng = total(A, N)
	sisa = budget - totalPeng

	fmt.Println("\n==============================")
	fmt.Println("   LAPORAN BUDGET TRAVELING   ")
	fmt.Println("==============================")

	for i := 0; i < N; i = i + 1 {
		fmt.Println("-", A[i].kategori, ":", A[i].jumlah)
	}

	fmt.Println("\nBudget        :", budget)
	fmt.Println("Total keluar  :", totalPeng)
	fmt.Println("Sisa budget   :", sisa)

	if totalPeng > budget {
		fmt.Println("Status        : OVER BUDGET")
		fmt.Println("Saran         : Kurangi pengeluaran tidak penting")
	} else {
		fmt.Println("Status        : AMAN")
	}
}

// ================= MAIN =================
func main() {
	var data tabData
	var n int
	var menu int
	var budget int
	var cari string
	var idx int

	fmt.Println("================================")
	fmt.Println("  APLIKASI BUDGET TRAVELING     ")
	fmt.Println("================================")

	fmt.Print("Masukkan budget perjalanan: ")
	fmt.Scan(&budget)

	for menu != 9 {
		fmt.Println("\n========= MENU =========")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Ubah Pengeluaran")
		fmt.Println("3. Hapus Pengeluaran")
		fmt.Println("4. Tampilkan Semua Data")
		fmt.Println("5. Cari Pengeluaran (Tanpa terurut)")
		fmt.Println("6. Cari Pengeluaran (Terurut)")
		fmt.Println("7. Urutkan berdasarkan Kategori (A-Z)")
		fmt.Println("8. Urutkan berdasarkan Jumlah (Kecil-Besar)")
		fmt.Println("9. Lihat Laporan & Keluar")
		fmt.Println("========================")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		if menu == 1 {
			tambahData(&data, &n)
		} else if menu == 2 {
			ubahData(&data, n)
		} else if menu == 3 {
			hapusData(&data, &n)
		} else if menu == 4 {
			cetakData(data, n)
		} else if menu == 5 {
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&cari)

			idx = SeqSearch(data, n, cari)

			if idx != -1 {
				fmt.Println("Ditemukan:", data[idx].kategori, "-", data[idx].jumlah)
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		} else if menu == 6 {
			insertionSort(&data, n)

			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&cari)

			idx = BinSearch(data, n, cari)

			if idx != -1 {
				fmt.Println("Ditemukan:", data[idx].kategori, "-", data[idx].jumlah)
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		} else if menu == 7 {
			insertionSort(&data, n)
			fmt.Println("Data diurutkan berdasarkan kategori (A-Z)")
		} else if menu == 8 {
			selectionSort(&data, n)
			fmt.Println("Data diurutkan berdasarkan jumlah (kecil ke besar)")
		} else if menu == 9 {
			laporan(data, n, budget)
			fmt.Println("Program selesai")
		}
	}
}