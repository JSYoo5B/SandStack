package config

import (
	"net"
	"os"
)

type Config struct {
	Port        string
	PublicURL   string
	Region      string
	Username    string
	Password    string
	UserID      string
	ProjectID   string
	ProjectName string
	RoleName    string
}

func Load() Config {
	return Config{
		Port:        env("SANDSTACK_PORT", "9696"),
		PublicURL:   os.Getenv("SANDSTACK_PUBLIC_URL"),
		Region:      env("SANDSTACK_REGION", "RegionOne"),
		Username:    env("SANDSTACK_DEFAULT_USERNAME", "admin"),
		Password:    env("SANDSTACK_DEFAULT_PASSWORD", "password"),
		UserID:      env("SANDSTACK_DEFAULT_USER_ID", "admin"),
		ProjectID:   env("SANDSTACK_DEFAULT_PROJECT_ID", "demo"),
		ProjectName: env("SANDSTACK_DEFAULT_PROJECT_NAME", "demo"),
		RoleName:    env("SANDSTACK_DEFAULT_ROLE_NAME", "admin"),
	}
}

func (c Config) ListenAddress() string {
	return net.JoinHostPort("", c.Port)
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
