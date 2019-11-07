package configs

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func config() *mongo.Client {
	var config string
	if os.Getenv("MONGO_USERNAME") != "" && os.Getenv("MONGO_PASSWORD") != "" {
		config = "mongodb://" + os.Getenv("MONGO_USERNAME") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	} else {
		config = "mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config))
	if err != nil {
		Server.Logger.Fatal(err.Error())
	}
	return client
}

//DB config
func DB() *mongo.Database {
	return config().Database(os.Getenv("MONGO_DATABASE"))
}
