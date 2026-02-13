package generators

import (
	"log"
	"os"
)

func HomepageGen() {
	w, err := os.OpenFile("public/index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = executeBase(w, "./templates/homepage.html", map[string]any{"Title": "Homepage"})
	if err != nil {
		log.Fatal(err)
		return
	}
}
