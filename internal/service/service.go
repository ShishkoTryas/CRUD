package service

import (
	"CRUDTest/internal/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type service struct {
	Router *httprouter.Router
	Config *config.Config
}

func NewServer() *service {
	return &service{
		Router: httprouter.New(),
		Config: config.NewConfig(),
	}
}

func (s *service) Run() error {

	return http.ListenAndServe(s.Config.HttpAddr, s.Router)
}
