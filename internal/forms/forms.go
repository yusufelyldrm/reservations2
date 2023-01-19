package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom from struct which embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0 // If there are no errors, return true
}

// New initializes a form struct
func New(data url.Values) *Form {
	// Return a pointer to a form struct containing the data
	return &Form{
		data,
		// Initialize the errors map
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post is not empty
func (f *Form) Has(field string, r *http.Request) bool {
	// Check if the form field is in the post data
	x := f.Get(field)
	if x == "" {
		// If it is not, add an error message
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}
	return true
}
