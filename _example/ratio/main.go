package main

import (
	"log"

	"github.com/xtdlib/rat"
)

func main() {
	log.Println(rat.Premium("1", "0.9991").DecimalString())
	log.Println(rat.Premium("1423", "1422").DecimalString())
	log.Println(rat.Premium("1420", "1419").DecimalString())
}
