package main

// nolint: gci //Remove after making a decision postgres / mysql and removing the other import
import (
	"os"
	"strings"

	"github.com/lfourky/go-rest-service-template/pkg/app"
	"github.com/lfourky/go-rest-service-template/pkg/repository/postgres"
	"github.com/lfourky/go-rest-service-template/pkg/server"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	"github.com/lfourky/go-rest-service-template/pkg/service/mail"
	"github.com/spf13/viper"

	// nolint: godot
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	serviceName = "Go REST service"
)

// nolint: gochecknoglobals
var (
	version    = "/"
	buildDate  = "/"
	commitHash = "/"
)

func main() {
	loadConfig()

	hostname, err := os.Hostname()
	if err != nil {
		panic("unable to retrieve hostname: " + err.Error())
	}

	cfg := app.Config{
		Info: app.InfoConfig{
			Version:     version,
			BuildDate:   buildDate,
			ServiceName: serviceName,
			CommitHash:  commitHash,
		},
		Server: server.Config{
			Address:      viper.GetString("server.address"),
			ReadTimeout:  viper.GetDuration("server.read_timeout"),
			WriteTimeout: viper.GetDuration("server.write_timeout"),
			Debug:        viper.GetBool("server.debug"),
			CORS: server.CORSConfig{
				AllowCredentials: viper.GetBool("server.cors.allow_credentials"),
				Headers:          viper.GetStringSlice("server.cors.headers"),
				Methods:          viper.GetStringSlice("server.cors.methods"),
				Origins:          viper.GetStringSlice("server.cors.origins"),
			},
		},
		Mail: mail.Config{
			SMTPPort:      viper.GetUint("mail.smtp.port"),
			SMTPHost:      viper.GetString("mail.smtp.host"),
			SMTPUser:      viper.GetString("mail.smtp.user"),
			SMTPPassword:  viper.GetString("mail.smtp.password"),
			SenderName:    viper.GetString("mail.sender_name"),
			SenderAddress: viper.GetString("mail.sender_address"),
		},
		PostgresConfig: postgres.Config{
			Host:         viper.GetString("database.host"),
			Port:         viper.GetInt("database.port"),
			DatabaseName: viper.GetString("database.name"),
			Username:     viper.GetString("database.user"),
			Password:     viper.GetString("database.password"),
			LogMode:      viper.GetBool("database.log_mode"),
		},
		Logger: log.Config{
			Level: viper.GetString("log.level"),
			Type:  viper.GetString("log.type"),
			DefaultFields: log.DefaultFields{
				Service:  serviceName,
				Version:  version,
				Build:    commitHash,
				Hostname: hostname,
			},
		},
	}

	app.Run(cfg)
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic("unable to read config: " + err.Error())
	}
}
