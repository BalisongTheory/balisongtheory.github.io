package generators

import (
	"log"
	"os"
)

func SloMoToolGen() {
	err := os.Remove("public/slo-mo-tool.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	w, err := os.OpenFile("public/slo-mo-tool.html", os.O_RDWR | os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = executeBase(w, "templates/slo-mo-tool.html", map[string]any{"Title": "SLO-MO Tool"})
	if err != nil {
		log.Fatal(err)
		return
	}
}

