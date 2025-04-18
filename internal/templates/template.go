package templates

import (
	"os"
)

type TemplateLoader struct{}

func NewTemplateLoader() *TemplateLoader {
	return &TemplateLoader{}
}

func (l *TemplateLoader) LoadTemplate(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	
	return string(content), nil
}