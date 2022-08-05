package ics

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-basic/uuid"
	"io"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"strconv"
	"strings"
	"time"
)

var urlSet = []string{
	"api.jlqwer.com",
	"node0.api.jlqwer.com",
	"node1.api.jlqwer.com",
}

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
	if err != nil {
		return []byte(""), err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
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

	var result []byte
	var err error
	for _, apiUrl := range urlSet {
		result, err := post(fmt.Sprintf("https://%s%s", apiUrl, url), param, "")
		if err == nil {
			return result, nil
		}
	}

	return result, err
}
