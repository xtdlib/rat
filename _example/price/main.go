package main

import (
	"log"

	"github.com/xtdlib/rat"
)

func main() {
	log.Println("start")
	PrintPrice(rat.Rat("3.4"))
}

func PrintPrice(price rat.Price) {
	log.Println(price.String())
}
