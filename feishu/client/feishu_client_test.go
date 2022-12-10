package client

import (
	"github.com/mohuani/notify/feishu/message"
	"reflect"
	"testing"
	"time"
)

const (
	TestAccessToken = "278470f5-ec50-4860-a21a-edb11cc4c49a" // 使用自己的 AccessToken
	TestKeyWord     = "监控报警"                                 // 使用自己创建自定义机器人时，安全设置里面，填写的自定义关键词
)

func TestClient_send(t *testing.T) {
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
		{
			name: "FeiShuClient_send",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				msg: message.NewTextMessage("新更新提醒，监控报警" + time.Now().String()),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := client.Send(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewFeiShuClient(t *testing.T) {
	type args struct {
		token   string
		keyWork string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "FeiShuClient_NewFeiShuClient",
			args: args{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			want: NewFeiShuClient(TestAccessToken, TestKeyWord),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFeiShuClient(tt.args.token, tt.args.keyWork); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFeiShuClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SendTextMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "FeiShuClient_SendTextMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				text: "新更新提醒" + time.Now().String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := client.SendTextMessage(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("SendTextMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendPostMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		title   string
		content [][]message.PostMessageContentPostZhCnContent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "FeiShuClient_SendPostMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				title:   "我是一个标题",
				content: [][]message.PostMessageContentPostZhCnContent{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := client.SendPostMessage(tt.args.title, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("SendPostMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendImageMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		imageKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "FeiShuClient_SendImageMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				imageKey: "img_7ea74629-9191-4176-998c-2e603c9c5e8g",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := client.SendImageMessage(tt.args.imageKey); (err != nil) != tt.wantErr {
				t.Errorf("SendImageMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendShareChatMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	type args struct {
		shareChatId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "FeiShuClient_SendShareChatMessage",
			fields: fields{
				token:   TestAccessToken,
				keyWork: TestKeyWord,
			},
			args: args{
				shareChatId: "oc_f5b1a7eb27ae2c7b6adc2a74faf339ff",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := client.SendShareChatMessage(tt.args.shareChatId); (err != nil) != tt.wantErr {
				t.Errorf("SendShareChatMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendInteractiveMessage(t *testing.T) {
	type fields struct {
		token   string
		keyWork string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				token:   tt.fields.token,
				keyWork: tt.fields.keyWork,
			}
			if err := client.SendInteractiveMessage(); (err != nil) != tt.wantErr {
				t.Errorf("SendInteractiveMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
