package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/kelseyhightower/envconfig"
)

func Start(conf *Config) error {
	srv := NewServer()
	srv.configLogger(conf)
	srv.InitServices(conf)
	initRouter(srv)

	return http.ListenAndServe(conf.BindAddr, srv)
}

func InitConfig(conf *Config) error {
	err := envconfig.Process("API", conf)
	if err != nil {
		return fmt.Errorf("error while parsing env config: %s", err)
	}
	return nil
}

func initRouter(s *server) {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/ping", s.handlePing()).Methods("GET")

	s.router.HandleFunc("/game", s.handleGameCreate()).Methods("POST")
	s.router.HandleFunc("/game", s.handleGameGetAll()).Methods("GET")

}
