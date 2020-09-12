package webserver

import (
	"html/template"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type webServer struct {
	router *mux.Router
	logger *zap.Logger
	config *Config
}

func newServer(c *Config) *webServer {
	return &webServer{
		router: mux.NewRouter(),
		logger: NewLogger(c.LogLevel),
		config: c,
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
}

func (server *webServer) index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../../resources/templates/index.html")
		server.templateError(err)
		tmpl.Execute(w, nil)
	})
}

func (server *webServer) login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../../resources/templates/login.html")
		server.templateError(err)
		tmpl.Execute(w, nil)
	})
}

func (server *webServer) templateError(err error) {
	if err != nil {
		server.logger.Error("Template error",
			zap.Error(err))
	}
}
