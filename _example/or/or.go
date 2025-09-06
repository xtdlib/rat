package main

import (
	"cmp"
	"log"

	"github.com/xtdlib/rat"
)

func main() {
	v := rat.Rat(3)
	log.Println(cmp.Or(v, rat.Rat(4)))
	log.Println(cmp.Or(rat.Rat(0), v))
	log.Println(rat.Or(rat.Rat(0), v))
}
