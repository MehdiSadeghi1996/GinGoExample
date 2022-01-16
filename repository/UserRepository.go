package repository

import (
	"Template/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]entity.Users, error)
}

type userRepository struct {
	client *mongo.Client
}

func NewUserRepository(cli *mongo.Client) UserRepository {
	return &userRepository{
		client: cli,
	}
}

func (r *userRepository) ListUsers(ctx context.Context) ([]entity.Users, error) {

	var users []entity.Users
	findOptions := options.Find()
	//Set the limit of the number of record to find
	findOptions.SetLimit(5)
	res, err := r.client.Database("TemplateGo").Collection("Users").Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	res.All(context.TODO(), &users)

	return users, nil
}
