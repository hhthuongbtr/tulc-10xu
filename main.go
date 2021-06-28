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

	//ping_unreachable
	//msg := `{"sessionId":"U4LkCLqoFkA5Hmksqp7HRg93","alarmStatus":"0","alarmType":"event","alarmObjInfo":{"region":"hk","dimensions":{"unInstanceId":"ins-7hw5sju6","deviceName":"Unnamed33","objDetail":{"deviceLanIp":"172.19.0.125","deviceWanIp":"129.226.166.84","uniqVpcId":"vpc-lfvj69yw"}}},"alarmPolicyInfo":{"policyType":"cvm_device","policyName":"AlarmForTetHoliday","policyTypeCName":"云服务器-基础监控","conditions":{"productName":"cvm","productShowName":"云服务器","eventName":"ping_unreachable","eventShowName":"ping不可达","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":""}},"additionalMsg":[{"key":"ip","value":"something unnormal"},{"key":"alias","value":"Unnamed33"},{"key":"deviceLanIp","value":"172.19.0.125"},{"key":"deviceWanIp","value":"129.226.166.84"},{"key":"uniqVpcId","value":"vpc-lfvj69yw"}],"firstOccurTime":"2021-04-20 10:23:24","durationTime":0,"recoverTime":"0"}`
	//Cpu over load
	//msg := `{"sessionId":"E7bIIH9qLkE7Oi6e8PYfG3aA","alarmStatus":"0","alarmType":"metric","alarmObjInfo":{"region":"sg","namespace":"qce/cvm","dimensions":{"deviceName":"gt-omg3sea-misc-01","objId":"a1192217-868b-44b6-937a-ec3dfadcba01","objName":"172.22.0.14#4112475","unInstanceId":"ins-age4ak12"}},"alarmPolicyInfo":{"policyId":"policy-8f5m0jo2","policyType":"cvm_device","policyName":"Default","policyTypeCName":"","policyTypeEname":"Cloud Virtual Machine","conditions":{"metricName":"cpu_load_1","metricShowName":"CPULoadAvg1m ","calcType":">","calcValue":"8","currentValue":"8.91","unit":"","period":"60","periodNum":"5","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":5}},"firstOccurTime":"2021-05-12 20:56:00","durationTime":0,"recoverTime":"0"}`
	//Out of memory
	//msg := `{"sessionId":"9LCvdYx3sawt1H1emvaaMpx0","alarmStatus":"1","alarmType":"metric","alarmObjInfo":{"region":"sg","namespace":"qce/cvm","dimensions":{"deviceName":"gt-omg3sea-gametest-misc-02","objId":"f49fc905-e8bf-4f00-a62e-a4933958e7e7","objName":"172.22.0.4#4112475","unInstanceId":"ins-jis5nk7e"}},"alarmPolicyInfo":{"policyId":"policy-8f5m0jo2","policyType":"cvm_device","policyName":"Default","policyTypeCName":"","policyTypeEname":"Cloud Virtual Machine","conditions":{"metricName":"mem_usage","metricShowName":"MemoryUtilization ","calcType":">","calcValue":"90","currentValue":"90.683","unit":"%","period":"60","periodNum":"5","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":5}},"firstOccurTime":"2021-05-11 20:31:00","durationTime":0,"recoverTime":"0"}`
	//CPU Util over threshold
	// msg := `{"sessionId":"LBzsq6xMHFfelYg7vRcH9Kkr","alarmStatus":"1","alarmType":"metric","alarmObjInfo":{"region":"sg","namespace":"qce/cdb","dimensions":{"objId":"09a82419-aca7-11eb-a85e-0c42a163eabd","objName":"cdb-59ds6dfv(instance name:gt-omg3sea-game-db-03,IP:172.22.0.47:3306)","uInstanceId":"cdb-59ds6dfv"}},"alarmPolicyInfo":{"policyId":"policy-r9mwiupe","policyType":"cdb_detail","policyName":"Default","policyTypeCName":"","policyTypeEname":"CDB-MySQL-MASTER","conditions":{"metricName":"cpu_use_rate","metricShowName":"cpu_use_rate","calcType":">","calcValue":"80","currentValue":"100","unit":"","period":"300","periodNum":"1","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":1}},"firstOccurTime":"2021-06-27 23:05:00","durationTime":0,"recoverTime":"0"}`
	// aa := issue_detect.RefactorMessage(msg)
	// log.Print(aa)
}




