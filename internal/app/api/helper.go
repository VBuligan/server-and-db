package api

import (
	"github.com/VBuligan/server-and-db/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
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
	api.router.HandleFunc(prefix+"/articles", api.GetAllArticles).Methods("GET")
	api.router.HandleFunc(prefix+"/articles/{id}", api.GetArticleById).Methods("GET")
	api.router.HandleFunc(prefix+"/articles/{id}", api.DeleteArticleById).Methods("DELETE")
	api.router.HandleFunc(prefix+"/articles", api.PostArticle).Methods("POST")
	api.router.HandleFunc(prefix+"/user/register", api.PostUserRegister).Methods("POST")
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
