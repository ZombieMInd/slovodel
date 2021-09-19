package server

import (
	"encoding/json"
	"net/http"

	"github.com/ZombieMInd/slovodel/internal/slovodel/game"
)

func (s *server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "Pong")
	}
}

func (s *server) handleGameCreate() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err := s.services.GameService.Create(&game.Game{Name: req.Name})
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, nil)
	}
}

func (s *server) handleGameGetAll() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		result, err := s.services.GameService.GetAll()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, result)
	}
}
