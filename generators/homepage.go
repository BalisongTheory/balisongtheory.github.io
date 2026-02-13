package generators

import (
	"log"
	"os"
)

func HomepageGen() {
	err := os.Remove("public/index.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	w, err := os.OpenFile("public/index.html", os.O_RDWR | os.O_CREATE, 0644)
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
