package repositories

import (
	"context"
	"time"

	"github.com/joaorodrs/linker/internals/core/domain"
	"github.com/joaorodrs/linker/internals/core/ports"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const (
	MongoClientTimeout = 5
)

type LinkRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.LinkRepository = (*LinkRepository)(nil)

func NewLinkRepository(conn string) (*LinkRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(conn))

	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &LinkRepository{
		client:     client,
		database:   client.Database("linker"),
		collection: client.Database("linker").Collection("links"),
	}, nil
}

func (r *LinkRepository) CreateLink(URL string, title string) error {
	//Here your code for login in mongo database
	return nil
}

func (r *LinkRepository) GetLink(ID int) (domain.Link, error) {
	//Here your code for login in mongo database
	return domain.Link{}, nil
}
