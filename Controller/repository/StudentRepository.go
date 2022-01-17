package repository

import (
	"Template/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository interface {
	CreateStudnet(ctx context.Context, student entity.Student) error
	FindByName(ctx context.Context, name string) (*entity.Student, error)
	FindByAgeOne(ctx context.Context, age string) (*entity.Student, error)
	FindByGreaterAge(ctx context.Context, age string) (*[]entity.Student, error)
}

type studentRepository struct {
	client *mongo.Client
}

func NewStudentRepository(cli *mongo.Client) StudentRepository {
	return &studentRepository{
		client: cli,
	}
}

func (r *studentRepository) FindByName(ctx context.Context, name string) (*entity.Student, error) {

	var studnet entity.Student

	err := r.client.Database("TemplateGo").Collection("Students").FindOne(ctx, bson.M{"name": name}).Decode(&studnet)
	if err != nil {
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &studnet, nil
}

func (r *studentRepository) CreateStudnet(ctx context.Context, student entity.Student) error {

	_, err := r.client.Database("TemplateGo").Collection("Students").InsertOne(ctx, student)
	if err != nil {
		return err
	}
	return nil
}

func (r *studentRepository) FindByAgeOne(ctx context.Context, age string) (*entity.Student, error) {

	var studnet entity.Student
	err := r.client.Database("TemplateGo").Collection("Students").FindOne(ctx, bson.M{"age": age}).Decode(&studnet)
	if err != nil {
		return nil, err
	}

	return &studnet, nil
}
func (r *studentRepository) FindByGreaterAge(ctx context.Context, age string) (*[]entity.Student, error) {
	var studnet []entity.Student
	ress, err := r.client.Database("TemplateGo").Collection("Students").Find(ctx, bson.M{"age": bson.M{"$gt": age}})
	if err != nil {
		return nil, err
	}
	ress.All(context.TODO(), &studnet)

	return &studnet, nil
}
