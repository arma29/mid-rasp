package queueServer

import (
)

type Publisher struct {
	host string
	port int
	publishQueues map[string]bool
}

type PublisherManager struct {
	pubList []Publisher
}

// Mark a host as a publisher
func (pubManager PublisherManager) PublishRequest(host string, port int, destination string) {

	if pubManager.pubList == nil {
		pubManager.pubList = make([]Publisher)
	}

	// Get or Create Publisher
	publisher := pubManager.getPublisher(host, port)

	if publisher == nil {
		publisher = Publisher{host: host, port: port, queuesSubscribed: {destination}}
		pubManager.pubList.append(publisher)
	}

	// Grant publishing permission
	Publisher.publishQueues[destination] = true
} 


func (pubManager PublisherManager) getPublisher(host string, port int) {

	publisher := nil

	for _, pub := range pubManager.pubList {
		if (pub.host == host && pub.port == port) {
			publisher = pub
		}
	}

	return publisher
}