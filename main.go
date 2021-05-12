package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hhthuongbtr/tulc-10xu/configuration"
	"github.com/hhthuongbtr/tulc-10xu/model"
	"github.com/hhthuongbtr/tulc-10xu/telegram"
	"log"
	"time"
)

type WebProxy struct {
	Conf	configuration.Conf
}

func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var conf configuration.Conf
	conf.LoadConf()
	runAsHttpMode(conf)
}

func runAsHttpMode(conf configuration.Conf) {
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

	msgToSend := refactorMessage(reqBody)
	telegram.SendMsgToTelegram(&w.Conf, msgToSend)
	ctx.String(200, msgToSend)
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



func refactorMessage(alarmMsgRecieve string) (msgToSend string) {
	//var alarmMsgRecieve string
	//alarmMsgRecieve = `{"sessionId":"ECkYDW5CbFIIPabdCq4y7w3t","alarmStatus":"1","alarmType":"metric","alarmObjInfo":{"region":"hk","namespace":"qce/cvm","dimensions":{"deviceName":"Unnamed1","objId":"dc2df8c0-e7a8-4f03-ac06-94fb93a427d5","objName":"172.19.0.34#3552563","unInstanceId":"ins-4ap7abk6"}},"alarmPolicyInfo":{"policyId":"policy-9dz6qbwy","policyType":"cvm_device","policyName":"AlarmForTetHoliday","policyTypeCName":"","policyTypeEname":"Cloud Virtual Machine","conditions":{"metricName":"disk_usage","metricShowName":"DiskUtilization ","calcType":">","calcValue":"95","currentValue":"96.739","unit":"%","period":"60","periodNum":"5","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":5}},"firstOccurTime":"2021-05-05 04:37:00","durationTime":86400,"recoverTime":"0"}`
	//alarmMsgRecieve = `{"sessionId":"GM9LWZPzIQEx62OzrpixQ056","alarmStatus":"0","alarmType":"metric","alarmObjInfo":{"region":"hk","namespace":"qce/cvm","dimensions":{"deviceName":"Unnamed1","objId":"dc2df8c0-e7a8-4f03-ac06-94fb93a427d5","objName":"172.19.0.34#3552563","unInstanceId":"ins-4ap7abk6"}},"alarmPolicyInfo":{"policyId":"policy-9dz6qbwy","policyType":"cvm_device","policyName":"AlarmForTetHoliday","policyTypeCName":"","policyTypeEname":"Cloud Virtual Machine","conditions":{"metricName":"disk_usage","metricShowName":"DiskUtilization ","calcType":">","calcValue":"95","currentValue":"26.14","unit":"%","period":"60","periodNum":"5","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":5}},"firstOccurTime":"2021-05-05 04:37:00","durationTime":101340,"recoverTime":"2021-05-06 08:46:00"}`
	var revieveData model.TencentAlarm
	revieveData.LoadFromJsonString(alarmMsgRecieve)
	log.Printf("data : %#v", revieveData)
	switch revieveData.AlarmStatus {
	case "1":
		log.Print("detect issue")
		msgToSend = detectIssue(alarmMsgRecieve)
	case "0":
		log.Print("resolve issue")
		msgToSend = resolveIssue(alarmMsgRecieve)
	default:
		msgToSend = alarmMsgRecieve
	}
	return msgToSend
}

func detectIssue(alarmMsgRecieve string) (msg string) {
	var revieveData model.TencentAlarm
	revieveData.LoadFromJsonString(alarmMsgRecieve)
	switch revieveData.AlarmPolicyInfo.Conditions.MetricName {
	case "disk_usage":
		msg = fmt.Sprintf(`Time issue: %s
Server: %s
Trigger: [%s][%s][Disk used %s %s%% Value=%s%%][Incident level 5]
Trigger status: PROBLEM
Trigger severity: Disaster`,
			revieveData.FirstOccurTime,
			revieveData.AlarmObjInfo.Dimensions.ObjName,
			revieveData.AlarmObjInfo.Dimensions.DeviceName,
			revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
			revieveData.AlarmPolicyInfo.Conditions.CalcType,
			revieveData.AlarmPolicyInfo.Conditions.CalcValue,
			revieveData.AlarmPolicyInfo.Conditions.CurrentValue)
		return msg
	default:
		return alarmMsgRecieve
	}
}

func resolveIssue(alarmMsgRecieve string) (msg string) {
	var revieveData model.TencentAlarm
	revieveData.LoadFromJsonString(alarmMsgRecieve)
	switch revieveData.AlarmPolicyInfo.Conditions.MetricName {
	case "disk_usage":
		msg = fmt.Sprintf(`Time issue: %s
Server: %s
Trigger: [%s][%s][Disk used %s %s%% Value=%s%%][Resolved]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
			revieveData.FirstOccurTime,
			revieveData.AlarmObjInfo.Dimensions.ObjName,
			revieveData.AlarmObjInfo.Dimensions.DeviceName,
			revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
			revieveData.AlarmPolicyInfo.Conditions.CalcType,
			revieveData.AlarmPolicyInfo.Conditions.CalcValue,
			revieveData.AlarmPolicyInfo.Conditions.CurrentValue,
			revieveData.RecoverTime,
			revieveData.DurationTime)
		return msg
	default:
		return alarmMsgRecieve
	}
}

