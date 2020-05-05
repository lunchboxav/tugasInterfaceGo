package model

import "fmt"

const kecepatanMotor float32 = 48

type Motor struct {
	Ban       int
	Gear      int
	Kecepatan float32
	Jenis     string
}

func (m *Motor) Cepat() float32 {
	gerakCepat := kecepatanMotor * 2
	return gerakCepat
}

func (m *Motor) Lambat() float32 {
	gerakLambat := kecepatanMotor / 2
	return gerakLambat
}

func (m *Motor) WaktuTempuh(jarak float32) float32 {
	waktuTempuh := jarak / kecepatanMotor * 60
	return waktuTempuh
}

func (m *Motor) CetakMajuCepat() {
	pesan := m.Cepat()
	fmt.Printf("Dengan bergerak cepat, kecepatan motor adalah %.f km/jam\n", pesan)
}

func (m *Motor) CetakWaktuTempuh() {
	var jarak float32 = 20
	waktu := m.WaktuTempuh(jarak)
	fmt.Printf("Normalnya, dengan kecepatan %.f km/jam, jarak sejauh %.f km ditempuh motor selama %.f menit\n", kecepatanMotor, jarak, waktu)
}
