package helpers

import (
	"github.com/spf13/viper"
)

func Path(dirName string) string {
	switch dirName {
	case "root":
		return viper.GetString("ROOT_PATH")
	case "storage":
		return viper.GetString("ROOT_PATH") + "/storage"
	case "temp":
		return viper.GetString("ROOT_PATH") + "/storage/temp"
	case "media":
		return viper.GetString("ROOT_PATH") + "/storage/media"
	default:
		return viper.GetString("ROOT_PATH")
	}
}
