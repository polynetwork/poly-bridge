package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

var (
	DingUrl = "https://oapi.dingtalk.com/robot/send?access_token=63395d10b3104b3b3817db7d6d673b4cd7452b7a375e333dd07b85f17c6c9ca6"
)

func PostDingCardSimple(title string, body map[string]interface{}, btns []map[string]string) error {
	content := fmt.Sprintf("## %s", title)
	for k, v := range body {
		content = fmt.Sprintf("%s\n- %s %v", content, k, v)
	}
	err := PostDingCard(title, content, btns)
	if err != nil {
		logs.Error("Post dingtalk error %s", err)
	}
	return err
}

func PostDingCard(title, body string, btns interface{}) error {
	payload := map[string]interface{}{}
	payload["msgtype"] = "actionCard"
	card := map[string]interface{}{}
	card["title"] = title
	card["text"] = body
	card["hideAvatar"] = 0
	card["btns"] = btns
	payload["actionCard"] = card
	return PostJson(DingUrl, payload)
}

func PostJson(url string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	logs.Info("PostJson response Body:", string(respBody))
	return nil
}
