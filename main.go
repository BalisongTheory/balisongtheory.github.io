package main

import (
	"log"
	"os"

	"github.com/fr3dr/balisongtheory/generators"
)

func main() {
	err := os.MkdirAll("public", 0755)
	err = os.MkdirAll("public/tricks", 0755)
	err = os.MkdirAll("public/combos", 0755)
	if err != nil {
		log.Fatal(err)
	}

	generators.HomepageGen()
	generators.TricksGen()
	generators.TrickPageGen()
	generators.CombosGen()
	generators.ComboPageGen()
}
