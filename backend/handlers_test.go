package backend

import (
	"testing"

	"github.com/213-team/recbot_backend/backend/core"
	"github.com/213-team/recbot_backend/dao/mongodb"
	"github.com/213-team/recbot_backend/subscriptionb"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/api/iterator"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type SubscriptionServiceTestCase struct {
	suite.Suite
}

func (suite *SubscriptionServiceTestCase) SetupTest() {
	viper.Set("mongo.dbname", "test_test")
}

func TestSubscriptionServiceTestCase(t *testing.T) {
	suite.Run(t, new(SubscriptionServiceTestCase))
}

func (suite *SubscriptionServiceTestCase) TestReadSubscriptions() {
	testChannelID1 := "testChannel1"
	testChannelID2 := "testChannel2"
	testUserID := "testUserId"
	subscriptionsColl, mongoCtx := mongodb.GetSubscriptionCollection()
	subscriptionsColl.InsertOne(mongoCtx, bson.M{
		"channel": bson.M{
			"id": testChannelID1,
		},
		"user": bson.M{
			"id": testUserID,
		},
	})
	subscriptionsColl.InsertOne(mongoCtx, bson.M{
		"channel": bson.M{
			"id": testChannelID2,
		},
		"user": bson.M{
			"id": testUserID,
		},
	})
	cursor, _ := core.ListSubscriptions(testUserID)
	var result []*subscriptionb.Subscription
	for {
		subscription, err := cursor.Next()
		if err == iterator.Done {
			break
		}
		result = append(result, subscription)
	}
}
