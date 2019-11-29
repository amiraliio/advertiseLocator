package configs

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func config() *mongo.Client {
	var config string
	if viper.GetString("DATABASES.MONGO.USERNAME") != "" && viper.GetString("DATABASES.MONGO.PASSWORD") != "" {
		config = "mongodb://" + viper.GetString("DATABASES.MONGO.USERNAME") + ":" + viper.GetString("DATABASES.MONGO.PASSWORD") + "@" + viper.GetString("DATABASES.MONGO.HOST") + ":" + viper.GetString("DATABASES.MONGO.PORT") + "/?authSource=" + viper.GetString("DATABASES.MONGO.DATABASE")
	} else {
		config = "mongodb://" + viper.GetString("DATABASES.MONGO.HOST") + ":" + viper.GetString("DATABASES.MONGO.PORT")
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
	return config().Database(viper.GetString("DATABASES.MONGO.DATABASE"))
}
