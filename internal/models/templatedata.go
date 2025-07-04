package models

import "github.com/notirawap/doctor-registration/internal/forms"

// TemplateData holds data sent from handlers to template set
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
	IsAuthenticated	bool
}
