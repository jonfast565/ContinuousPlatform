package templating

import (
	"github.com/aymerick/raymond"
	"io/ioutil"
)

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
	template, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	templateString := string(template)
	result, err := RunTemplate(templateString, input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
