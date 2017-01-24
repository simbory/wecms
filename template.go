package wecms

import "time"

type TemplateField struct {
	Name            string
	DisplayTitle    string
	FieldType       string
	Mandatory       bool
	ValidationRegex string
	DefaultValue    string
}

type TemplateSection struct {
	Name   string
	Fields []TemplateField
}

func (s *TemplateSection) GetField(name string) *TemplateField {
	if len(s.Fields) == 0 {
		return nil
	}
	for _, f := range s.Fields {
		if f.Name == name {
			return &f
		}
	}
	return nil
}

type Template struct {
	Id         ID `bson:"__id"`
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
	CreatedBy  string
	UpdatedBy  string
	Sections   []TemplateSection
}

func (t *Template) GetSection(name string) *TemplateSection {
	if len(t.Sections) == 0 {
		return nil
	}
	for _, section := range t.Sections {
		if section.Name == name {
			return &section
		}
	}
	return nil
}

func (t *Template) GetField(name string) *TemplateField {
	if len(t.Sections) == 0 {
		return nil
	}
	for _, section := range t.Sections {
		if len(section.Fields) == 0 {
			continue
		}
		if f := section.GetField(name); f != nil {
			return f
		}
	}
	return nil
}