// ipify-api/error_handlers
//
// This package holds our API error handlers which we use to service REST API
// errors.

package api

import (
	"net/http"
)

// NotFound renders a not found response for invalid API endpoints.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

// MethodNotAllowed renders a method not allowed response for invalid request
// types.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
}
