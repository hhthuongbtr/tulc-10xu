package model

import "encoding/json"


type ObjDetail struct {
	DeviceLanIp				string	`json:"deviceLanIp"`
	DeviceWanIp				string	`json:"deviceWanIp"`
	UniqVpcId				string	`json:"uniqVpcId"`
	IP						string	`json:"IP"`
	PeeringConnectionName	string	`json:"PeeringConnectionName"`
	QosBandwidth			string	`json:"QosBandwidth"`
	VpcName					string	`json:"VpcName"`
	VpcId					string	`json:"VpcId"`
	VpnGatewayName			string	`json:"VpnGatewayName"`
	InternetMaxBandwidthOut	string	`json:"InternetMaxBandwidthOut"`
	Vip						string	`json:"vip"`
	Ar						string	`json:"ar"`
	Bandwidth				string	`json:"bandwidth"`
	CircuitNumber			string	`json:"circuitNumber"`
	DcType					string	`json:"dcType"`
	ConnLocalIp				string	`json:"connLocalIp"`
	ConnPeerIp				string	`json:"connPeerIp"`
}

type Dimension struct {
	UnInstanceId	string	`json:"unInstanceId"`
	ObjDetail		string	`json:"objDetail"`
}

type AlarmObjInfo struct {
	Region			string	`json:"region"`
	Dimensions		Dimension	`json:"dimensions"`
}

type Conditions struct {
	ProductName				string	`json:"productName"`
	ProductShowName			string	`json:"productShowName"`
	EventName				string	`json:"eventName"`
	EventShowName			string	`json:"eventShowName"`
	AlarmNotifyType			string	`json:"alarmNotifyType"`
	AlarmNotifyPeriod		string	`json:"alarmNotifyPeriod"`
}

type AlarmPolicyInfo struct {
	PolicyId				string	`json:"policyId`
	PolicyType				string	`json:"policyType"`
	PolicyName				string	`json:"policyName"`
	Conditions				Conditions	`json:"conditions"`
	PolicyTypeCName			string	`json:"policyTypeCName"`
}

type TencentAlarm struct {
	SessionId				string	`json:"sessionId"`
	AlarmStatus				string	`json:"alarmStatus"`
	AlarmType				string	`json:"alarmType"`
	AlarmObjInfo			AlarmObjInfo	`json:"alarmObjInfo"`
	AlarmPolicyInfo			AlarmPolicyInfo	`json:"alarmPolicyInfo"`
	FirstOccurTime			string	`json:"firstOccurTime"`
	RecoverTime				string	`json:"recoverTime"`
}

func (ta *TencentAlarm) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(ta)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (ta *TencentAlarm) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), ta)
	if err != nil {
		return err
	}
	return
}