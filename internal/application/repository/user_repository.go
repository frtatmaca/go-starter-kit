package repository

import (
	"context"
	"go-starter-kit/internal/application/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	MongoClientTimeout = 5
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(conn string) *UserRepository {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		return nil
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}

	return &UserRepository{
		client:     client,
		database:   client.Database("database"),
		collection: client.Database("database").Collection("users"),
	}
}

func (r *UserRepository) Login(email string, password string) error {
	return nil
}

func (r *UserRepository) Register(email string, password string) error {
	return nil
}
