package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mohuani/notify/message/dingtalk"
	"io"
	"net/http"
	"strings"
)

const (
	Webhook = "https://oapi.dingtalk.com/robot/send"
)

type Client struct {
	token   string
	keyWork string // 自定义关键词
}

func NewDingTalkClient(token string, keyWork string) *Client {
	return &Client{
		token:   token,
		keyWork: keyWork,
	}
}

func (client *Client) Send(msg any) error {
	url := Webhook + "?access_token=" + client.token

	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	payload := strings.NewReader(string(message))

	httpClient := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	messageResponse := &dingtalk.MessageResponse{}
	err = json.Unmarshal(body, messageResponse)
	if err != nil {
		return err
	}
	if messageResponse.ErrCode != 0 {
		_ = fmt.Errorf("messageResponse error %v", messageResponse)
		return errors.New(messageResponse.ErrMsg)
	}

	return nil
}

func (client *Client) SendTextMessage(at dingtalk.At, text dingtalk.Text) error {
	text.Content = text.Content + client.keyWork
	textMessage := dingtalk.NewTextMessage(at, text)
	return client.Send(textMessage)
}

func (client *Client) SendMarkDownMessage(markdown dingtalk.Markdown, at dingtalk.At) error {
	markdown.Title = markdown.Title + client.keyWork
	markDownMessage := dingtalk.NewMarkDownMessage(markdown, at)
	return client.Send(markDownMessage)
}

func (client *Client) SendLinkMessage(link dingtalk.Link) error {
	link.Title = link.Title + client.keyWork
	linkMessage := dingtalk.NewLinkMessage(link)
	return client.Send(linkMessage)
}

func (client *Client) SendActionCardMessage(actionCard dingtalk.ActionCard) error {
	actionCard.Title = actionCard.Title + client.keyWork
	actionCardMessage := dingtalk.NewActionCardMessage(actionCard)
	return client.Send(actionCardMessage)
}

func (client *Client) SendFeedCardMessage(feedCard dingtalk.FeedCard) error {
	for i, link := range feedCard.Links {
		feedCard.Links[i].Title = link.Title + client.keyWork
	}

	feedCardMessage := dingtalk.NewFeedCardMessage(feedCard)
	return client.Send(feedCardMessage)
}
