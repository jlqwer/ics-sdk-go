# ics-sdk-go

## 初始化

调用接口前先初始化：
```go
package main
import "github.com/jlqwer/icsSdkGo"
func main() {
	appid := ""
	secretId := ""
	secretKey := ""
    ics.Init(appid, secretId, secretKey)
}
```

## 获取黑白名单列表

```go
package main
import (
    "fmt"
    "github.com/jlqwer/icsSdkGo"
)
func main() {
	appid := ""
	secretId := ""
	secretKey := ""
	ics.Init(appid, secretId, secretKey)
    res := ics.GetBlacklistIp() //黑名单
    fmt.Println(res)
    res = ics.GetWhitelistIp() //白名单
    fmt.Println(res)
}
```

## 检查IP信息(自动添加黑名单)

```go
package main
import (
    "fmt"
    "github.com/jlqwer/icsSdkGo"
)
func main() {
	appid := ""
	secretId := ""
	secretKey := ""
	ics.Init(appid, secretId, secretKey)
    res := ics.CheckIp(ip, useragent, uri)
    fmt.Println(res)
}
```
## IP归属地查询

```go
package main
import (
    "fmt"
    "github.com/jlqwer/icsSdkGo"
)
func main() {
	appid := ""
	secretId := ""
	secretKey := ""
	ics.Init(appid, secretId, secretKey)
    res := ics.GetIpGeo(ip)
    fmt.Println(res)
}
```
