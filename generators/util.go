package generators

import (
	"html/template"
	"io"
)

func executeBase(w io.Writer, file string, title string, url string, data map[string]any) error {
	if data == nil {
		data = map[string]any{}
	}
	data["Title"] = title
	data["URL"] = ROOTURL + url
	data["RootURL"] = ROOTURL
	tmpl, err := template.ParseFiles("./templates/base.html", file)
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		return err
	}
	return nil
}
