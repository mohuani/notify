package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mohuani/notify/dingtalk/message"
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

	messageContent, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	payload := strings.NewReader(string(messageContent))

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

	messageResponse := &message.MessageResponse{}
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

func (client *Client) SendTextMessage(at message.At, text message.Text) error {
	text.Content = text.Content + client.keyWork
	textMessage := message.NewTextMessage(at, text)
	return client.Send(textMessage)
}

func (client *Client) SendMarkDownMessage(markdown message.Markdown, at message.At) error {
	markdown.Title = markdown.Title + client.keyWork
	markDownMessage := message.NewMarkDownMessage(markdown, at)
	return client.Send(markDownMessage)
}

func (client *Client) SendLinkMessage(link message.Link) error {
	link.Title = link.Title + client.keyWork
	linkMessage := message.NewLinkMessage(link)
	return client.Send(linkMessage)
}

func (client *Client) SendActionCardMessage(actionCard message.ActionCard) error {
	actionCard.Title = actionCard.Title + client.keyWork
	actionCardMessage := message.NewActionCardMessage(actionCard)
	return client.Send(actionCardMessage)
}

func (client *Client) SendFeedCardMessage(feedCard message.FeedCard) error {
	for i, link := range feedCard.Links {
		feedCard.Links[i].Title = link.Title + client.keyWork
	}

	feedCardMessage := message.NewFeedCardMessage(feedCard)
	return client.Send(feedCardMessage)
}
