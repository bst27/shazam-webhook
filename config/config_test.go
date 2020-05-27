package config

import (
	"bytes"
	"testing"
)

func TestLoad(t *testing.T) {
	json := `
		{
			"SendInitialHooks": true,
			"PollingInterval": 10800,
			"MonitoredShazamCharts": [
				"81310",
				"80321"
			],
			"Tracklists": [
				{
					"WatchedShazamCharts": [
						"81310",
						"80321"
					],
					"WebhookTargets": [
						{
							"URL": "https://eni6i5nghyfmd.x.pipedream.net"
						},
						{
							"URL": "https://webhook.site/444feec8-1e97-4693-bf93-f272071b8415"
						}
					]
				},
				{
					"WatchedShazamCharts": [
						"81310"
					],
					"WebhookTargets": [
						{
							"URL": "https://2047ebf9-7c2c-4f00-8734-49044499d5f3.requestcatcher.com/test"
						}
					]
				}
			]
		}
	`

	appConfig, err := load(bytes.NewBufferString(json).Bytes())

	if err != nil {
		t.Errorf("Has error; want json")
		return
	}

	if appConfig.SendInitialHooks != true {
		t.Errorf("SendInitialHooks is false; want true")
	}

	if got := appConfig.PollingInterval; got != 10800 {
		t.Errorf("PollingInterval is %d; want 10800", got)
	}

	if got := len(appConfig.MonitoredShazamCharts); got != 2 {
		t.Errorf("MonitoredShazamCharts has %d items; want 2", got)
	} else {
		if got := appConfig.MonitoredShazamCharts[0]; got != "81310" {
			t.Errorf("MonitoredShazamCharts[0] is %s; want 81310", got)
		}

		if got := appConfig.MonitoredShazamCharts[1]; got != "80321" {
			t.Errorf("MonitoredShazamCharts[1] is %s; want 80321", got)
		}
	}

	if got := len(appConfig.Tracklists); got != 2 {
		t.Errorf("Tracklists has %d items; want 2", got)
	} else {
		if got := len(appConfig.Tracklists[0].WatchedShazamCharts); got != 2 {
			t.Errorf("Tracklists[0].WatchedShazamCharts has %d items; want 2", got)
		} else {
			if got := appConfig.Tracklists[0].WatchedShazamCharts[0]; got != "81310" {
				t.Errorf("Tracklists[0].WatchedShazamCharts[0] is %s; want 81310", got)
			}

			if got := appConfig.Tracklists[0].WatchedShazamCharts[1]; got != "80321" {
				t.Errorf("Tracklists[0].WatchedShazamCharts[1] is %s; want 80321", got)
			}
		}

		if got := len(appConfig.Tracklists[0].WebhookTargets); got != 2 {
			t.Errorf("Tracklists[0].WebhookTargets has %d items; want 2", got)
		} else {
			if got := appConfig.Tracklists[0].WebhookTargets[0].URL; got != "https://eni6i5nghyfmd.x.pipedream.net" {
				t.Errorf("Tracklists[0].WebhookTargets[0].URL is %s; want https://eni6i5nghyfmd.x.pipedream.net", got)
			}

			if got := appConfig.Tracklists[0].WebhookTargets[1].URL; got != "https://webhook.site/444feec8-1e97-4693-bf93-f272071b8415" {
				t.Errorf("Tracklists[0].WebhookTargets[1].URL is %s; want https://webhook.site/444feec8-1e97-4693-bf93-f272071b8415", got)
			}
		}

		// ---

		if got := len(appConfig.Tracklists[1].WatchedShazamCharts); got != 1 {
			t.Errorf("Tracklists[1].WatchedShazamCharts has %d items; want 1", got)
		} else {
			if got := appConfig.Tracklists[1].WatchedShazamCharts[0]; got != "81310" {
				t.Errorf("Tracklists[1].WatchedShazamCharts[0] is %s; want 81310", got)
			}
		}

		if got := len(appConfig.Tracklists[1].WebhookTargets); got != 1 {
			t.Errorf("Tracklists[1].WebhookTargets has %d items; want 1", got)
		} else {
			if got := appConfig.Tracklists[1].WebhookTargets[0].URL; got != "https://2047ebf9-7c2c-4f00-8734-49044499d5f3.requestcatcher.com/test" {
				t.Errorf("Tracklists[1].WebhookTargets[0].URL is %s; want https://2047ebf9-7c2c-4f00-8734-49044499d5f3.requestcatcher.com/test", got)
			}
		}
	}
}
