package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joaorodrs/linker/internals/core/domain"
	"github.com/joaorodrs/linker/internals/core/ports"
	"github.com/joaorodrs/linker/internals/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (r *LinkRepository) CreateLink(URL string) error {
	hash := helpers.GenerateHash(URL)

	newURL := fmt.Sprintf("/r/%s", hash)

	id := uuid.New().String()

	link, err := domain.NewLink(id, URL, newURL)
	if err != nil {
		return fmt.Errorf("[Link]: error creating new link: %w", err)
	}

	document := bson.M{
		"id":           link.ID,
		"URL":          link.URL,
		"shortenedUrl": link.ShortenedURL,
	}

	_, err = r.collection.InsertOne(context.Background(), document)
	if err != nil {
		return fmt.Errorf("[MongoDB]: error creating new link: %w", err)
	}

	return nil
}

func (r *LinkRepository) GetLink(hash string) (domain.Link, error) {
	filter := bson.M{"shortenedUrl": fmt.Sprintf("/r/%s", hash)}

	var link domain.Link
	err := r.collection.FindOne(context.TODO(), filter).Decode(&link)
	if err == mongo.ErrNoDocuments {
		return domain.Link{}, fmt.Errorf("Link not found")
	} else if err != nil {
		return domain.Link{}, fmt.Errorf("failed to query MongoDB: %w", err)
	}

	log.Printf("Link found: %v", link)

	return link, nil
}

func (r *LinkRepository) GetAllLinks() ([]domain.Link, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to query MongoDB: %w", err)
	}
	defer cursor.Close(context.TODO())

	var links []domain.Link
	if err := cursor.All(context.TODO(), &links); err != nil {
		return nil, fmt.Errorf("failed to decode links: %w", err)
	}

	return links, nil
}
