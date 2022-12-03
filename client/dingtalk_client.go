package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"notify/message"
	"strings"
)

const (
	DingTalkWebhook = "https://oapi.dingtalk.com/robot/send"
)

type DingTalkClient struct {
	token   string
	keyWork string // 自定义关键词
}

func NewDingTalkClient(token string, keyWork string) *DingTalkClient {
	return &DingTalkClient{
		token:   token,
		keyWork: keyWork,
	}
}

func (dingTalkClient *DingTalkClient) Send(msg any) error {
	url := DingTalkWebhook + "?access_token=" + dingTalkClient.token

	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	payload := strings.NewReader(string(message))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}

func (dingTalkClient *DingTalkClient) SendTextMessage(at message.At, text message.Text) error {
	text.Content = text.Content + dingTalkClient.keyWork
	textMessage := message.NewDingTalkTextMessage(at, text)
	return dingTalkClient.Send(textMessage)
}

func (dingTalkClient *DingTalkClient) SendMarkDownMessage(markdown message.Markdown, at message.At) error {
	markdown.Title = markdown.Title + dingTalkClient.keyWork
	markDownMessage := message.NewDingTalkMarkDownMessage(markdown, at)
	return dingTalkClient.Send(markDownMessage)
}

func (dingTalkClient *DingTalkClient) SendLinkMessage(link message.Link) error {
	link.Title = link.Title + dingTalkClient.keyWork
	linkMessage := message.NewDingTalkLinkMessage(link)
	return dingTalkClient.Send(linkMessage)
}

func (dingTalkClient *DingTalkClient) SendActionCardMessage(actionCard message.ActionCard) error {
	actionCard.Title = actionCard.Title + dingTalkClient.keyWork
	actionCardMessage := message.NewDingTalkActionCardMessage(actionCard)
	return dingTalkClient.Send(actionCardMessage)
}

func (dingTalkClient *DingTalkClient) SendFeedCardMessage(feedCard message.FeedCard) error {
	for i, link := range feedCard.Links {
		feedCard.Links[i].Title = link.Title + dingTalkClient.keyWork
	}

	feedCardMessage := message.NewDingTalkFeedCardMessage(feedCard)
	return dingTalkClient.Send(feedCardMessage)
}
