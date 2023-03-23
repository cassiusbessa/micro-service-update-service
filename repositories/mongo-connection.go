package repositories

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Repo = New()

type Repository struct {
	Client        *mongo.Client
	Ctx           context.Context
	CtxCancel     context.CancelFunc
	ClientOptions *options.ClientOptions
}

func New() *Repository {
	logrus.Warnf("Instantiating repository")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	options := options.Client().ApplyURI("mongodb+srv://omicronrpl:oIe3hKERFNXPtiyF@hsa.sqnczja.mongodb.net/?authSource=admin")
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		logrus.Fatalf("%v", err)
		panic(err)
	}
	defer logrus.Infof("Repository instance Created")
	return &Repository{
		Client:        client,
		Ctx:           ctx,
		CtxCancel:     cancel,
		ClientOptions: options,
	}
}

func (r *Repository) Ping() {
	logrus.Warn("Pinging MongoDB ...")
	if er := r.Client.Ping(r.Ctx, r.ClientOptions.ReadPreference); er != nil {
		logrus.Fatal("Mongo Database Unavailable")
		panic(er)
	}
	logrus.Infof("MongoDB connection established")
}
