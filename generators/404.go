package generators

import (
	"log"
	"os"
)

func PageNotFoundGen() {
	w, err := os.OpenFile("public/404.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = executeBase(w, "templates/404.html", "404", "/404", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
