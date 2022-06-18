package main

import (
	"html/template"
	"io"
)

type TemplateType int

const (
	ConsoleTemplate TemplateType = 0
	WebTemplate     TemplateType = 1
)

type StoryArcProvider struct {
	Story        *Story
	TemplateType TemplateType
	tpl          *template.Template
}

func (sap *StoryArcProvider) Initialise() error {
	templateName := "arc.tpl"
	if sap.TemplateType == 0 {
		templateName = "arc-console.tpl"
	}

	t := template.New(templateName)
	var err error
	sap.tpl, err = t.ParseFiles(templateName)
	return err
}

func (sap StoryArcProvider) WriteTemplateText(w io.Writer, arcName string) (*StoryArc, error) {
	arc, err := sap.Story.GetArc(arcName)
	if err != nil {
		return nil, err
	}

	er := sap.tpl.Execute(w, arc)
	if er != nil {
		return nil, er
	}

	return arc, nil
}
