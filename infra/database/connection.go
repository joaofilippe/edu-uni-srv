package database

import (
	"fmt"
	"github.com/joaofilippe/edu-uni-srv/common/logger"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joaofilippe/edu-uni-srv/config"
	_ "github.com/lib/pq" //necessary to config
	"gopkg.in/yaml.v3"
)

var (
	dbConfig   *DBConfig
	connection *DBConnection
)

type DBConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"dbname" env:"DB_NAME"`
	Dsn      string
}

type DBConnection struct {
	dbConfig     *DBConfig
	DBConnection *sqlx.DB
}

func NewConnection(log *logger.Logger, appConfig *appconfig.App) *DBConnection {
	setDBConfigFromYaml(log, appConfig)
	dbConfig.Dsn = dbConfig.getDsn()

	db, err := sqlx.Open("postgres", dbConfig.Dsn)
	if err != nil {
		log.Fatalf(fmt.Errorf("can't open connection to database: %w", err))
		return connection
	}

	connection = &DBConnection{
		dbConfig,
		db,
	}

	return connection
}

func (c *DBConfig) getDsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
	)

}

func setDBConfigFromYaml(log *logger.Logger, appConfig *appconfig.App) {
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/db.yaml", appConfig.ConfigPath))
	if err != nil {
		log.Fatalf(fmt.Errorf("can't load db.yaml file"))
	}

	dbConfig = &DBConfig{}

	err = yaml.Unmarshal(yamlFile, dbConfig)
	if err != nil {
		log.Fatalf(fmt.Errorf("can't unmarshal db.yaml file"))
	}
}
