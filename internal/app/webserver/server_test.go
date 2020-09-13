package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonTest))
			req.Header.Set("Content-Type", "application/json")
			handler := http.HandlerFunc(s.auth().ServeHTTP)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}
}
