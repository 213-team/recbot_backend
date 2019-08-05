package backend

import (
	"context"
	"errors"
	"fmt"

	"github.com/213-team/recbot_backend/backend/core"
	"github.com/213-team/recbot_backend/subscriptionb"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//SubscriptionService manages subscripttions
type SubscriptionService struct {
}

//ReadSubscription returns information about a subscription of a user
func (s *SubscriptionService) ReadSubscription(_ context.Context, req *subscriptionb.ReadSubscriptionReq) (*subscriptionb.ReadSubscriptionRes, error) {
	return nil, errors.New("Not implemented")
}

//AddSubscription adds a new subsciption to a user
func (s *SubscriptionService) AddSubscription(_ context.Context, req *subscriptionb.AddSubscriptionReq) (*subscriptionb.AddSubscriptionRes, error) {
	core.AddSubscription(req.GetSubscription())
	return &subscriptionb.AddSubscriptionRes{Status: &subscriptionb.Status{Success: true}}, nil
}

//DeleteSubscription deletes a subscription for a user
func (s *SubscriptionService) DeleteSubscription(_ context.Context, req *subscriptionb.DeleteSubscriptionReq) (*subscriptionb.DeleteSubscriptionRes, error) {
	core.DeleteSubscription(req.GetSubscription())
	return &subscriptionb.DeleteSubscriptionRes{Status: &subscriptionb.Status{Success: true}}, nil
}

//ListSubscriptions lists all user subscriptions
func (s *SubscriptionService) ListSubscriptions(req *subscriptionb.ListSubscriptionsReq, stream subscriptionb.SubscriptionService_ListSubscriptionsServer) error {
	subscriptionsCursor, err := core.ListSubscriptions(req.User.Id)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	for {
		subscription, err := subscriptionsCursor.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		stream.Send(&subscriptionb.ListSubscriptionsRes{
			Subscription: subscription,
		})
	}
	return nil
}
