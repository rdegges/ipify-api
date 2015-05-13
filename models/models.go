// ipify-api/models
//
// This package contains all models used in the ipify service.

package models

// IPAddress is a struct we use to represent JSON API responses.
type IPAddress struct {
	IP string `json:"ip"`
}
