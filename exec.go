package goapitestgen

import "log"

func Exec(openAPISpecPath, filename string) {

	openAPISpec, err := ReadJson(openAPISpecPath)
	if err != nil {
		log.Fatal(err)
	}

	generatedCode, err := GenerateTest(openAPISpec, "api_test_template.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := WriteToFile(generatedCode, filename); err != nil {
		log.Fatal(err)
	}
}
