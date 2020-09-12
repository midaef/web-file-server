package models

import (
	"encoding/base64"

	"github.com/gorilla/securecookie"
)

// Token ...
type Token struct {
	Token string `json:"token"`
}

// GenerateToken ...
func GenerateToken() string {
	var token = securecookie.GenerateRandomKey(64)
	strToken := base64.StdEncoding.EncodeToString(token)
	return strToken
}
