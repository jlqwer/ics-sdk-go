package ics

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-basic/uuid"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"strconv"
	"strings"
	"time"
)

func sha256Encode(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func post(url string, data netUrl.Values, contentType string) ([]byte, error) {

	body := ioutil.NopCloser(strings.NewReader(data.Encode()))
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Ics/3.1; +https://api.jlqwer.com/api/about)")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return []byte(""), err
	}
	result, err := ioutil.ReadAll(resp.Body)
	return result, err
}

type errorInfo struct {
	Code int
	Msg  string
}

func request(url string, postData interface{}) ([]byte, error) {
	if app.AppId == "" || app.SecretId == "" || app.SecretKey == "" {
		info := errorInfo{
			Code: 1005,
			Msg:  "Request parameter error",
		}
		return json.Marshal(info)
	}
	dataStr, _ := json.Marshal(postData)
	param := make(netUrl.Values)
	param["appid"] = []string{app.AppId}
	param["ak"] = []string{app.SecretId}
	param["data"] = []string{string(dataStr)}
	param["timestamp"] = []string{strconv.FormatInt(time.Now().Unix(), 10)}
	param["nonce"] = []string{uuid.New()}
	param["sign"] = []string{sha256Encode(fmt.Sprintf("%s%s%s%s", param["data"][0], param["timestamp"][0], param["nonce"][0], app.SecretKey))}

	result, err := post(fmt.Sprintf("%s%s", "https://api.jlqwer.com", url), param, "")
	return result, err
}
