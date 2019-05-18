package backend

import (
	"context"
	"errors"
	"fmt"

	"github.com/213-team/recbot_backend/subscriptionb"
)

//SubscriptionService manages subscripttions
type SubscriptionService struct {
}

//ReadSubscription returns information about a subscription of a user
func (s *SubscriptionService) ReadSubscription(_ context.Context, req *subscriptionb.ReadSubscriptionReq) (*subscriptionb.ReadSubscriptionRes, error) {
	fmt.Println(req.User.GetId())
	return &subscriptionb.ReadSubscriptionRes{
		Subscription: &subscriptionb.Subscription{
			Channel: &subscriptionb.Channel{Id: "CatsChannel"},
			User:    &subscriptionb.User{Id: "FakeUser"}}}, nil
}

//AddSubscription adds a new subsciption to a user
func (s *SubscriptionService) AddSubscription(context.Context, *subscriptionb.AddSubscriptionReq) (*subscriptionb.AddSubscriptionRes, error) {
	return nil, errors.New("Not implemented")
}

//DeleteSubscription deletes a subscription for a user
func (s *SubscriptionService) DeleteSubscription(context.Context, *subscriptionb.DeleteSubscriptionReq) (*subscriptionb.DeleteSubscriptionRes, error) {
	return nil, errors.New("Not implemented")
}

//ListSubscriptions lists all user subscriptions
func (s *SubscriptionService) ListSubscriptions(*subscriptionb.ListSubscriptionsReq, subscriptionb.SubscriptionService_ListSubscriptionsServer) error {
	return errors.New("Not implemented")
}
