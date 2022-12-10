package message

const (
	TEXT     = "text"
	MARKDOWN = "markdown"
	IMAGE    = "image"
	NEWS     = "news"
	FILE     = "file"
)

// TextMessage
type TextMessage struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}
type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

func NewTextMessage(content string, mentionedList []string, mentionedMobileList []string) *TextMessage {
	return &TextMessage{
		Msgtype: TEXT,
		Text: Text{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	}
}

// MarkdownMessage
type MarkdownMessage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}
type Markdown struct {
	Content string `json:"content"`
}

func NewMarkdownMessage(content string) *MarkdownMessage {
	return &MarkdownMessage{
		Msgtype: MARKDOWN,
		Markdown: Markdown{
			Content: content,
		},
	}
}

// ImageMessage
type ImageMessage struct {
	Msgtype string `json:"msgtype"`
	Image   Image  `json:"image"`
}
type Image struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

func NewImageMessage(base64 string, md5 string) *ImageMessage {
	return &ImageMessage{
		Msgtype: IMAGE,
		Image: Image{
			Base64: base64,
			Md5:    md5,
		},
	}
}

// NewsMessage
type NewsMessage struct {
	Msgtype string `json:"msgtype"`
	News    News   `json:"news"`
}
type News struct {
	Articles []Articles `json:"articles"`
}
type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}

func NewNewsMessage(articles []Articles) *NewsMessage {
	return &NewsMessage{
		Msgtype: NEWS,
		News: News{
			Articles: articles,
		},
	}
}

// FileMessage
type FileMessage struct {
	Msgtype string `json:"msgtype"`
	File    File   `json:"file"`
}
type File struct {
	MediaID string `json:"media_id"`
}

func NewFileMessage(mediaId string) *FileMessage {
	return &FileMessage{
		Msgtype: FILE,
		File: File{
			MediaID: mediaId,
		},
	}
}
