package main

import (
	"github.com/hhthuongbtr/tulc-10xu/configuration"
	"github.com/hhthuongbtr/tulc-10xu/server"
	"log"
)



func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var conf configuration.Conf
	conf.LoadConf()
	server.RunAsHttpMode(conf)
	//msg := `{"sessionId":"U4LkCLqoFkA5Hmksqp7HRg93","alarmStatus":"0","alarmType":"event","alarmObjInfo":{"region":"hk","dimensions":{"unInstanceId":"ins-7hw5sju6","deviceName":"Unnamed33","objDetail":{"deviceLanIp":"172.19.0.125","deviceWanIp":"129.226.166.84","uniqVpcId":"vpc-lfvj69yw"}}},"alarmPolicyInfo":{"policyType":"cvm_device","policyName":"AlarmForTetHoliday","policyTypeCName":"云服务器-基础监控","conditions":{"productName":"cvm","productShowName":"云服务器","eventName":"ping_unreachable","eventShowName":"ping不可达","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":""}},"additionalMsg":[{"key":"ip","value":"something unnormal"},{"key":"alias","value":"Unnamed33"},{"key":"deviceLanIp","value":"172.19.0.125"},{"key":"deviceWanIp","value":"129.226.166.84"},{"key":"uniqVpcId","value":"vpc-lfvj69yw"}],"firstOccurTime":"2021-04-20 10:23:24","durationTime":0,"recoverTime":"0"}`
	//aa := refactorMessage(msg)
	//log.Print(aa)
}




