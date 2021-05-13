package server

import (
	"github.com/gin-gonic/gin"
	issue_detect "github.com/hhthuongbtr/tulc-10xu/issue-detect"
	"github.com/hhthuongbtr/tulc-10xu/telegram"
	"log"
)

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

	msgToSend := issue_detect.RefactorMessage(reqBody)
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
