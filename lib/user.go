package model

import (
	"encoding/json"
	"io"
)

// User blahblah
type User struct {
	// TODO NEVER store password plaintext
	Login       string `json:"login"`
	Password    string `json:"password"`
	FirstName   string `json:"fistName,omitempty"`
	MidName     string `json:"midName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	DisplayName string `json:"displayName"`
}

// JSONUser blahblah
func JSONUser(s string) (User, error) {
	var u User
	err := json.Unmarshal([]byte(s), &u)
	return u, err
}

// ReadUser blahblah
func ReadUser(r io.Reader) (Cert, error) {
	var u Cert
	err := json.NewDecoder(r).Decode(&u)
	return u, err
}
