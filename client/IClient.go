package client

type Client struct {
	token   string
	message string
}

type IClient interface {
	Send(msg any) error
}
