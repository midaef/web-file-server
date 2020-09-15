package models

import (
	"sync"
)

// Sessions ...
type Sessions struct {
	usersMutex *sync.Mutex
	Users      map[string]*User
}

// User ...
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// NewSessions ...
func NewSessions() *Sessions {
	return &Sessions{
		usersMutex: new(sync.Mutex),
		Users:      make(map[string]*User),
	}
}

// Write ...
func (sessions *Sessions) Write(user *User) string {
	sessions.usersMutex.Lock()
	defer sessions.usersMutex.Unlock()
	token := GenerateToken()
	sessions.Users[token] = user
	return token
}

// Read ...
func (sessions *Sessions) Read(user *User) string {
	sessions.usersMutex.Lock()
	defer sessions.usersMutex.Unlock()
	for k, v := range sessions.Users {
		if v == user {
			return k
		}
	}
	return ""
}
