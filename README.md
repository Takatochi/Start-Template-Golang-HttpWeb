# StartTemplateGolangHttpWeb
 
 Example golang for start.
 
 /En/
 
It is not a 100% correct solution, the example was created for a quick start, which can be upgraded

 /Ua/
 
Не является 100% вірним рішеням, приклад був створиний для бистрого старту,який можна модернізувавати

+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*
 
 include in package main

	import (
	
    "encoding/json"
    
	"html/template"
    
	"net/http"
    
	"project/handler"
    
	"project/server"
    
	"github.com/gin-gonic/gin"
	)
+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*

		
/En/
Struct Server
	
/UA/	
Структура сервера


	func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()

	}

+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*


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
	result := []map[string]any{}

	err := add(&result)
	if err != nil {
		log.Fatal(err)
	}
	r(&hadler.Index, result, router)
	hadler.Contact.Routing(result, "contact", "/contact/", router)
	}
	func r(g handler.Routined, result any, router *gin.Engine) {
		g.Routing(result, "index", "/", router)
	}
