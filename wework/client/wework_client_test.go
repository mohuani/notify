package client

import (
	"github.com/mohuani/notify/wework/message"
	"reflect"
	"testing"
	"time"
)

const TestAccessToken = "3697adcf-fed9-4bf2-b906-2f47cbc3470e"

func TestClient_Send(t *testing.T) {
	type fields struct {
		token string
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
		{
			name: "WeworkClient_Send",
			fields: fields{
				token: TestAccessToken,
			},
			args: args{
				msg: message.NewTextMessage("Send 这是一条企业微信的测试通知", nil, nil),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token: tt.fields.token,
			}
			if err := client.Send(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendFileMessage(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		mediaId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WeworkClient_SendFileMessage",
			fields: fields{
				token: TestAccessToken,
			},
			args: args{
				mediaId: "3a8asd892asd8asd",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token: tt.fields.token,
			}
			if err := client.SendFileMessage(tt.args.mediaId); (err != nil) != tt.wantErr {
				t.Errorf("SendFileMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendImageMessage(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		base64 string
		md5    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WeworkClient_SendImageMessage",
			fields: fields{
				token: TestAccessToken,
			},
			args: args{
				md5: "80b9a31e0c8ba9d25bd4a32a94fcf3e9",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token: tt.fields.token,
			}
			if err := client.SendImageMessage(tt.args.base64, tt.args.md5); (err != nil) != tt.wantErr {
				t.Errorf("SendImageMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendMarkdownMessage(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WeworkClient_SendMarkdownMessage",
			fields: fields{
				token: TestAccessToken,
			},
			args: args{
				content: "实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n>类型:<font color=\"comment\">用户反馈</font>\n>普通用户反馈:<font color=\"comment\">117例</font>\n >VIP用户反馈:<font color=\"comment\">15例</font>",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token: tt.fields.token,
			}
			if err := client.SendMarkdownMessage(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("SendMarkdownMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendNewsMessage(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		articles []message.Articles
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WeworkClient_SendNewsMessage",
			fields: fields{
				token: TestAccessToken,
			},
			args: args{
				articles: []message.Articles{
					{
						Title:       "中秋节礼品领取",
						Description: "今年中秋节公司有豪礼相送",
						URL:         "www.qq.com",
						Picurl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token: tt.fields.token,
			}
			if err := client.SendNewsMessage(tt.args.articles); (err != nil) != tt.wantErr {
				t.Errorf("SendNewsMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendTextMessage(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		content             string
		mentionedList       []string
		mentionedMobileList []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WeworkClient_SendTextMessage",
			fields: fields{
				token: TestAccessToken,
			},
			args: args{
				content:             "这是一条企业微信的测试通知" + time.Now().String(),
				mentionedList:       []string{"\"@all\""},
				mentionedMobileList: []string{"13598055910", "@all"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token: tt.fields.token,
			}
			if err := client.SendTextMessage(tt.args.content, tt.args.mentionedList, tt.args.mentionedMobileList); (err != nil) != tt.wantErr {
				t.Errorf("SendTextMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewWeworkClient(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "WeworkClient_NewWeworkClient",
			args: args{
				token: TestAccessToken,
			},
			want: NewWeworkClient(TestAccessToken),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeworkClient(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeworkClient() = %v, want %v", got, tt.want)
			}
		})
	}
}