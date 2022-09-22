package configs

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ServerHost  string `requred:"true"`
	ServerPort  int    `requred:"true"`
	DatabaseUri string `requred:"true"`
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {
		if err := envconfig.Process("", &Cfg); err != nil {
			log.Panicf("Error parsing environtments vars %#v", err)
		}
	})
	return Cfg
}
