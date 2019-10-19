package configs

import (
	"log"
	"os"

	"github.com/amiraliio/advertiselocator/helpers"
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
	context, _ := helpers.TimeOut(10)
	client, err := mongo.Connect(context, options.Client().ApplyURI(config))
	//TODO mongo defer close
	if err != nil {
		log.Println(err.Error())
	}
	return client
}

//DB config
func DB() *mongo.Database {
	return config().Database(os.Getenv("MONGO_DATABASE"))
}
