package forms

// Error is a type that holds a map of errors
type errors map[string][]string

// Add adds an error message to the map of errors
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get gets the first error message for a given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
