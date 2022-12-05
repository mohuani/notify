package feishu

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mohuani/notify/message/feishu"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	token   string
	keyWork string // 自定义关键词
}

func NewFeiShuClient(token string, keyWork string) *Client {
	return &Client{
		token:   token,
		keyWork: keyWork,
	}
}

const (
	Webhook = "https://open.feishu.cn/open-apis/bot/v2/hook/"
)

func (client *Client) Send(msg any) error {
	url := Webhook + client.token

	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(message))
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

	if response.StatusCode != http.StatusOK {
		errMessageResponse := &feishu.ErrMessageResponse{}
		err := json.Unmarshal(body, errMessageResponse)
		if err != nil {
			return err
		}
		return errors.New(errMessageResponse.Msg)
	}

	return nil
}

func (client *Client) SendTextMessage(text string) error {
	text = text + client.keyWork
	message := feishu.NewTextMessage(text)
	return client.Send(message)
}

func (client *Client) SendPostMessage(title string, content [][]feishu.PostMessageContentPostZhCnContent) error {
	title = title + client.keyWork
	message := feishu.NewPostMessage(title, content)
	return client.Send(message)
}

func (client *Client) SendImageMessage(imageKey string) error {
	message := feishu.NewImageMessage(imageKey)
	return client.Send(message)
}

func (client *Client) SendShareChatMessage(shareChatId string) error {
	message := feishu.NewShareChatMessage(shareChatId)
	return client.Send(message)
}

func (client *Client) SendInteractiveMessage() error {
	return nil
}
