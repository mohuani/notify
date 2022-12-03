package dingtalk

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

type TextMessage struct {
	Msgtype string `json:"msgtype"`
	At      At     `json:"at"`
	Text    Text   `json:"text"`
}

func NewTextMessage(at At, text Text) *TextMessage {
	return &TextMessage{At: at, Text: text, Msgtype: TEXT}
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type LinkMessage struct {
	Msgtype string `json:"msgtype"`
	Link    Link   `json:"link"`
}

func NewLinkMessage(link Link) *LinkMessage {
	return &LinkMessage{Msgtype: LINK, Link: link}
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type MarkDownMessage struct {
	Msgype   string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}

func NewMarkDownMessage(markdown Markdown, at At) *MarkDownMessage {
	return &MarkDownMessage{Msgype: MARKDOWN, Markdown: markdown, At: at}
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

type ActionCardMessage struct {
	Msgtype    string     `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

func NewActionCardMessage(actionCard ActionCard) *ActionCardMessage {
	return &ActionCardMessage{Msgtype: ACTION_CARD, ActionCard: actionCard}
}

type Links []struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

type FeedCard struct {
	Links Links `json:"links"`
}

type FeedCardMessage struct {
	Msgtype  string   `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
}

func NewFeedCardMessage(feedCard FeedCard) *FeedCardMessage {
	return &FeedCardMessage{Msgtype: FEED_CARD, FeedCard: feedCard}
}
