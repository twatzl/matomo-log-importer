package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Param1 string
	Param2 bool
}

type appData struct {
	config *Config
	logger *logrus.Logger
}

type App interface {
	PrintHello()
}

func New(config *Config, logger *logrus.Logger) App {
	return &appData{
		config: config,
		logger: logger,
	}
}

func (a* appData) PrintHello() {
	a.logger.Infoln("printing hello")
	fmt.Println("Hello")
}


