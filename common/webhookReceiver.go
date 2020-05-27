// Common holds shared code.
package common

// A WebhookReceiver can receive a webhook.
type WebhookReceiver interface {
	ReceiveWebhook(artist, title string) error
}
