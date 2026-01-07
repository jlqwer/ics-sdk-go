package ics

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-basic/uuid"
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

func PostJson(url string, postJson string) ([]byte, error) {
	client := &http.Client{}
	var req *http.Request
	body := io.NopCloser(strings.NewReader(postJson))
	req, _ = http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; OpenApi/4.0; +https://api.jlqwer.com/api/about)")

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	result, err := io.ReadAll(resp.Body)
	return result, err
}

type errorInfo struct {
	Code int
	Msg  string
}

type apiParam struct {
	AppId     int    `json:"appid"`
	SecretId  string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Sign      string `json:"sign"`
	Data      string `json:"data"`
}

func request(url string, postData interface{}) ([]byte, error) {
	if app.AppId == 0 || app.SecretId == "" || app.SecretKey == "" {
		info := errorInfo{
			Code: 1005,
			Msg:  "Request parameter error",
		}
		return json.Marshal(info)
	}
	dataStr, _ := json.Marshal(postData)

	var param = apiParam{
		AppId:     app.AppId,
		SecretId:  app.SecretId,
		Timestamp: time.Now().Unix(),
		Nonce:     uuid.New(),
		Data:      string(dataStr),
	}
	param.Sign = sha256Encode(fmt.Sprintf("%s%d%s%s", param.Data, param.Timestamp, param.Nonce, app.SecretKey))
	paramJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var result []byte
	for _, apiUrl := range urlSet {
		result, err := PostJson(fmt.Sprintf("https://%s%s", apiUrl, url), string(paramJson))
		if err == nil {
			return result, nil
		}
	}

	return result, err
}
