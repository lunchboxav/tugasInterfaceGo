package model

import "fmt"

type Motor struct {
	Ban       int
	Gear      int
	Kecepatan float32
	Jenis     string
	Maju
}

func (m Motor) Cepat(word float32) float32 {
	fast := (word * 2.0) * 50.0
	return fast
}

func (m Motor) Lambat(word float32) float32 {
	slow := (word / 2.0) * 50.0
	return slow
}

func (m Motor) waktu(speed float32) float32 {
	waktuTempuh := speed * 50.0

	return waktuTempuh
}

func LajuMotor() {
	var JnsMotor = [...]Motor{
		Motor{Ban: 2, Gear: 6, Jenis: "Beat", Kecepatan: 4.0},
		// Motor{Ban: 2, Gear: 6, Jenis: "Mio", Kecepatan: 3.0},
		// Motor{Ban: 2, Gear: 6, Jenis: "Honda", Kecepatan: 2.5},
		// Motor{Ban: 2, Gear: 6, Jenis: "Revo", Kecepatan: 2.50},
		// Motor{Ban: 2, Gear: 6, Jenis: "Yamaha", Kecepatan: 3.5},
	}

	for _, value := range JnsMotor {
		fmt.Println("Jenis Motor = ", value.Jenis)
		fmt.Println("Jumlah Ban Motor = ", value.Ban)
		fmt.Println("Jumlah Gear di sepeda = ", value.Gear)
		fmt.Println("Dalam 40 Kilometer yang waktu adalah ", value.waktu(value.Kecepatan))
		fmt.Println("dipercepat  ", value.Cepat(value.Kecepatan), "Km")
		fmt.Println("diperlambat", value.Lambat(value.Kecepatan), "Km")
		Hitung(value)
	}

}
