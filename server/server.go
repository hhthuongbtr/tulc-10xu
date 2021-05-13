package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"log"
	"time"
	"fmt"
	"github.com/hhthuongbtr/tulc-10xu/configuration"
)

type WebProxy struct {
	Conf	configuration.Conf
}

func RunAsHttpMode(conf configuration.Conf) {
	log.Println("http mode, please wait")
	webContext := WebProxy{
		Conf: conf,
	}
	server := initializeServer()
	setupRoute(server, &webContext)
	log.Print("begin run http server...")
	listenAdd := fmt.Sprintf("%s:%d", webContext.Conf.Server.Host, webContext.Conf.Server.Port)
	log.Printf("serve on %s\n", listenAdd)
	err := server.Run(listenAdd)
	if err != nil {
		log.Println(err)
	}
}

func setupRoute(server *gin.Engine, webContext *WebProxy) {
	v1 := server.Group("/api/v1")
	{
		//----------------CCU-------------------
		users := v1.Group("/Callback")
		{
			users.POST("", webContext.CallBack)
			users.GET("", webContext.Ping)
		}
	}
}

func initializeServer() *gin.Engine {
	server := gin.New()
	gin.SetMode(gin.ReleaseMode)
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 30 seconds
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           30 * time.Second,
	}))
	return server
}

