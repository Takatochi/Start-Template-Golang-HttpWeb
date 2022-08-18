package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"project/app"
	srv "project/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	//server ...

	flag.Parse()

	config := srv.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	app.Run(config)

}
