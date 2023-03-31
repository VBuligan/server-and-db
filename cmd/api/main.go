package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/VBuligan/standart-web-server/internal/app/api"
	"log"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	// * На этапе запуска наше приложение будет получать config из внешнего мира (cli)
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	// * Инициализируем переменную configPath значением
	flag.Parse()
	log.Println("It works")

	// * server instance init
	config := api.NewConfig()                     // * Дефолтное значение из config.go
	_, err := toml.DecodeFile(configPath, config) // * Десеариализация .toml файла
	if err != nil {
		log.Fatal("Can not find configs file. Using default values", err)
	}
	server := api.New(config)

	// * api server start
	log.Fatal(server.Start())
}
