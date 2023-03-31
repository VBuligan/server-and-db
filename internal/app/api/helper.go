package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// * Конфигурируем api instance, поле Logger
func (api *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(logLevel)
	return nil
}

// * Конфигурируем маршрутизатор, поле router api
func (api *API) configureRouterField() {
	api.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hi! This is REST API"))
	})
}
