package queueServer

type Subscriber struct {
	Host string
	Port int
	QueuesSubscribed map[string]bool
}

type SubscriberManager struct {
	SubList []Subscriber
}

// Subscribe to destination queue
func (subManager SubscriberManager) SubscribeRequest(host string, port int, destination string) {

	if subManager.SubList == nil {
		subManager.SubList = make([]Subscriber, 1)
	}

	// Get or Create Subscriber
	subscriber, exists := subManager.getSubscriber(host, port)

	if !exists {
		subscriber = Subscriber{Host: host, Port: port, QueuesSubscribed: make(map[string]bool)}
		subManager.SubList = append(subManager.SubList, subscriber)
	}

	// Subscribe
	subscriber.QueuesSubscribed[destination] = true
	
} 

// Subscribe to destination queue
func (subManager SubscriberManager) UnsubscribeRequest(host string, port int, destination string) {

	if subManager.SubList == nil {
		return 
	}

	// Get or Create Subscriber
	subscriber, exists := subManager.getSubscriber(host, port)

	if !exists{
		return
	}

	// Remove destination string from subscribe list
	subscriber.QueuesSubscribed[destination] = false	
} 

func (subManager SubscriberManager) getSubscriber(host string, port int) (Subscriber, bool) {

	var subscriber Subscriber
	exists := false

	for _, sub := range subManager.SubList {
		if (sub.Host == host && sub.Port == port) {
			subscriber = sub
			exists = true
		}
	}

	return subscriber, exists
}	