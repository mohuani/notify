# Notify

> Golang 推送通知 SDK (钉钉群机器人、飞书群机器人、企业微信群机器人、邮件)。


## 支持的平台

- [x] [钉钉群机器人](https://developers.dingtalk.com/document/app/custom-robot-access)
- [x] [飞书群机器人](https://www.feishu.cn/hc/zh-CN/articles/360024984973)
- [x] [企业微信群机器人](https://open.work.weixin.qq.com/api/doc/90000/90136/91770)
- [ ] [邮件]()

## 安装

### 环境要求
- go 1.19+

### 引入本包
```shell
go get -u github.com/mohuani/notify
```


## 使用

<details>
<summary><b>钉钉机器人</b></summary>


### 前置准备

请先阅读一遍官方的文档 [钉钉群机器人](https://developers.dingtalk.com/document/app/custom-robot-access)，熟悉里面的各种名词概念。

#### 支持的消息类型 

- [x] text类型
- [x] link类型
- [x] markdown类型
- [x] ActionCard类型
- [x] FeedCard类型

### 调用

```go

// Text Message


// Link Message


// Markdown Message


// Feed Card Message


// Single Action Card Message


// Btns Action Card Message



```

</details>


<details>
<summary><b>飞书群群机器人</b></summary>


### 前置准备

请先阅读一遍官方的文档 [飞书群机器人配置说明](https://www.feishu.cn/hc/zh-CN/articles/360024984973)，熟悉里面的各种名词概念。

#### 支持的消息类型
- [x] 文本 text
- [x] 富文本 post
- [x] 群名片 share_chat
- [x] 图片 image
- [ ] 消息卡片 interactive

### 调用

```go

```

</details>



<details>
<summary><b>企业微信群机器人</b></summary>


### 前置准备

请先阅读一遍官方的文档 [企业微信群机器人配置说明](https://developer.work.weixin.qq.com/document/path/91770)，熟悉里面的各种名词概念。

#### 支持的消息类型
- [x] 文本 text
- [x] markdown
- [x] 图片 image
- [x] 图文 news
- [ ] 模版卡片 template_card

### 调用

```go

```

</details>


</details>




## 特别感谢

本项目借鉴了下面开源项目的思路，在此表示感谢，排名不分先后。
- [notify](https://github.com/guanguans/notify)
- [dingtalk](https://github.com/blinkbean/dingtalk)
- [dingtalk-notify-go-sdk](https://github.com/JetBlink/dingtalk-notify-go-sdk)


## License

[MIT](https://opensource.org/licenses/MIT)


----
