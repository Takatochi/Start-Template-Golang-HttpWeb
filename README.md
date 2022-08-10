# StartTemplateGolangHttpWeb
 
 Example golang for start.
 
 /En/
 
It is not a 100% correct solution, the example was created for a quick start, which can be upgraded

 /Ua/
 
Не является 100% вірним рішеням, приклад був створиний для бистрого старту,який можна модернізувавати

+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*

  /En/
 include in package main
 
 /Ua/
 
 підключить до пакету main, нижче вказані модулі, фреймворки і тд

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

/En/

initialization of templates, path and route addressing.

/Ua/
ініціалізація шаблонів, шляху та адресація маршрута.

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


	r(&hadler.Index, result, router)
	hadler.Contact.Routing(result, "contact", "/contact/", router)
	}
	func r(g handler.Routined, result any, router *gin.Engine) {
		g.Routing(result, "index", "/", router)
	}
	
+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*

/En/

package hadler,the underhood part of the routing, functionality, redirection of the get route

/Ua/

пакет hadler, підкапотна частина маршрутизації, функціональності, переадресації get маршруту

  package handler

	import (
	
		"net/http"

		"github.com/gin-gonic/gin"
	)

	type Routined interface {
		Routing(post any, maineroot string, handlename string, router *gin.Engine)
	}
	type Handler struct {
		Index   index
		Contact contact
		router  *gin.Engine
	}
	type path struct {
		maineroot  string
		handlename string
	}
	type index struct {
		post any
		path
	}
	type contact struct {
		post any
		path
	}

	func InitHandler(router *gin.Engine) *Handler {
		return &Handler{
			router: router,
		}

	}

	func (s *index) Routing(post any, maineroot string, handlename string, router *gin.Engine) {
		s.post = post
		s.maineroot = maineroot
		s.handlename = handlename
		router.GET(s.handlename, s.ServeHTTP)
	}

	func (s index) ServeHTTP(ctx *gin.Context) {

		ctx.Request.ParseForm()
		get := ctx.Request.Form
		ctx.HTML(http.StatusOK, s.maineroot, gin.H{
			"Post": s.post,
			"Rget": get,
		})

	}
	func (s *index) Whole(a int, b int) bool {
	if a%b == 0 {
		return true
	} else if a%b == 1 {
		return false
	}
	return true
	}
	func (s *contact) Routing(post any, maineroot string, handlename string, router *gin.Engine) {
		s.post = post
		s.maineroot = maineroot
		s.handlename = handlename
		router.GET(s.handlename, s.ServeHTTP)
	}

	func (s contact) ServeHTTP(ctx *gin.Context) {

		ctx.Request.ParseForm()
		get := ctx.Request.Form
		ctx.HTML(http.StatusOK, s.maineroot, gin.H{
			"Post": s.post,
			"Rget": get,
		})

	}

