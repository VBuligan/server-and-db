package api

import (
	"github.com/VBuligan/server-and-db/storage"
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

// * Конфигурируем хранилище storage
func (api *API) configureStorageField() error {
	storage := storage.New(api.config.Storage)
	// * Пытаемся установить соединение, если не выходит возвращаем ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	api.storage = storage
	return nil
}
