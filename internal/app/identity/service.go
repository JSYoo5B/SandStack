package identity

import "github.com/JSYoo5B/SandStack/internal/platform/config"

type Service struct {
	config config.Config
}

func NewService(cfg config.Config) Service {
	return Service{config: cfg}
}
