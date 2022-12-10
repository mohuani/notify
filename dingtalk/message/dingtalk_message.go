package message

const (
	TEXT        string = "text"
	LINK        string = "link"
	MARKDOWN    string = "markdown"
	ACTION_CARD string = "actionCard"
	FEED_CARD   string = "feedCard"
)

// TextMessage
type TextMessage struct {
	Msgtype string `json:"msgtype"`
	At      At     `json:"at"`
	Text    Text   `json:"text"`
}
type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}
type Text struct {
	Content string `json:"content"`
}

func NewTextMessage(content string, at At) *TextMessage {
	return &TextMessage{
		Msgtype: TEXT,
		Text: Text{
			Content: content,
		},
		At: at,
	}
}

// LinkMessage
type LinkMessage struct {
	Msgtype string `json:"msgtype"`
	Link    Link   `json:"link"`
}
type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

func NewLinkMessage(link Link) *LinkMessage {
	return &LinkMessage{Msgtype: LINK, Link: link}
}

// MarkDownMessage
type MarkDownMessage struct {
	Msgype   string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewMarkDownMessage(markdown Markdown, at At) *MarkDownMessage {
	return &MarkDownMessage{Msgype: MARKDOWN, Markdown: markdown, At: at}
}

// ActionCardMessage
type ActionCardMessage struct {
	Msgtype    string     `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
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

func NewActionCardMessage(actionCard ActionCard) *ActionCardMessage {
	return &ActionCardMessage{Msgtype: ACTION_CARD, ActionCard: actionCard}
}

// FeedCardMessage
type FeedCardMessage struct {
	Msgtype  string   `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
}
type FeedCardLink struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

type FeedCard struct {
	Links []FeedCardLink `json:"links"`
}

func NewFeedCardMessage(links []FeedCardLink) *FeedCardMessage {
	return &FeedCardMessage{
		Msgtype: FEED_CARD,
		FeedCard: FeedCard{
			Links: links,
		},
	}
}
