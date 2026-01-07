package ics

type appInfo struct {
	AppId     int
	SecretId  string
	SecretKey string
}

var (
	app appInfo
)

func Init(appid int, secretId string, secretKey string) {
	app.AppId = appid
	app.SecretId = secretId
	app.SecretKey = secretKey
}
