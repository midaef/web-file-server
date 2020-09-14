package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"packages/internal/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleUserAuth(t *testing.T) {
	config := NewConfig()
	s := newServer(config)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    "test",
				"password": config.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"login":    "test",
				"password": "test",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid",
			payload:      map[string]string{},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid",
			payload: map[string]int{
				"login":    1,
				"password": 2,
			},
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			jsonTest, _ := json.Marshal(test.payload)
			req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(jsonTest))
			req.Header.Set("Content-Type", "application/json")
			handler := http.HandlerFunc(s.auth().ServeHTTP)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}
}

func TestHandleToken(t *testing.T) {
	config := NewConfig()
	s := newServer(config)
	user := &models.User{
		Login:    "midaef",
		Password: config.Password,
	}
	token := s.sessions.Write(user)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"token": token,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"token": "",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid",
			payload:      map[string]string{},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid",
			payload: map[string]int{
				"token": 1,
			},
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			jsonTest, _ := json.Marshal(test.payload)
			req, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(jsonTest))
			req.Header.Set("Content-Type", "application/json")
			handler := http.HandlerFunc(s.token().ServeHTTP)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}
}
