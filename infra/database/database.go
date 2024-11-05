package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joaofilippe/edu-uni-srv/config"
	"github.com/joaofilippe/edu-uni-srv/utils/logger"
	_ "github.com/lib/pq" //necessary to config
	"gopkg.in/yaml.v3"
)

var (
	configDB   *ConfigDB
	connection *Connection
)

type ConfigDB struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"dbname" env:"DB_NAME"`
	Dsn      string
}

type Connection struct {
	ConfigDB   *ConfigDB
	Connection *sqlx.DB
}

func (c *ConfigDB) getDsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
	)

}

func NewConnection(config *ConfigDB) *Connection {
	config.Dsn = config.getDsn()

	db, err := sqlx.Open("postgres", config.Dsn)
	if err != nil {
		log.Fatalf("can't connect to database. Error: %v", err)
		return connection
	}

	connection = &Connection{
		config,
		 db,
	}

	return connection
}

func GetConnection(log *logger.Logger, appConfig *config.App) *Connection {
	return GetConfigFromYaml(log, appConfig)
}

func GetConfigFromYaml(log *logger.Logger, appConfig *config.App) *Connection {
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/db.yaml", appConfig.ConfigPath))
	if err != nil {
		log.Fatalf(fmt.Errorf("can't load db.yaml file"))
	}

	configDB = &ConfigDB{}

	err = yaml.Unmarshal(yamlFile, configDB)
	if err != nil {
		log.Fatalf(fmt.Errorf("can't unmarshal db.yaml file"))
	}

	conn := NewConnection(configDB)

	if err := conn.Connection.Ping(); err != nil {
		log.Fatalf(fmt.Errorf("can't connect to database. Error: %w", err))
	}

	return conn
}
