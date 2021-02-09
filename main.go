package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hhthuongbtr/tulc-10xu/telegram"
	"log"
	"time"
)

type WebProxy struct {
	Host	string	`json:"host"`
	Port	int16	`json:"port"`
	UriApiFromPartner	string	`json:"uri_api_from_partner"`
	SecretKey	string	`json:"secret_key"`
	ServerlistFilePath	string	`json:"serverlist_file_path"`
	ServerListFilePathForStaging	string	`json:"server_list_file_path_for_staging"`
	ConcurrencyThread int	`json:"concurrency_thread"`
}

func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("http mode, please wait")
	webContext := WebProxy{
		Host: "0.0.0.0",
		Port: 8000,
	}
	server := initializeServer()
	setupRoute(server, &webContext)
	log.Print("begin run http server...")
	listenAdd := fmt.Sprintf("%s:%d", webContext.Host, webContext.Port)
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

func (w *WebProxy) CallBack(ctx *gin.Context) {
	//var subject string
	//var body string
	//var alarmMsg string
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	log.Println(reqBody)
	//var callbackData model.TencentAlarm
	//ctx.BindJSON(&callbackData)
	//switch callbackData.AlarmStatus {
	//case "0":
	//	switch callbackData.AlarmPolicyInfo.Conditions.EventName {
	//	case "ping_unreachable":
	//		subject = fmt.Sprintf("PROBLEM: [%s][HOST][%s]\n", callbackData.AlarmObjInfo.Dimensions.DeviceName, callbackData.AlarmPolicyInfo.Conditions.EventName)
	//	}
	//case "1":
	//}

	telegram.SendMsgToTelegram(reqBody)
	ctx.String(200, reqBody)
	return
}
func (w *WebProxy) Ping(ctx *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	log.Println(reqBody)
	ctx.String(200, "pong")
	return
}


