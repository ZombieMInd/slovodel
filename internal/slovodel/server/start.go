package server

import (
	"fmt"

	"github.com/ZombieMInd/slovodel/internal/slovodel/config"
	"github.com/kelseyhightower/envconfig"
)

func InitConfig(app *config.App) error {
	err := envconfig.Process("API", app)
	if err != nil {
		return fmt.Errorf("error while parsing env config: %s", err)
	}
	return nil
}
