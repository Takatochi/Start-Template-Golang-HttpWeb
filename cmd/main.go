package main

import (
	"flag"
	"html/template"
	"log"
	"project/pkg/handler"
	"project/pkg/server"
	"project/pkg/store"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	router := gin.New()
	g := handler.InitHandler(router)

	srv := new(server.Server)

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	st := store.New(config)

	if err := st.Open(); err != nil {
		log.Fatal(err)
	}
	if err := Run(g, router); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(config, router); err != nil {
		log.Fatal(err)
	}

}
func Run(hadler *handler.Handler, router *gin.Engine) error {

	router.Static("/static", "./static/")

	router.SetFuncMap(template.FuncMap{
		"whole": handler.Whole,
	})

	router.LoadHTMLGlob("templates/*.html")

	result := []map[string]any{}

	r(&hadler.Index, result, router)
	hadler.Contact.Routing(result, "contact", "/contact/", router)

	return nil
}
func r(g handler.Routined, result any, router *gin.Engine) {
	g.Routing(result, "index", "/", router)
}
