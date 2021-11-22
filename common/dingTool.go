package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
)

var DingUrl string

func PostDingCardSimple(title string, body map[string]interface{}, btns []map[string]string, dingUrl string) error {
	content := fmt.Sprintf("## %s", title)
	for k, v := range body {
		content = fmt.Sprintf("%s\n- %s %v", content, k, v)
	}
	err := PostDingCard(title, content, btns, dingUrl)
	if err != nil {
		logs.Error("Post dingtalk error %s", err)
	}
	return err
}

func PostDingCard(title, body string, btns interface{}, dingUrl string) error {
	payload := map[string]interface{}{}
	payload["msgtype"] = "actionCard"
	card := map[string]interface{}{}
	card["title"] = title
	card["text"] = body
	card["hideAvatar"] = 0
	card["btns"] = btns
	payload["actionCard"] = card
	return PostJson(dingUrl, payload)
}
func PostDingmarkdown(title, body string) error {
	payload := map[string]interface{}{}
	payload["msgtype"] = "markdown"
	card := map[string]interface{}{}
	card["title"] = title
	card["text"] = body
	payload["markdown"] = card
	return PostJson(DingUrl, payload)

}

func PostDingtext(body string, dingURL string) error {
	DingUrl = dingURL
	payload := map[string]interface{}{}
	payload["msgtype"] = "text"
	card := map[string]interface{}{}
	card["content"] = body
	payload["text"] = card
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
