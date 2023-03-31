package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// API * Base API server instance description
type API struct {
	// * UNEXPORTED FIELD
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New * API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start * Start http server/configure loggers, router, db connections and etc..
func (api *API) Start() error {
	// * Trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	// * Trying to configure Logger
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	// * Trying to configure Router
	api.configureRouterField()

	// * Trying to start Http-serv
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
