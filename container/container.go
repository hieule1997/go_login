package container

import (
	"time"
	"github.com/BurntSushi/toml"
)

type Config struct {
	MongoURI       string
	DATABASE_NAME  string
	JWT_SECRET_KEY string
	EXPIRAION_TIME time.Duration
}
type Container struct {
	Config *Config
}
func NewContainer() *Container {
	var container = new(Container)
	return container
}
func (container *Container) Read() error {
	if _, err := toml.DecodeFile("config.toml", &container.Config); err != nil {
		return err
	}
	return nil
}