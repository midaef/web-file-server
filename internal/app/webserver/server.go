package webserver

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"packages/internal/app/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

type webServer struct {
	router   *mux.Router
	logger   *zap.Logger
	config   *Config
	sessions *models.Sessions
}

func newServer(c *Config) *webServer {
	return &webServer{
		router:   mux.NewRouter(),
		logger:   NewLogger(c.LogLevel),
		config:   c,
		sessions: models.NewSessions(),
	}
}

// Run ...
func Run(config *Config) error {
	server := newServer(config)
	server.routers()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	server.logger.Info("Starting web-file-server",
		zap.String("host", server.config.Host),
		zap.String("port", server.config.Port),
		zap.String("password", server.config.Password),
		zap.String("log level", server.config.LogLevel))
	return http.ListenAndServe(config.Port, handlers.CORS(headers, methods, origins)(server.router))
}

func (server *webServer) routers() {
	server.router.PathPrefix("../resources/static/").Handler(http.StripPrefix("../resources/static/", http.FileServer(http.Dir(".././resources/static/"))))
	server.router.Handle("/", server.index())
	server.router.Handle("/login", server.login())
	server.router.Handle("/auth", server.auth()).Methods("POST")
	server.router.Handle("/token", server.token()).Methods("POST")
}

func (server *webServer) auth() http.Handler {
	var user *models.User
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes := server.requestReader(r)
		err := json.Unmarshal(bytes, &user)
		if err != nil {
			server.logger.Error("Unmarshal error", zap.Error(err))
			server.responseWriter(500, map[string]interface{}{"status": "Internal server error"}, w)
		} else {
			if user.Password == server.config.Password {
				token := server.sessions.Write(user)
				server.responseWriter(200, &models.Token{Token: token}, w)
			} else {
				server.responseWriter(400, map[string]interface{}{"status": "Bad request"}, w)
			}
		}
	})
}

func (server *webServer) index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../../resources/templates/index.html")
		server.templateError(err)
		tmpl.Execute(w, nil)
	})
}

func (server *webServer) token() http.Handler {
	var token *models.Token
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes := server.requestReader(r)
		err := json.Unmarshal(bytes, &token)
		if err != nil {
			server.logger.Error("Unmarshal error", zap.Error(err))
			server.responseWriter(500, map[string]interface{}{"status": "Internal server error"}, w)
		} else {
			if _, ok := server.sessions.Users[token.Token]; ok {
				server.responseWriter(200, map[string]interface{}{"status": "OK"}, w)
			} else {
				server.responseWriter(404, map[string]interface{}{"status": "Bad request"}, w)
			}
		}
	})
}

func (server *webServer) login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../../resources/templates/login.html")
		server.templateError(err)
		tmpl.Execute(w, nil)
	})
}

func (server *webServer) requestReader(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.logger.Error("Request reader error", zap.Error(err))
	}
	return body
}

func (server *webServer) responseWriter(statusCode int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json, err := json.Marshal(data)
	if err != nil {
		server.logger.Error("Json marshal error", zap.Error(err))
	}
	_, err = w.Write(json)
	if err != nil {
		server.logger.Error("Response writer error", zap.Error(err))
	}
}

func (server *webServer) templateError(err error) {
	if err != nil {
		server.logger.Error("Template error", zap.Error(err))
	}
}
