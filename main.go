package ics

type appInfo struct {
	AppId     string
	SecretId  string
	SecretKey string
}

var (
	app appInfo
)

func Init(appid string, secretId string, secretKey string) {
	app.AppId = appid
	app.SecretId = secretId
	app.SecretKey = secretKey
}
