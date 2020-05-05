package model

// Maju adalah nama interface yang mencakup perhitungan kecepatan dan percetakan data
type Maju interface {
	Cepat() float32
	Lambat() float32
	WaktuTempuh(float32) float32
	CetakMajuCepat()
	CetakWaktuTempuh()
}
