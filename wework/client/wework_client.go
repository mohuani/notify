package client

import (
	"encoding/json"
	"fmt"
	"github.com/mohuani/notify/wework/message"
	"io"
	"net/http"
	"strings"
)

const (
	Webhook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
)

type Client struct {
	token string
}

func NewWeworkClient(token string) *Client {
	return &Client{token: token}
}

func (client *Client) Send(msg any) error {
	url := Webhook + client.token

	messageContent, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(messageContent))
	request, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	request.Header.Add("Content-Type", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(body))

	return nil
}

func (client *Client) SendTextMessage(content string, mentionedList []string, mentionedMobileList []string) error {
	message := message.NewTextMessage(content, mentionedList, mentionedMobileList)
	return client.Send(message)
}

func (client *Client) SendMarkdownMessage(content string) error {
	message := message.NewMarkdownMessage(content)
	return client.Send(message)
}

func (client *Client) SendImageMessage(base64 string, md5 string) error {
	message := message.NewImageMessage(base64, md5)
	return client.Send(message)
}

func (client *Client) SendNewsMessage(articles []message.Articles) error {
	message := message.NewNewsMessage(articles)
	return client.Send(message)
}

func (client *Client) SendFileMessage(mediaId string) error {
	message := message.NewFileMessage(mediaId)
	return client.Send(message)
}
