package models

import (
	"sync"
)

// UsersStorage ...
type UsersStorage struct {
	usersMutex *sync.Mutex
	Users      map[string]*User
}

// User ...
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// NewUsersStorage ...
func NewUsersStorage() *UsersStorage {
	return &UsersStorage{
		usersMutex: new(sync.Mutex),
		Users:      make(map[string]*User),
	}
}

// Write ...
func (usersStorage *UsersStorage) Write(user *User) string {
	usersStorage.usersMutex.Lock()
	defer usersStorage.usersMutex.Unlock()
	token := GenerateToken()
	usersStorage.Users[token] = user
	return token
}
