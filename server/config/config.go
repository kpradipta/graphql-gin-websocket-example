package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() error {

	configFile := flag.String("c", "", "configuration file")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	fileName := *configFile

	splitsStr := strings.Split(filepath.Base(fileName), ".")
	viper.SetConfigName(filepath.Base(splitsStr[0]))
	viper.AddConfigPath(filepath.Dir(fileName))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Cannot read config file")
		return err
	}
	return nil
}

func keyIsExist(key string) {
	if !viper.IsSet(key) {
		fmt.Printf("Key %s not found; Check It Again \n", key)
		os.Exit(1)
	}
}

func MustGetString(key string) string {
	keyIsExist(key)
	return viper.GetString(key)
}

func MustGetInt(key string) int {
	keyIsExist(key)
	return viper.GetInt(key)
}

func MustGetBool(key string) bool {
	keyIsExist(key)
	return viper.GetBool(key)
}
