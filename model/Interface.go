package model

import "fmt"

type Maju interface {
	Cepat(word float32) float32
	Lambat(word float32) float32
}

//motor
func Hitung(ma Maju) {
	fmt.Println("jika dipercepat akan menghasilkan kecepatan : ", ma.Cepat(40))

	fmt.Println("jika diperlambat akan menghasilkan kecepatan : ", ma.Lambat(40))
}

//sepeda
func HitungKecepatan(ma Maju) {
	fmt.Println("jika dipercepat akan menghasilkan kecepatan : ", ma.Cepat(20))

	fmt.Println("jika diperlambat akan menghasilkan kecepatan : ", ma.Lambat(20))
}
