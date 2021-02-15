// +build integration

package integration_test

import (
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/lfourky/go-rest-service-template/pkg/app"
	"github.com/lfourky/go-rest-service-template/pkg/repository/postgres"
	"github.com/lfourky/go-rest-service-template/pkg/server"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	"github.com/lfourky/go-rest-service-template/pkg/service/mail"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	address string
	config  app.Config
)

func TestMain(m *testing.M) {
	loadConfig()

	// This is a blocking operation, so it needs to go in a separate go routine.
	go app.Run(config)

	// The server needs some time to boot up, so we'll wait a predefined time for it,
	// but otherwise fail if it takes too long.
	waitForAppOrCrash()

	os.Exit(m.Run())
}

func waitForAppOrCrash() {
	isReady := false
	for i := 0; i < 100; i++ {
		if isReady = isAppReady(); isReady {
			break
		}

		time.Sleep(time.Millisecond * 10)
	}

	// If it's still not up, abort.
	if !isReady {
		panic("timed out while waiting for server to boot up")
	}
}

func isAppReady() bool {
	res, err := http.Get(address + "/version")
	if err != nil {
		return false
	}

	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}

func loadConfig() {
	viper.SetConfigName("config_integration")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic("unable to read config: " + err.Error())
	}

	// Basically most values are hardcoded here, since there's no actual need
	// for them to be configurable. If such a need arises, then use config values appropriately, instead.
	config = app.Config{
		Info: app.InfoConfig{
			Version:     "integration-tests-version",
			BuildDate:   "integration-tests-build-date",
			ServiceName: "integration-tests-service-name",
			CommitHash:  "integration-tests-commit-hash",
		},
		Server: server.Config{
			Address:      viper.GetString("server.address"),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
			Debug:        true,
		},
		Mail: mail.Config{
			SMTPPort:      viper.GetUint("mail.smtp.port"),
			SMTPHost:      viper.GetString("mail.smtp.host"),
			SMTPUser:      "smtp-user",
			SMTPPassword:  "smtp-password",
			SenderName:    "integration-tests-sender-name",
			SenderAddress: "user@integrationtests",
		},
		PostgresConfig: postgres.Config{
			Host:         viper.GetString("database.host"),
			Port:         viper.GetInt("database.port"),
			DatabaseName: viper.GetString("database.name"),
			Username:     viper.GetString("database.user"),
			Password:     viper.GetString("database.password"),
			LogMode:      true,
		},
		Logger: log.Config{
			Level: "INFO",
			Type:  "text",
		},
	}

	address = "http://" + config.Server.Address
}
