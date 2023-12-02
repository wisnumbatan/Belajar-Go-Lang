package main

import (
	"fmt"
)

func metodeSimplex(c []float64, A [][]float64, b []float64) ([]float64, float64) {
	// Inisialisasi variabel
	// m mewakili banyak baris dari A
	// n mewakili banyak kolom dari A
	m := len(A)
	n := len(A[0])

	// Membuat matriks kosong
	// m = 2, n = 2
	// Baris = 3 dan kolom = 5
	matriks := make([][]float64, m+1)
	for i := range matriks {
		matriks[i] = make([]float64, m+n+1)
	}

	// Inisialisasi variabel basis
	basis := make([]int, m)
	for i := range basis {
		basis[i] = n + i
	}

	// Mengisi matriks
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matriks[i+1][j] = A[i][j]
		}
	}
	for i := 0; i < m; i++ {
		matriks[i+1][m+n] = b[i]
	}

	iterasi := 0
	for {
		fmt.Printf("### Iterasi %d ###\n", iterasi)
		// Menampilkan tabel Simpleks pada setiap iterasi
		for _, row := range matriks {
			fmt.Println(row)
		}

		// Mencari nilai kolom dengan koefisien negatif terbesar di baris pertama
		pivotKolom := 0
		for j := 0; j < n; j++ {
			if matriks[0][j] < matriks[0][pivotKolom] {
				pivotKolom = j
			}
		}

		// Jika koefisien dalam baris pertama adalah non-negatif,
		// itu berarti tidak ada kemungkinan peningkatan lebih lanjut
		// dalam nilai fungsi tujuan. Dengan kata lain, sudah
		// mencapai solusi yang optimal
		if matriks[0][pivotKolom] >= 0 {
			// Solusi optimal ditemukan
			break
		}

		// Cari indeks tiap baris
		indeks := make([]float64, m)
		for i := 0; i < m; i++ {
			indeks[i] = matriks[i+1][m+n] / matriks[i+1][pivotKolom]
		}
		pivotBaris := 0
		for i := 1; i < m; i++ {
			if indeks[i] < indeks[pivotBaris] {
				pivotBaris = i
			}
		}

		// Lakukan operasi baris untuk membuat elemen pivot menjadi 1
		elemenPivot := matriks[pivotBaris+1][pivotKolom]
		for j := 0; j <= m+n; j++ {
			matriks[pivotBaris+1][j] /= elemenPivot
		}

		// Lakukan operasi baris lainnya untuk membuat elemen lain di kolom pivot menjadi 0
		for i := 0; i <= m; i++ {
			if i != pivotBaris+1 {
				rasio := matriks[i][pivotKolom]
				for j := 0; j <= m+n; j++ {
					matriks[i][j] -= rasio * matriks[pivotBaris+1][j]
				}
			}
		}

		// Perbarui basis
		basis[pivotBaris] = pivotKolom

		iterasi++
	}

	// Ekstrak solusi dari tabel
	solusi := make([]float64, n)
	for i := 0; i < m; i++ {
		if basis[i] < n {
			solusi[basis[i]] = matriks[i+1][m+n]
		}
	}

	nilaiOptimal := -matriks[0][m+n]

	return solusi, nilaiOptimal
}

func main() {
	// Contoh masalah pemrograman linear
	c := []float64{-9, -12}
	A := [][]float64{
		{12, 1},
		{3, 1},
	}
	b := []float64{15, 18}

	solusi, nilaiOptimal := metodeSimplex(c, A, b)
	fmt.Println("Solusi optimal:")
	for i, val := range solusi {
		fmt.Printf("X%d = %.2f\n", i+1, val)
	}
	fmt.Printf("Nilai maksimum Z = %.2f\n", nilaiOptimal)
}
