package itischemes

import (
	"net/http"
	"strings"
)

// SchemesMatcher Store the schema to match with the request Path
type SchemesMatcher struct {
	schemes []string
}

// New is the constructor to ItiHost
func New(template ...string) *SchemesMatcher {
	for i := range template {
		template[i] = strings.ToLower(template[i])
	}

	return &SchemesMatcher{
		schemes: template,
	}
}

// Match returns if the request can be handled by this Route.
func (h *SchemesMatcher) Match(req *http.Request) bool {
	s := strings.ToLower(req.URL.Scheme)
	for _, v := range h.schemes {
		if v == s {
			return true
		}
	}
	return false
}
