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

	// a := rat.Parse("1387")
	//
	// log.Println(a.LessThan(rat.Parse("1388")))
	// log.Println(rat.Rat(1).Add(rat.Rat(2)))
	// log.Println(rat.Rat("0.1").Add(rat.Rat("0.2")))
	// log.Println(rat.Rat("0.1").Add(rat.Rat("0.2")).IsEqual(rat.Rat("0.3")))
	// log.Println(rat.Rat("0.1").Add(rat.Rat("0.2")))

	// log.Println(rat.Rat("0.1").Add(rat.Rat("0.2")).IsEqual(rat.Rat("0.3")))
	log.Println(rat.Rat("0.0000000003").Mul(rat.Rat("10000")))
}
