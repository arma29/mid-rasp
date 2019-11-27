package queue

import (
)

type Publisher struct {
	Host string
	Port int
	PublishQueues map[string]bool
}

type PublisherManager struct {
	PubList []Publisher
}

// Mark a host as a publisher
func (pubManager PublisherManager) PublishRequest(host string, port int, destination string) {

	if pubManager.PubList == nil {
		pubManager.PubList = make([]Publisher, 0)
	}

	// Get or Create Publisher
	publisher, exists := pubManager.GetPublisher(host, port)

	if !exists {
		publisher = Publisher{Host: host, Port: port, PublishQueues: make(map[string]bool)}
		pubManager.PubList = append(pubManager.PubList, publisher)
	}

	// Grant publishing permission
	publisher.PublishQueues[destination] = true
} 


func (pubManager PublisherManager) GetPublisher(host string, port int) (Publisher, bool) {

	var publisher Publisher
	exists := false

	for _, pub := range pubManager.PubList {
		if (pub.Host == host && pub.Port == port) {
			publisher = pub
			exists = true
		}
	}

	return publisher, exists
}