package main

import (
	"github.com/zoowen/interfacego/model"
)

func main() {
	// hanya mencetak pesan-pesan yang dibuat di masing-masing struct
	model.Maju.CetakMajuCepat(&model.Motor{})
	model.Maju.CetakWaktuTempuh(&model.Motor{})

	model.Maju.CetakMajuCepat(&model.Sepeda{})
	model.Maju.CetakWaktuTempuh(&model.Sepeda{})
}
