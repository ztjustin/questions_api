package user

import (
	"context"
	"log"

	"github.com/ztjustin/questions_api/config"
	domain "github.com/ztjustin/questions_api/domain/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	FindById(string) (*domain.User, error)
	Create(*domain.User) (*domain.User, error)
}

type usersRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
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

func (r *usersRepositoryImpl) FindById(id string) (*domain.User, error) {
	var user *domain.User

	objectId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		return nil, errId
	}

	err := collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *usersRepositoryImpl) Create(newUser *domain.User) (*domain.User, error) {

	result, err := collection.InsertOne(context.TODO(), &newUser)
	if err != nil {
		return nil, err
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	newUser.Id = objectID

	return newUser, nil

}
