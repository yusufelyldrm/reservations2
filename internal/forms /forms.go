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

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if a form field has a given value
func (f *Form) Has(field string, r *http.Request) bool {
	x := f.Get(field)
	if x == "" {
		return false
	}
	return true
}
