package model

import "fmt"

const kecepatanSepeda float32 = 24

type Sepeda struct {
	Warna string
	Ban   int
	Gigi  int
}

func (s *Sepeda) Cepat() float32 {
	gerakCepat := kecepatanSepeda * 2
	return gerakCepat
}

func (s *Sepeda) Lambat() float32 {
	gerakLambat := kecepatanSepeda / 2
	return gerakLambat
}

func (s *Sepeda) WaktuTempuh(jarak float32) float32 {
	waktuTempuh := jarak / kecepatanSepeda * 60
	return waktuTempuh
}

func (s *Sepeda) CetakMajuCepat() {
	pesan := s.Cepat()
	fmt.Printf("Dengan bergerak cepat, kecepatan sepeda adalah %.f km/jam\n", pesan)
}

func (s *Sepeda) CetakWaktuTempuh() {
	var jarak float32 = 20
	waktu := s.WaktuTempuh(jarak)
	fmt.Printf("Normalnya, dengan kecepatan %.f km/jam, jarak sejauh %.f km ditempuh sepeda selama %.f menit\n", kecepatanSepeda, jarak, waktu)
}
