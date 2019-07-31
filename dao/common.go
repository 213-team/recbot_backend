package dao

import (
	"github.com/213-team/recbot_backend/dao/interfaces"
	"github.com/213-team/recbot_backend/dao/mongodb"
)

// GetStorage returns an entity that is used for working with data storage
func GetStorage() interfaces.SubscriptionDAO {
	return mongodb.SubscriptionDAOMongoDB{}
}
