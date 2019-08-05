package core

import (
	"github.com/213-team/recbot_backend/dao"
	"github.com/213-team/recbot_backend/dao/interfaces"
	"github.com/213-team/recbot_backend/subscriptionb"
)

//AddSubscription is an adding subscription core routine
func AddSubscription(subscription *subscriptionb.Subscription) error {
	return dao.GetStorage().Add(subscription)
}

//DeleteSubscription is a deleting subscription core routine
func DeleteSubscription(subscription *subscriptionb.Subscription) error {
	return dao.GetStorage().Delete(subscription)
}

//ListSubscriptions is a listing subscriptions core routine
func ListSubscriptions(userID string) (interfaces.SubscriptionCursor, error) {
	return dao.GetStorage().GetAllByUserID(userID)
}
