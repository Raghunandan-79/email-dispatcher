package main

import (
	"bytes"
	"html/template"
)

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}