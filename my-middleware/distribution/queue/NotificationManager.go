package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
	// "time"
)

type Notification struct {
	Host    string
	Port    int
	Message message.Message
}

type NotificationManager struct{}

// Create new Notifications for all subscribers for that message destination
func (manager *NotificationManager) NewMessageToNotify(msg message.Message, subList []Subscriber) []Notification {

	notifList := make([]Notification, 0)

	for _, sub := range subList {
		notif := Notification{Host: sub.Host, Port: sub.Port, Message: msg}
		notifList = append(notifList, notif)
	}

	return notifList
}


// Create new Notifications for all messages in the Queue that sub has subscribed
func (manager *NotificationManager) NewSubscriberToNotify(sub Subscriber, queue Queue) []Notification {

	msgList := queue.MsgList
	notifList := make([]Notification, 0)

	if len(msgList) > 0 {
		for _, msg := range msgList {
			notif := Notification{Host:sub.Host, Port:sub.Port, Message: msg}
			notifList = append(notifList, notif)
		}
	}

	return notifList
}


// TODO: Function to clean old notifications
// TODO: Implement function to clean Old	 Messages From Queue
