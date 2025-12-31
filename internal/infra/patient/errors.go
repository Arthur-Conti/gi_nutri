package patient

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

type FormulaDependencyError struct {
	Formula    string
	Dependency string
	Message    string
}

func (e *FormulaDependencyError) Error() string {
	return fmt.Sprintf("formula dependency error: %s requires %s to be calculated first. %s",
		e.Formula, e.Dependency, e.Message)
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

func NewFormulaDependencyError(formula, dependency, message string) *FormulaDependencyError {
	return &FormulaDependencyError{
		Formula:    formula,
		Dependency: dependency,
		Message:    message,
	}
}
