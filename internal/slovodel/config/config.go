package config

type App struct {
	Name string `envconfig:"NAME" default:"some"`
}
