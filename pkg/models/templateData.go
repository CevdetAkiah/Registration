package models

import "github.com/CevdetAkiah/Registration/pkg/forms"

type TemplateData struct {
	StringMap map[string]string
	intMap    map[string]int
	FloatMap  map[string]string
	Data      map[string]interface{} //Data can hold data of any type
	CSRFToken string                 //This is for cross site security
	Flash     string                 //For flash message
	Warning   string                 //Warning messages
	Error     string
	Form      *forms.Form
}
