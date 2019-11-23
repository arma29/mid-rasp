package queueServer

type Subscriber struct {
	host string
	port int
	queuesSubscribed map[string]bool
}

type SubscriberManager struct {
	subList []Subscriber
}

// Subscribe to destination queue
func (subManager SubscriberManager) SubscribeRequest(host string, port int, destination string) {

	if subManager.subList == nil {
		subManager.subList = make([]Subscriber)
	}

	// Get or Create Subscriber
	subscriber := subManager.getSubscriber(host, port)

	if subscriber == nil {
		subscriber = Subscriber{host: host, port: port, queuesSubscribed: make(map[string]bool)}
		subManager.subList.append(subscriber)
	}

	// Subscribe
	subscriber.queuesSubscribed[destination] = true
	
} 

// Subscribe to destination queue
func (subManager SubscriberManager) UnsubscribeRequest(host string, port int, destination string) {

	if subManager.subList == nil {
		return 
	}

	// Get or Create Subscriber
	subscriber := subManager.getSubscriber(host, port)

	if subscriber == nil {
		return
	}

	// Remove destination string from subscribe list
	subscriber.queuesSubscribed[destination] = false	
} 

func (subManager SubscriberManager) getSubscriber(host string, port int) {

	subscriber := nil

	for _, sub := range subManager.subList {
		if (sub.host == host && sub.port == port) {
			subscriber = sub
		}
	}

	return subscriber
}