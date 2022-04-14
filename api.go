package ics

import "fmt"

func GetWhitelistIp() {
	result := request("/Api/Ipc/getIpcList", map[string]string{"type": "ipw"})
	fmt.Println(result)
}
