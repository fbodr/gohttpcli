package lib

import (
	"github.com/spf13/viper"
	"os"
)

var contextFile *viper.Viper

func ContextInit(appCfgPath string) {
	var err error

	contextFile = viper.New()
	contextFile.SetConfigType("yaml")
	contextFile.AddConfigPath(appCfgPath)
	contextFile.SetConfigName("context")
	_, err = os.Stat(appCfgPath + "/context.yaml")
	if err != nil {
		emptyFile, cannotCreateErr := os.Create(appCfgPath + "/context.yaml")
		if cannotCreateErr != nil {
			panic(cannotCreateErr)
		} else {
			emptyFile.Close()
		}
	}
}

func ContextSet(key string, value string) {
	var err error

	contextFile.Set(key, value)
	err = contextFile.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func ContextHasKey(key string) bool {
	return len(contextFile.GetString("")) > 0
}

func ContextGetValue(key string) string {
	return contextFile.GetString(key)
}
