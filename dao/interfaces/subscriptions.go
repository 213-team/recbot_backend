package interfaces

import (
	"github.com/213-team/recbot_backend/subscriptionb"
)

//SubscriptionDAO represents an interface for interacting with DP for subscriptions
type SubscriptionDAO interface {
	Add(subscription *subscriptionb.Subscription) error
	Delete(subscription *subscriptionb.Subscription) error
	GetOne(userID string, channelID string) (*subscriptionb.Subscription, error)
	GetAllByUserID(userID string) (SubscriptionCursor, error)
}

//SubscriptionCursor is an iterator for listing subscriptions
type SubscriptionCursor interface {
	Next() (*subscriptionb.Subscription, error)
}
