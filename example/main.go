package main

import (
	"log"

	"github.com/xtdlib/rat"
)

// Geometric progression
// 1 + 1*3 + 1*3^2 + 1*3^3 + 1*3^4
func main() {
	a := rat.Int(1)
	r := rat.Int(3)
	sum := rat.Zero()

	for n := 0; n < 5; n++ {
		log.Println(a.Mul(r.PowInt(n)))
		sum = sum.Add(a.Mul(r.PowInt(n)))
	}
	log.Println(sum.String())
}
