package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Wolfxxxz/Dictionary/internal/app/apiserver"
)

var (
	configPath string //= "configs/api.toml"
	//configPath string = "configs/.env"
)

func init() {

	flag.StringVar(&configPath, "pathtoml", "configs/api.toml", "path to config file in .toml format")
	// - path to config file in .env format - В каком файде хранится инфо (дескриптор)
	//flag.StringVar(&configPath, "pathenv", "configs/.env", "path to config file in .env format")
	// Параметры path - Должны отличатся
}

func StartServer() {
	//Запускаем функцию init и flag.StringVar()
	flag.Parse()
	log.Println("It works")
	//Server instance initialization
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can not find configs file. Using default value", err)
	}
	//Теперь нужно попробовать прочитать из .toml/.env , так как там может быть новая информация
	server := apiserver.New(config)

	//И сохранить в config
	/*if configPath == "configs/.env" {
		//Если на FLAG пришол запрос на puthenv
		err := godotenv.Load("configs/.env")
		if err != nil {
			log.Fatal("Could not find .env file:", err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")
		//fmt.Println(config)
	} else if configPath == "configs/api.toml" {
		//Если на FLAG пришол запрос на puthtoml
		//Достаём конфиг из томл var configPath string = "configs/api.toml"
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println("Can not find configs file. Using default value", err)
		}
	}*/

	//api server Start
	//log.Fatal(server.Start())

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
