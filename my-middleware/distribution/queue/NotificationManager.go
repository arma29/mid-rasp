package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
)

const (
	NOTIFICATION_READY = 0
	NOTICATION_SENT = 1
)

type Notification struct {
	Host string
	Port int
	Status int
	Message message.Message
}

type NotificationManager struct{
	NotificationList []Notification
}

// Append notifications to NotificationList
func (manager *NotificationManager) addNotification(notif Notification) {

	if (manager.NotificationList == nil) {
		manager.NotificationList = make([]Notification, 0)
	}
	manager.NotificationList = append(manager.NotificationList, notif)
}


// Create new Notifications for all subscribers for that message destination
func (manager *NotificationManager) NewMessageToNotify(msg message.Message, subList []Subscriber) {

	for _, sub := range subList {
		notif := Notification{Host:sub.Host, Port:sub.Port, Status: NOTIFICATION_READY, Message:msg}
		manager.addNotification(notif)
	}
}


// Create new Notifications for all messages in the Queue that sub has subscribed
func (manager *NotificationManager) NewSubscriberToNotify(sub Subscriber, queue Queue) {
	
	msgList := queue.MsgList

	if len(msgList) > 0 {
		for _, msg := range msgList {
			notif := Notification{Host:sub.Host, Port:sub.Port, Status:NOTIFICATION_READY, Message: msg}
			manager.addNotification(notif)
		}
	}
}


// TODO: Function to clean old notifications
// TODO: Implement function to clean Old	 Messages From Queue
