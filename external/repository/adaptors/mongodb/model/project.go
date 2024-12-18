package mongo_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	OwnerId     string             `bson:"owner_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Flow        string             `bson:"flow"`
}
