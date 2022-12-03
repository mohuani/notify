package message

const (
	TEXT        string = "text"
	LINK        string = "link"
	MARKDOWN    string = "markdown"
	ACTION_CARD string = "actionCard"
	FEED_CARD   string = "feedCard"
)

type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

type Text struct {
	Content string `json:"content"`
}

type DingTalkTextMessage struct {
	Msgtype string `json:"msgtype"`
	At      At     `json:"at"`
	Text    Text   `json:"text"`
}

func NewDingTalkTextMessage(at At, text Text) *DingTalkTextMessage {
	return &DingTalkTextMessage{At: at, Text: text, Msgtype: TEXT}
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type DingTalkLinkMessage struct {
	Msgtype string `json:"msgtype"`
	Link    Link   `json:"link"`
}

func NewDingTalkLinkMessage(link Link) *DingTalkLinkMessage {
	return &DingTalkLinkMessage{Msgtype: LINK, Link: link}
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DingTalkMarkDownMessage struct {
	Msgype   string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}

func NewDingTalkMarkDownMessage(markdown Markdown, at At) *DingTalkMarkDownMessage {
	return &DingTalkMarkDownMessage{Msgype: MARKDOWN, Markdown: markdown, At: at}
}

type Btns []struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           Btns   `json:"btns"`
}

type DingTalkActionCardMessage struct {
	Msgtype    string     `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

func NewDingTalkActionCardMessage(actionCard ActionCard) *DingTalkActionCardMessage {
	return &DingTalkActionCardMessage{Msgtype: ACTION_CARD, ActionCard: actionCard}
}

type Links []struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

type FeedCard struct {
	Links Links `json:"links"`
}

type DingTalkFeedCardMessage struct {
	Msgtype  string   `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
}

func NewDingTalkFeedCardMessage(feedCard FeedCard) *DingTalkFeedCardMessage {
	return &DingTalkFeedCardMessage{Msgtype: FEED_CARD, FeedCard: feedCard}
}
