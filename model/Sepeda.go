package model

import "fmt"

type Sepeda struct {
	warna string
	ban   int
	gigi  int
}

func (s Sepeda) Cepat(word float32) float32 {
	fast := (word * 2) * 50
	return fast
}

func (s Sepeda) Lambat(word float32) float32 {
	slow := (word / 2) * 50
	return slow
}

func (s Sepeda) waktuSepeda(speed float32) float32 {
	waktuTempuh := speed * 50.0

	return waktuTempuh
}

func LajuSepeda() {
	var arraySepeda = [...]Sepeda{
		Sepeda{warna: "biru", ban: 2, gigi: 5},
		// Sepeda{warna: "hitam", ban: 2, gigi: 6},
		// Sepeda{warna: "putih", ban: 2, gigi: 1},
		// Sepeda{warna: "ungu", ban: 3, gigi: 1},
		// Sepeda{warna: "merah", ban: 2, gigi: 3},
	}

	fmt.Println(arraySepeda)

	for _, value := range arraySepeda {
		fmt.Println("Warna Sepeda = " + value.warna)
		fmt.Println("jumlah ban   = ", value.ban)
		fmt.Println("jumlah gigi  = ", value.gigi)
		fmt.Println("waktu tempuh sejauh 20 km adalah ", value.waktuSepeda(20), "menit ")
		HitungKecepatan(value)
	}
}
