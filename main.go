// Deliver webhook notifications when shazam.com charts are updated.
package main

import (
	"fmt"
	"time"

	"os"
	"os/signal"
	"syscall"

	"github.com/bst27/shazam-webhook/chartlist"
	"github.com/bst27/shazam-webhook/config"
	"github.com/bst27/shazam-webhook/service/shazam"
	"github.com/bst27/shazam-webhook/service/webhook"
	"github.com/bst27/shazam-webhook/tracklist"
)

var (
	appConfig       *config.AppConfig
	webhookService  webhook.Service
	shazamClient    shazam.Client
	tracklistRepo   tracklist.Repository
	chartlistRepo   chartlist.Repository
	shutdownPending bool
)

func bootstrap() {
	appConfig = config.Get()
	webhookService = webhook.Get()
	shazamClient = shazam.New()
	tracklistRepo = tracklist.Get()
	chartlistRepo = chartlist.Get()
	shutdownPending = false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// Wait for signal to stop execution
		<-sigs

		fmt.Println("Stopping application ...")
		shutdownPending = true
	}()
}

func shutdown() {
	_ = webhookService.Close()
}

func main() {
	bootstrap()
	run()
	shutdown()
}

func run() {
	initChartlists()
	initTracklists()

	if !appConfig.SendInitialHooks {
		webhookService.Disable()
	}

	lastPolling := time.Time{}

	for !shutdownPending {
		time.Sleep(1 * time.Second)

		if lastPolling.Add(time.Duration(appConfig.PollingInterval) * time.Second).After(time.Now()) {
			continue
		}

		lastPolling = time.Now()

		for _, v := range chartlistRepo.GetAll() {
			req := shazamClient.CreateFetchChartsRequest(v.ShazamKey())

			tracks, err := req.Send()

			if err != nil {
				// silently ignore; will be updated in the next iteration
				continue
			}

			for _, t := range tracks {
				v.AddTrackIfMissing(t.Subtitle, t.Title)
			}
		}

		for _, c := range chartlistRepo.GetAll() {
			for _, t := range tracklistRepo.GetAll() {
				if t.IsWatching(c.ShazamKey()) {
					c.WriteTracks(t)
				}
			}
		}

		if !appConfig.SendInitialHooks {
			webhookService.Enable()
		}
	}
}

func initChartlists() {
	for _, shazamChartKey := range appConfig.MonitoredShazamCharts {
		chartlistRepo.Add(chartlist.New(shazamChartKey))
	}
}

func initTracklists() {
	for _, tracklistConfig := range appConfig.Tracklists {
		tl := tracklist.New()

		for _, shazamChartKey := range tracklistConfig.WatchedShazamCharts {
			tl.WatchShazamCharts(shazamChartKey)
		}

		for _, webhookTarget := range tracklistConfig.WebhookTargets {
			tl.RegisterWebhook(webhook.NewReceiver(webhookTarget.URL, webhookService))
		}

		tracklistRepo.Add(tl)
	}
}
