package main

import (
	"log"

	"github.com/xtdlib/rat"
)

func main() {
	// a := big.Rat{}
	// b := big.Rat{}
	// a.SetFrac64(1, 3)
	// b.SetFrac64(2, 3)
	// log.Println(a, b)
	// a.Add(&a, &b)
	// log.Println(a, b)
	// log.Println(a.IsInt())

	a := rat.Parse("1387")
	log.Println(a)
	a.AddInt(1)
	log.Println(a)
	a.AddInt(1)
	log.Println(a)
}
