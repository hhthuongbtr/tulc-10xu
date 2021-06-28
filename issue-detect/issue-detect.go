package issue_detect

import (
	"fmt"
	"github.com/hhthuongbtr/tulc-10xu/model"
	"log"
)

func RefactorMessage(alarmMsgRecieve string) (msgToSend string) {
	//var alarmMsgRecieve string
	//alarmMsgRecieve = `{"sessionId":"ECkYDW5CbFIIPabdCq4y7w3t","alarmStatus":"1","alarmType":"metric","alarmObjInfo":{"region":"hk","namespace":"qce/cvm","dimensions":{"deviceName":"Unnamed1","objId":"dc2df8c0-e7a8-4f03-ac06-94fb93a427d5","objName":"172.19.0.34#3552563","unInstanceId":"ins-4ap7abk6"}},"alarmPolicyInfo":{"policyId":"policy-9dz6qbwy","policyType":"cvm_device","policyName":"AlarmForTetHoliday","policyTypeCName":"","policyTypeEname":"Cloud Virtual Machine","conditions":{"metricName":"disk_usage","metricShowName":"DiskUtilization ","calcType":">","calcValue":"95","currentValue":"96.739","unit":"%","period":"60","periodNum":"5","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":5}},"firstOccurTime":"2021-05-05 04:37:00","durationTime":86400,"recoverTime":"0"}`
	//alarmMsgRecieve = `{"sessionId":"GM9LWZPzIQEx62OzrpixQ056","alarmStatus":"0","alarmType":"metric","alarmObjInfo":{"region":"hk","namespace":"qce/cvm","dimensions":{"deviceName":"Unnamed1","objId":"dc2df8c0-e7a8-4f03-ac06-94fb93a427d5","objName":"172.19.0.34#3552563","unInstanceId":"ins-4ap7abk6"}},"alarmPolicyInfo":{"policyId":"policy-9dz6qbwy","policyType":"cvm_device","policyName":"AlarmForTetHoliday","policyTypeCName":"","policyTypeEname":"Cloud Virtual Machine","conditions":{"metricName":"disk_usage","metricShowName":"DiskUtilization ","calcType":">","calcValue":"95","currentValue":"26.14","unit":"%","period":"60","periodNum":"5","alarmNotifyType":"singleAlarm","alarmNotifyPeriod":5}},"firstOccurTime":"2021-05-05 04:37:00","durationTime":101340,"recoverTime":"2021-05-06 08:46:00"}`
	var revieveData model.TencentAlarm
	revieveData.LoadFromJsonString(alarmMsgRecieve)
	//log.Printf("data : %#v", revieveData)
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
	switch revieveData.AlarmType {
	case "metric":
		switch revieveData.AlarmPolicyInfo.Conditions.MetricName {
		case "disk_usage":
			log.Print("Disk issue")
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Disk used %s %s%% Value=%s%%][Incident level 5]
Trigger status: PROBLEM
Trigger severity: Disaster`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjName,
				revieveData.AlarmObjInfo.Dimensions.DeviceName,
				revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
				revieveData.AlarmPolicyInfo.Conditions.CalcType,
				revieveData.AlarmPolicyInfo.Conditions.CalcValue,
				revieveData.AlarmPolicyInfo.Conditions.CurrentValue)
			return msg
		case "cpu_load_1":
			log.Print("CPULoadAvg1m")
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Server Load %s %s%% in 1 mins][Value=%s]
Trigger status: PROBLEM
Trigger severity: Disaster`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjName,
				revieveData.AlarmObjInfo.Dimensions.DeviceName,
				revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
				revieveData.AlarmPolicyInfo.Conditions.CalcType,
				revieveData.AlarmPolicyInfo.Conditions.CalcValue,
				revieveData.AlarmPolicyInfo.Conditions.CurrentValue)
			return msg
		case "cpu_load_5":
			log.Print("CPULoadAvg5m")
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Server Load %s %s%% in 5 mins][Value=%s]
Trigger status: PROBLEM
Trigger severity: Disaster`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjName,
				revieveData.AlarmObjInfo.Dimensions.DeviceName,
				revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
				revieveData.AlarmPolicyInfo.Conditions.CalcType,
				revieveData.AlarmPolicyInfo.Conditions.CalcValue,
				revieveData.AlarmPolicyInfo.Conditions.CurrentValue)
			return msg
		case "cpu_load_15":
			log.Print("CPULoadAvg15m")
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Server Load %s %s%% in 15 mins][Value=%s]
Trigger status: PROBLEM
Trigger severity: Disaster`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjName,
				revieveData.AlarmObjInfo.Dimensions.DeviceName,
				revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
				revieveData.AlarmPolicyInfo.Conditions.CalcType,
				revieveData.AlarmPolicyInfo.Conditions.CalcValue,
				revieveData.AlarmPolicyInfo.Conditions.CurrentValue)
			return msg
		case "mem_usage":
			log.Print("MemoryUtilization")
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Memory used %s %s%%][Value=%s]
Trigger status: PROBLEM
Trigger severity: Disaster`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjName,
				revieveData.AlarmObjInfo.Dimensions.DeviceName,
				revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
				revieveData.AlarmPolicyInfo.Conditions.CalcType,
				revieveData.AlarmPolicyInfo.Conditions.CalcValue,
				revieveData.AlarmPolicyInfo.Conditions.CurrentValue)
			return msg
		// CPU Usage rate
		case "cpu_use_rate":
			log.Print("cpu_use_rate")
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][CPU Util %s %s%%][Value=%s] in %s seconds
Trigger status: PROBLEM
Trigger severity: Disaster`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjName,
				revieveData.AlarmObjInfo.Dimensions.UInstanceId,
				revieveData.AlarmPolicyInfo.Conditions.MetricShowName,
				revieveData.AlarmPolicyInfo.Conditions.CalcType,
				revieveData.AlarmPolicyInfo.Conditions.CalcValue,
				revieveData.AlarmPolicyInfo.Conditions.CurrentValue
				revieveData.AlarmPolicyInfo.Conditions.Period)
			return msg
		default:
			return alarmMsgRecieve
		}
	case "event":
		switch revieveData.AlarmPolicyInfo.Conditions.EventName {
		case "ping_unreachable":
			log.Print("uncreachable")
			msg = fmt.Sprintf(`Alarm id: %s
Time issue: %s
Server: %s
Trigger: [%s][Server is unreachable]
Trigger status: PROBLEM
Trigger severity: Disaster`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjDetail.DeviceLanIp,
				revieveData.AlarmObjInfo.Dimensions.DeviceName)
			return msg
		default:
			return alarmMsgRecieve
		}
	default:
		return alarmMsgRecieve
	}
}

func resolveIssue(alarmMsgRecieve string) (msg string) {
	var revieveData model.TencentAlarm
	revieveData.LoadFromJsonString(alarmMsgRecieve)

	switch revieveData.AlarmType {
	case "metric":
		switch revieveData.AlarmPolicyInfo.Conditions.MetricName {
		case "disk_usage":
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Disk used %s %s%% Value=%s%%][Resolved]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
				revieveData.SessionId,
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
		case "cpu_load_1":
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Server Load %s %s%% in 1 mins][Value=%s][Resolved]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
				revieveData.SessionId,
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
		case "cpu_load_5":
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Server Load %s %s%% in 5 mins][Value=%s][Resolved]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
				revieveData.SessionId,
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
		case "cpu_load_15":
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Server Load %s %s%% in 15 mins][Value=%s][Resolved]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
				revieveData.SessionId,
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
		case "mem_usage":
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][Memory %s %s%%][Value=%s]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
				revieveData.SessionId,
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
		// CPU use rate
		case "cpu_use_rate":
			msg = fmt.Sprintf(`Alarm ID: %s
Time issue: %s
Server: %s
Trigger: [%s][%s][CPU Util %s %s%%][Value=%s]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjName,
				revieveData.AlarmObjInfo.Dimensions.UInstanceId,
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
	case "event":
		switch revieveData.AlarmPolicyInfo.Conditions.EventName {
		case "ping_unreachable":
			log.Print("uncreachable")
			msg = fmt.Sprintf(`Alarm id: %s
Time issue: %s
Server: %s
Trigger: [%s][Server is unreachable][Resolved]
Trigger status: RESOLVED
Trigger severity: Info
Time resolved: %s
Duration time in second: %d`,
				revieveData.SessionId,
				revieveData.FirstOccurTime,
				revieveData.AlarmObjInfo.Dimensions.ObjDetail.DeviceLanIp,
				revieveData.AlarmObjInfo.Dimensions.DeviceName,
				revieveData.RecoverTime,
				revieveData.DurationTime)
			return msg
		default:
			return alarmMsgRecieve
		}
	default:
		return alarmMsgRecieve
	}
}


