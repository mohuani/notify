package client

import (
	"notify/message"
	"reflect"
	"testing"
	"time"
)

const (
	TestAccessToken = "fec0514acf57fe7118dc9e1cd1491bee1c263396c18d90e7aab75cd72f1c4565" // 使用自己的 AccessToken
	TestKeyWord     = "账单"                                                               // 使用自己创建自定义钉钉机器人时，安全设置里面，填写的自定义关键词
)

func TestDingTalkClient_Send(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		msg any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Send",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				msg: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dingTalkClient := &DingTalkClient{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := dingTalkClient.Send(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDingTalkClient_SendActionCardMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		actionCard message.ActionCard
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SendActionCardMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				actionCard: message.ActionCard{
					Title:          "我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
					Text:           "![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
					BtnOrientation: "0",
					Btns: message.Btns{
						{
							Title:     "内容不错",
							ActionURL: "https://www.dingtalk.com/",
						},
						{
							Title:     "不感兴趣",
							ActionURL: "https://www.dingtalk.com/",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dingTalkClient := &DingTalkClient{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := dingTalkClient.SendActionCardMessage(tt.args.actionCard); (err != nil) != tt.wantErr {
				t.Errorf("SendActionCardMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDingTalkClient_SendFeedCardMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		feedCard message.FeedCard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SendFeedCardMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				feedCard: message.FeedCard{
					Links: message.Links{
						{
							Title:      "时代的火车向前开1",
							MessageURL: "https://www.dingtalk.com/",
							PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
						},
						{
							Title:      "时代的火车向前开2",
							MessageURL: "https://www.dingtalk.com/",
							PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dingTalkClient := &DingTalkClient{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := dingTalkClient.SendFeedCardMessage(tt.args.feedCard); (err != nil) != tt.wantErr {
				t.Errorf("SendFeedCardMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDingTalkClient_SendLinkMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		link message.Link
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SendLinkMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				link: message.Link{
					Text:       "这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林",
					Title:      "时代的火车向前开",
					PicUrl:     "",
					MessageUrl: "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dingTalkClient := &DingTalkClient{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := dingTalkClient.SendLinkMessage(tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("SendLinkMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDingTalkClient_SendMarkDownMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		markdown message.Markdown
		at       message.At
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SendMarkDownMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				markdown: message.Markdown{
					Title: "this is a markdown message title",
					Text:  "#### 杭州天气 \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n",
				},
				at: message.At{
					AtMobiles: nil,
					AtUserIds: nil,
					IsAtAll:   true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dingTalkClient := &DingTalkClient{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := dingTalkClient.SendMarkDownMessage(tt.args.markdown, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("SendMarkDownMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDingTalkClient_SendTextMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		at   message.At
		text message.Text
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SendTextMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				at: message.At{
					AtMobiles: []string{"13598055910"},
					IsAtAll:   false,
				},
				text: message.Text{
					Content: "this is a test text message" + time.Now().String(),
				},
			},
			wantErr: false,
		},
		{
			name: "SendTextMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				at: message.At{
					IsAtAll: true,
				},
				text: message.Text{
					Content: "this is a test text message" + time.Now().String(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dingTalkClient := &DingTalkClient{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := dingTalkClient.SendTextMessage(tt.args.at, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("SendTextMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDingTalkClient(t *testing.T) {
	type args struct {
		token   string
		keyWork string
	}
	tests := []struct {
		name string
		args args
		want *DingTalkClient
	}{
		{
			name: "111",
			args: args{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			want: NewDingTalkClient(TestAccessToken, TestKeyWord),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDingTalkClient(TestAccessToken, TestKeyWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDingTalkClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
