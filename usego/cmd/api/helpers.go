package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type envelope map[string]any

func (app *application) parseTemplate(w http.ResponseWriter, filepath string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	if err = t.Execute(w, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, data envelope, status int, header http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range header {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil

}
