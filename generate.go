package goapitestgen

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
)

// ReadJson reads the OpenAPI spec JSON file and unmarshals it into a map.
func ReadJson(openAPISpecPath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(openAPISpecPath)
	if err != nil {
		return nil, err
	}

	var swaggerJSON map[string]interface{}
	if err := json.Unmarshal(data, &swaggerJSON); err != nil {
		return nil, err
	}

	return swaggerJSON, nil
}

// GenerateTest generates test code based on the provided OpenAPI spec and Go template.
func GenerateTest(openAPISpec map[string]interface{}, templatePath string) (string, error) {
	// Read template file content
	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	// Create and configure template
	tmpl, err := template.New("apiTest").Funcs(template.FuncMap{
		"title":               titleCase,
		"cut":                 cutWord,
		"upper":               upper,
		"extractDtoName":      extractDtoName,
		"swaggerTypeToGoType": swaggerTypeToGoType,
		"toPascalCase":        ToPascalCase,
		"makeFnName":          makeFnName,
		"getAPIPath":          getAPIPath,
	}).Parse(string(tmplContent))
	if err != nil {
		return "", err
	}

	// Buffer to hold generated code
	var buff bytes.Buffer
	if err = tmpl.Execute(&buff, openAPISpec); err != nil {
		return "", err
	}

	// Format generated code with goimports
	formattedCode, err := imports.Process("", buff.Bytes(), &imports.Options{
		Fragment:   true,
		FormatOnly: false,
	})
	if err != nil {
		return "", err
	}

	return string(formattedCode), nil
}

// WriteToFile writes the generated content to a file.
func WriteToFile(content, filename string) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}

	file, err := os.Create(filename+"_test.go")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
