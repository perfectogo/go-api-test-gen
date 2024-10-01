package goapitestgen

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// cutWord extracts meaningful parts from the path and formats them appropriately
func cutWord(path string) string {
	parts := strings.Split(path, "/")
	switch len(parts) {
	case 0, 1:
		return path
	case 2:
		return titleCase(parts[1])
	default:
		if parts[2] == "{id}" {
			return titleCase(parts[1]) + "ById"
		}
		return titleCase(parts[1]) + titleCase(parts[2])
	}
}

// titleCase returns the title-cased version of a string, trimming any surrounding whitespace
func titleCase(s string) string {
	return cases.Title(language.English).String(strings.TrimSpace(s))
}

// upper converts a string to uppercase
func upper(s string) string {
	return strings.ToUpper(s)
}

// extractDtoName extracts the DTO name from a JSON reference string (e.g., "#/definitions/DtoName")
func extractDtoName(ref string) string {
	const prefix = "#/definitions/"
	if !strings.HasPrefix(ref, prefix) {
		return ""
	}
	return strings.TrimSpace(ref[len(prefix):])
}

// swaggerTypeToGoType maps Swagger types to Go types
func swaggerTypeToGoType(swaggerType string) string {
	switch swaggerType {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		return "[]interface{}"
	case "object":
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

// ToPascalCase converts a string to PascalCase, handling acronyms and underscores
func ToPascalCase(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == '_'
	})

	acronyms := map[string]string{
		"id":  "ID",
		"url": "URL",
		"api": "API",
	}

	for i, word := range words {
		word = strings.ToLower(word)
		if acronym, found := acronyms[word]; found {
			words[i] = acronym
		} else {
			words[i] = capitalizeFirstLetter(word)
		}
	}

	return strings.Join(words, "")
}

// makeFnName creates a function name from an API path, converting path variables (e.g., {id}) to "ById"
func makeFnName(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			part = strings.Trim(part, "{}")
			parts[i] = "By" + titleCase(part)
		} else {
			parts[i] = titleCase(part)
		}
	}
	return strings.Join(parts, "")
}

// capitalizeFirstLetter capitalizes the first letter of a string
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// getAPIPath returns the API path by removing the leading part before the first "/"
func getAPIPath(basePath string) string {
	if slashIndex := strings.Index(basePath, "/"); slashIndex != -1 {
		return basePath[slashIndex:]
	}
	return basePath
}
