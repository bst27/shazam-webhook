package webhook

type Receiver struct {
	url     string
	webhook Service
}

func NewReceiver(url string, webhook Service) *Receiver {
	return &Receiver{
		url:     url,
		webhook: webhook,
	}
}

func (r Receiver) ReceiveWebhook(artist, title string) error {
	return r.webhook.Send(r.url, artist, title)
}
