package generators

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type TrickVideo struct {
	Filename       string `json:"filename"`
	Format         string `json:"format"`
	ThumbnailImage string `json:"thumbnailImage"`
	Credit         string `json:"credit"`
}

type Trick struct {
	ID          string
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Videos      []TrickVideo `json:"videos"`
	Tags        []string     `json:"tags"`
	Combos      []Combo
}

type Combo struct {
	ID          string
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Videos      []TrickVideo `json:"videos"`
	Tags        []string     `json:"tags"`
	TrickIDs    []string     `json:"tricks"`
	Tricks      []Trick
}

func TricksGen() {
	entries, err := os.ReadDir("./tricks/")
	if err != nil {
		log.Fatal(err)
		return
	}

	tricks := []Trick{}
	for _, v := range entries {
		if v.Type().IsRegular() {
			data, err := os.ReadFile(fmt.Sprintf("./tricks/%s", v.Name()))
			if err != nil {
				log.Fatal(err)
				return
			}

			trick := Trick{}
			trick.ID = strings.TrimSuffix(v.Name(), filepath.Ext(v.Name()))
			err = json.Unmarshal(data, &trick)
			if err != nil {
				log.Fatalf("failed to unmarshal %s, error: %v\n", v.Name(), err)
				return
			}
			tricks = append(tricks, trick)
		}
	}

	w, err := os.OpenFile("public/tricks.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = executeBase(w, "templates/tricks.html", map[string]any{"Title": "Tricktionary", "Tricks": tricks})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func TrickPageGen() {
	entries, err := os.ReadDir("./tricks/")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, v := range entries {
		if v.Type().IsRegular() {
			id := strings.TrimSuffix(v.Name(), filepath.Ext(v.Name()))
			data, err := os.ReadFile("./tricks/" + id + ".json")
			if err != nil {
				log.Fatalln("error opening trick file:", err)
				return
			}

			trick := Trick{ID: id}
			err = json.Unmarshal(data, &trick)
			if err != nil {
				log.Fatalln("error unmarshaling trick json:", err)
				return
			}

			entries, err := os.ReadDir("./combos/")
			if err != nil {
				log.Fatal(err)
				return
			}

			for _, v := range entries {
				if v.Type().IsRegular() {
					data, err := os.ReadFile(fmt.Sprintf("./combos/%s", v.Name()))
					if err != nil {
						log.Fatal(err)
						continue
					}

					combo := Combo{}
					combo.ID = strings.TrimSuffix(v.Name(), filepath.Ext(v.Name()))
					err = json.Unmarshal(data, &combo)
					if err != nil {
						log.Fatalf("failed to unmarshal %s, error: %v\n", v.Name(), err)
						continue
					}

					if slices.Contains(combo.TrickIDs, trick.ID) {
						trick.Combos = append(trick.Combos, combo)
					}
				}
			}

			w, err := os.OpenFile("public/tricks/"+id+".html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatal(err)
				return
			}

			err = executeBase(w, "templates/trickpage.html", map[string]any{"Title": trick.Name, "Trick": trick})
			if err != nil {
				log.Fatalln("error executing base template:", err)
				return
			}
		}
	}
}

func CombosGen() {
	entries, err := os.ReadDir("./combos/")
	if err != nil {
		log.Println(err)
		return
	}

	combos := []Combo{}
	for _, v := range entries {
		if v.Type().IsRegular() {
			data, err := os.ReadFile(fmt.Sprintf("./combos/%s", v.Name()))
			if err != nil {
				log.Fatal(err)
				return
			}

			combo := Combo{}
			combo.ID = strings.TrimSuffix(v.Name(), filepath.Ext(v.Name()))
			err = json.Unmarshal(data, &combo)
			if err != nil {
				log.Fatalf("failed to unmarshal %s, error: %v\n", v.Name(), err)
				return
			}
			combos = append(combos, combo)
		}
	}

	w, err := os.OpenFile("public/combos.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = executeBase(w, "templates/combos.html", map[string]any{"Title": "Combos", "Combos": combos})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ComboPageGen() {
	entries, err := os.ReadDir("./combos/")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, v := range entries {
		if v.Type().IsRegular() {
			id := strings.TrimSuffix(v.Name(), filepath.Ext(v.Name()))
			data, err := os.ReadFile("./combos/" + id + ".json")
			if err != nil {
				log.Fatalln("error opening combo file:", err)
				return
			}

			combo := Combo{ID: id}
			err = json.Unmarshal(data, &combo)
			if err != nil {
				log.Fatalln("error unmarshaling combo json:", err)
				return
			}

			for _, v := range combo.TrickIDs {
				data, err := os.ReadFile("./tricks/" + v + ".json")
				if err != nil {
					log.Fatal(err)
					return
				}

				trick := Trick{}
				trick.ID = v
				err = json.Unmarshal(data, &trick)
				if err != nil {
					log.Fatal(err)
					return
				}

				combo.Tricks = append(combo.Tricks, trick)
			}

			w, err := os.OpenFile("public/combos/"+id+".html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatal(err)
				return
			}

			err = executeBase(w, "templates/combopage.html", map[string]any{"Title": combo.Name, "Combo": combo})
			if err != nil {
				log.Fatal("error executing base template:", err)
				return
			}
		}
	}
}
