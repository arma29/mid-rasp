package queue

type Subscriber struct {
	Host string
	Port int
	QueuesSubscribed map[string]bool
}

type SubscriberManager struct {
	SubList []Subscriber
}

// Subscribe to destination queue
func (subManager *SubscriberManager) SubscribeRequest(host string, port int, destination string) *Subscriber{

	if subManager.SubList == nil {
		subManager.SubList = make([]Subscriber, 0)
	}

	// Get or Create Subscriber
	subscriber, exists := subManager.getSubscriber(host, port)

	if !exists {
		subscriber = Subscriber{Host: host, Port: port, QueuesSubscribed: make(map[string]bool)}
		subManager.SubList = append(subManager.SubList, subscriber)
	}

	// Subscribe
	subscriber.QueuesSubscribed[destination] = true
	
	return &subscriber
} 


// Subscribe to destination queue
func (subManager SubscriberManager) UnsubscribeRequest(host string, port int, destination string) {

	if subManager.SubList != nil {
		// Get or Create Subscriber
		subscriber, exists := subManager.getSubscriber(host, port)
		if exists{
			subscriber.QueuesSubscribed[destination] = false	
		}
	}
} 


// Get Subscriber object from Host and Port
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


// Get Subscriber List attached to a Queue
func (subManager SubscriberManager) GetSubscribersByQueue(queueName string) []Subscriber {

	subscriberList := make([]Subscriber, 0)

	for _, sub := range subManager.SubList {
		if sub.QueuesSubscribed[queueName] {
			subscriberList = append(subscriberList, sub)
		}
	}
	return subscriberList
}