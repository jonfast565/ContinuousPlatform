package utilities

import "github.com/aymerick/raymond"

func RunTemplate(template string, input interface{}) (*string, error) {
	tpl, err := raymond.Parse(template)
	if err != nil {
		return nil, err
	}
	result, err := tpl.Exec(input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func RunTemplateFromFile(path string, input interface{}) (*string, error) {
	return nil, nil
}
