package utils

import "calebsideras.com/temporary/src/db"

type TLinkContent struct {
	Title       string
	Description string
	Src         string
	Alt         string
	Href        string
	Boost       string
	Disabled    bool
}

// Config is a struct that holds configuration settings for the application.
type Config struct {
	AppConfig
	Datastore db.Datastore
}

type AppConfig struct {
	AppName       string
	APIKey        string
	OtherSettings map[string]string // Additional settings as a map for flexibility
}

// initConfig initializes the configuration with a Database.
func InitConfig(appConfig AppConfig, dbAddr string) Config {
	rdb := db.NewRedisDB(dbAddr)
	return Config{
		AppConfig: appConfig,
		Datastore: rdb,
	}
}
