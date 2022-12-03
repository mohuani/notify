# Notify

> Golang 推送通知 SDK (钉钉群机器人、飞书群机器人、企业微信群机器人、邮件)。


## 支持的平台

- [x] [钉钉群机器人](https://developers.dingtalk.com/document/app/custom-robot-access)
- [ ] [飞书群机器人](https://www.feishu.cn/hc/zh-CN/articles/360024984973)
- [ ] [企业微信群机器人](https://open.work.weixin.qq.com/api/doc/90000/90136/91770)
- [ ] [邮件]()

## 安装

### 环境要求
- go 1.19+

### 引入本包
```shell
go install github.com/mohuani/notify
```


## 使用

<details>
<summary><b>钉钉机器人</b></summary>


### 前置准备

 请先阅读一遍官方的文档 [钉钉群机器人](https://developers.dingtalk.com/document/app/custom-robot-access)，熟悉里面的各种名词概念。

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

## 特别感谢

本项目借鉴了下面开源项目的思路，在此表示感谢，排名不分先后。
- [notify](https://github.com/guanguans/notify)
- [dingtalk](https://github.com/blinkbean/dingtalk)
- [dingtalk-notify-go-sdk](https://github.com/JetBlink/dingtalk-notify-go-sdk)


## License

[MIT](https://opensource.org/licenses/MIT)


----
