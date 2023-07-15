package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type env struct {
	KafkaTopic     string   `mapstructure:"PIPELINE_KAFKA_TOPIC"`
	KafkaServers   []string `mapstructure:"PIPELINE_KAFKA_SERVERS"`
	KafkaPartition int      `mapstructure:"PIPELINE_KAFKA_PARTITION"`
	KafkaGroupId   string   `mapstructure:"PIPELINE_KAFKA_GROUP_ID"`
}

var Env *env

func init() {
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Str("err", err.Error()).Msg("load env to struct error")
		os.Exit(1)
	}

	Env = &env{}
	if err := viper.Unmarshal(Env); err != nil {
		log.Error().Str("err", err.Error()).Msg("load env to struct error")
		os.Exit(1)
	}
}
