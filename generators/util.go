package generators

import (
	"html/template"
	"io"
)

func executeBase(w io.Writer, file string, data any) error {
	tmpl, err := template.ParseFiles("./templates/base.html", file)
	if err != nil {
		return err
	}
	tmpl.ExecuteTemplate(w, "base", data)
	return nil
}
