package queue

type QueueServer struct {
	QueueManager QueueManager
	PubManager PublisherManager
	SubManager SubscriberManager
	NotifManager NotificationManager
	EventNotifManager EventNotification
}


