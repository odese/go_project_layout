package config

import (
	// "strings"

	// log ".../backend/pkg/infrastructure/logger"
	// ".../backend/pkg/utils"
	// "github.com/spf13/viper"
)

// // Represents the instance of config file.
// var config *viper.Viper

// // Init, reads yml file and initializes config instance.
// func Init() {
// 	config = viper.New()
// 	config.SetConfigType("yml")
// 	config.AddConfigPath("configs")
// 	config.SetConfigName(getNameOfConfigFile())

// 	err := config.ReadInConfig()
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("Error on reading config file.")
// 	}
// }

// // getNameOfConfigFile, chooses proper config file name according to work environment.
// func getNameOfConfigFile() (fileName string) {
// 	log.Info().Str(utils.WorkEnvironmentKey, utils.WorkEnvironment).Send()
// 	fileName = strings.ToLower(utils.WorkEnvironment)
// 	return fileName
// }

// // Call, returns instance of config file.
// func Call() *viper.Viper {
// 	return config
// }
