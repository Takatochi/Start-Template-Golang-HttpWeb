// Шаблон для роботи с go та http запросами( веб сайти і тд)
package main

import (
	"html/template"
	"log"
	"project/handler"
	"project/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	g := handler.InitHandler(router)

	srv := new(server.Server)
	Run(*g, router)

	if err := srv.Run("8088", router); err != nil {
		log.Fatal(err)
	}

}
func Run(hadler handler.Handler, router *gin.Engine) {

	router.Static("/static", "./static/")

	router.SetFuncMap(template.FuncMap{
		"whole": hadler.Index.Whole,
	})

	router.LoadHTMLGlob("templates/*.html")

	r(&hadler.Index, 2, router)
	hadler.Contact.Routing(2, "contact", "/contact/", router)
}
func r(g handler.Routined, result any, router *gin.Engine) {
	g.Routing(result, "index", "/", router)
}
