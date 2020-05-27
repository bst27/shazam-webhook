// Webhook offers a web service to deliver webhooks.
package webhook

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

const (
	concurrentRequests = 3
)

var (
	serviceInstance *service
)

// Service describes a web service to deliver webhooks.
type Service interface {
	Send(targetURL string, artist, title string) error
	Disable()
	Enable()
	Close() error
}

type service struct {
	reqChannel chan int
	reqWaitGrp sync.WaitGroup
	enabled    bool
}

// Get returns a singleton service.
func Get() Service {
	if serviceInstance == nil {
		serviceInstance = newService()
	}

	return serviceInstance
}

func newService() *service {
	return &service{
		reqChannel: make(chan int, concurrentRequests),
		reqWaitGrp: sync.WaitGroup{},
		enabled:    true,
	}
}

func (s *service) sendAsync(targetURL string, artist, title string) {
	defer s.reqWaitGrp.Done()

	resp, err := http.Post(
		targetURL,
		"application/json",
		bytes.NewBufferString(fmt.Sprintf(`{"value1":"%s","value2":"%s"}`, title, artist)),
	)

	<-s.reqChannel

	if err != nil {
		return
	}

	defer resp.Body.Close()
}

// Send delivers a webhook.
func (s *service) Send(targetURL string, artist, title string) error {
	if !s.enabled {
		return nil
	}

	s.reqChannel <- 1
	s.reqWaitGrp.Add(1)
	go s.sendAsync(targetURL, artist, title)
	return nil
}

func (s *service) Close() error {
	s.reqWaitGrp.Wait()
	return nil
}

func (s *service) Disable() {
	s.enabled = false
}

func (s *service) Enable() {
	s.enabled = true
}
