package langs

import "fmt"

func GenerateValidationMessage(field string, rule string) (message string) {
	switch rule {

	// required rule
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)

	// I can add more validator.v10 rules here 
	default: 
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}