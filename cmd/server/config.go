package main

import (
	"context"
	"path"
	"strings"

	"github.com/spf13/viper"

	"github.com/ez-api/internal/app/server"
)

// Configuration executor configruation strcture
type Configuration struct {
	// LoggingConfig *log.Configuration
	ServerConfig *server.Configruation
}

// New Config Object
func Init(ctx context.Context, filepath string) (*viper.Viper, error) {

	dir, file := path.Split(filepath)
	if dir == "" {
		dir = "./"
	}
	ext := path.Ext(filepath)
	filename := strings.TrimSuffix(file, ext)

	if len(ext) < 2 {
		// logger.Fatalf("check config file path: %s", filepath)
	}

	conf := viper.New()
	conf.AddConfigPath(dir)
	conf.SetConfigType(ext[1:])
	conf.SetConfigName(filename)
	conf.SetConfigFile(filepath)

	// Find and read the config file
	if err := conf.ReadInConfig(); err != nil {
		// fmt.Fatalf("Error reading config file, %s", err)
		return nil, err
	}
	// Confirm which config file is used
	// logger.Infof("Config: %s\n", conf.ConfigFileUsed())

	return conf, nil
}

// GetConfigruation get configruation by file
func GetConfigruation(ctx context.Context, filepath string) (*Configuration, error) {

	conf, err := Init(ctx, filepath)
	if err != nil {
		return nil, err
	}

	// set default values
	// // logging
	// conf.SetDefault("logging.verbose", true)
	// conf.SetDefault("logging.disableColors", false)
	// server
	conf.SetDefault("server.port", 9001)
	// conf.SetDefault("server.jwtToken.secret", helper.GenerateRandomBytes(32))
	// conf.SetDefault("server.jwtToken.expirationTimeMilisecond", 6000*60*60*8)

	conf.SetDefault("server.httpServer.port", 9001)

	// conf.SetDefault("server.postgresql.userID", "")
	// conf.SetDefault("server.postgresql.userPassword", "")
	// conf.SetDefault("server.postgresql.url", "")
	// conf.SetDefault("server.postgresql.port", "5432")
	// conf.SetDefault("server.postgresql.database", "")
	// conf.SetDefault("server.postgresql.MaxConnectionCount", 10)
	// conf.SetDefault("server.postgresql.IdleConnectionCount", 100)

	return &Configuration{
		// LoggingConfig: &log.Configuration{
		// 	Verbose:       conf.GetBool("logging.verbose"),
		// 	DisableColors: conf.GetBool("logging.disableColors"),
		// 	ElasticSearchConfig: &log.ElasticSearchConfiguration{
		// 		Enabled:   conf.GetBool("logging.elasticSearch.enabled"),
		// 		URL:       conf.GetString("logging.elasticSearch.url"),
		// 		Sniff:     conf.GetBool("logging.elasticSearch.sniff"),
		// 		IndexName: conf.GetString("logging.elasticSearch.indexName"),
		// 	},
		// },

		ServerConfig: &server.Configruation{
			Port: conf.GetInt("server.port"),

			// 	DataSourceConfig: &datasource.Configuration{
			// 		NotiConfig: &slackbot.Configuration{
			// 			SlackBotToken: conf.GetString("notification.slackBotToken"),
			// 		},
			// 		JWTTokenConfig: &token.Configuration{
			// 			Secret:                   conf.GetString("server.jwtToken.secret"),
			// 			ExpirationTimeMilisecond: conf.GetInt("server.jwtToken.expirationTimeMilisecond"),
			// 		},
			// 		PostgresConfig: &postgres.Configuration{
			// 			UserID:              conf.GetString("server.postgresql.userID"),
			// 			UserPassword:        conf.GetString("server.postgresql.userPassword"),
			// 			URL:                 conf.GetString("server.postgresql.url"),
			// 			Port:                conf.GetInt("server.postgresql.port"),
			// 			Database:            conf.GetString("server.postgresql.database"),
			// 			MaxConnectionCount:  conf.GetInt("server.postgresql.MaxConnectionCount"),
			// 			IdleConnectionCount: conf.GetInt("server.postgresql.IdleConnectionCount"),
			// 		},
			// 	},
		},
	}, nil
}
