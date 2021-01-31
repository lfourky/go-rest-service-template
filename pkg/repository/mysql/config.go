package mysql

type Config struct {
	Host         string
	Port         int `default:"3306"`
	DatabaseName string
	User         string
	Password     string
	LogMode      bool `default:"true"`
}
