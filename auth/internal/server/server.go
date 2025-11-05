package server

import (
	"auth/internal/config"
	"auth/internal/handler"
	"auth/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) (*Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("конфигурация сервера не можеть быть nil")
	}

	handler := handler.NewHandler(cfg)
	if handler == nil {
		return nil, fmt.Errorf("не удалось создать обработчик сервера")
	}
	fmt.Println("Обработчик сервера успешно создан")
	router := routes.SetupRouter(handler)

	return &Server{
		router: router,
		cfg:    cfg,
	}, nil
}

func (s *Server) Stop() error {
	fmt.Println("Сервер остановлен")
	return nil
}

func (s *Server) Serve() error {
	address := fmt.Sprintf("%s: %s", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Сервер готов к обработке запросов на %s...\n", address)
	return s.router.Run(address)
}
