package user

import (
	"context"
	"log"

	"github.com/ztjustin/questions_api/config"
	domain "github.com/ztjustin/questions_api/domain/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RestUserRepository interface {
	GetAll() ([]*domain.User, error)
}

type usersRepositoryImpl struct {
}

func NewUserRestUserRepository() RestUserRepository {
	return &usersRepositoryImpl{}
}

var (
	collection = config.ClientDB().Database("reports").Collection("users")
)

func (r *usersRepositoryImpl) GetAll() ([]*domain.User, error) {

	var results []*domain.User

	findOpts := options.Find()
	findOpts.SetLimit(15)

	cursor, err := collection.Find(context.Background(), bson.D{{}}, findOpts)

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var user domain.User
		err := cursor.Decode(&user)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &user)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.TODO())

	return results, nil

}
