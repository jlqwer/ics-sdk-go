package ics

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

type IpList struct {
	Code  int
	Msg   string
	Data  []string
	Count int
}

func getIpList(ipType string) IpList {
	var ipList = IpList{}
	result, err := request("/Api/Ipc/getIpcList", map[string]string{"type": ipType})
	if err != nil {
		ipList.Code = -1
		ipList.Msg = err.Error()
		return ipList
	}
	err = json.Unmarshal(result, &ipList)
	if err != nil {
		ipList.Code = -1
		ipList.Msg = err.Error()
	}

	return ipList
}

// GetBlacklistIp 获取IP黑名单列表
func GetBlacklistIp() IpList {
	return getIpList("ipb")
}

// GetWhitelistIp 获取IP白名单列表
func GetWhitelistIp() IpList {
	return getIpList("ipw")
}

type checkResultData struct {
	Code      int
	Appkey    string
	Ip        string
	Useragent string
	Visit     int
	Ipb       int
	Ipw       int
	IpbTime   int
	IpwTime   int
	Isbot     int
	VisitTtl  int
}

type CheckResult struct {
	Code  int
	Msg   string
	Data  checkResultData
	Count int
}

// CheckIp 检查IP黑白名单&真假蜘蛛
func CheckIp(ip string, useragent string, uri string) CheckResult {
	useragent = base64.StdEncoding.EncodeToString([]byte(useragent))
	uri = hex.EncodeToString([]byte(uri))
	var checkResult CheckResult
	result, err := request("/Api/Ipc/check", map[string]string{"ip": ip, "useragent": useragent, "uri": uri})
	if err != nil {
		checkResult.Code = -1
		checkResult.Msg = err.Error()
		return checkResult
	}
	err = json.Unmarshal(result, &checkResult)
	if err != nil {
		checkResult.Code = -1
		checkResult.Msg = err.Error()
	}
	return checkResult
}
