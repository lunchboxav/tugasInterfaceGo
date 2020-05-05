package model

type Maju interface {
	Cepat() float32
	Lambat() float32
	WaktuTempuh(float32) float32
	CetakMajuCepat()
	CetakWaktuTempuh()
}

// func CetakMaju(m Maju) {
// 	fmt.Println(m)
// 	v := m.Cepat()
// 	fmt.Printf("Dengan bergerak cepat, kecepatan %s adalah %.f km/jam\n", v)
// }

// func CetakWaktuTempuh(m Maju) {
// 	var jarak float32 = 20
// 	waktu := m.WaktuTempuh(jarak)
// 	fmt.Printf("Normalnya, dengan kecepatan %.f km/jam, jarak sejauh %.f km ditempuh selama %.f menit\n", kecepatanSepeda, jarak, waktu)
// }
