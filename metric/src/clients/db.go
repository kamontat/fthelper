package clients

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/utils"
)

type Database struct {
	Enabled bool
	Type    string
	Cluster string

	context    context.Context
	connection *pgxpool.Pool
	url        string
	dbname     string
	username   string
	password   string

	logger *loggers.Logger
}

func (c *Database) Initial() error {
	if c.Type != "postgres" {
		return fmt.Errorf("current database only support 'postgres' type")
	}

	conn, err := pgxpool.Connect(
		c.context,
		// postgres://username:password@localhost:5432/database_name
		fmt.Sprintf("postgres://%s:%s@%s/%s", c.username, c.password, c.url, c.dbname),
	)
	if err != nil {
		return err
	}

	c.connection = conn
	return nil
}

func (c *Database) Cleanup() error {
	if c.connection != nil {
		c.connection.Close()
	}
	return nil
}

func (c *Database) Query(query string, args ...interface{}) pgx.Row {
	return c.connection.QueryRow(c.context, query, args...)
}

func (c *Database) Queries(query string, args ...interface{}) (pgx.Rows, error) {
	return c.connection.Query(c.context, query, args...)
}

func (c *Database) String() string {
	if !c.Enabled {
		return "disabled"
	}

	return fmt.Sprintf("%s (%s) '%s:%s'", c.url, c.Type, c.username, utils.MaskString(c.password, utils.MEDIUM))
}

func NewDatabase(cluster string, config maps.Mapper) (*Database, error) {
	var enabled = config.Bo("enabled", true)

	dbType, err := config.Se("type")
	if err != nil {
		return nil, err
	}

	url, err := config.Se("url")
	if err != nil {
		return nil, err
	}

	dbname, err := config.Se("name")
	if err != nil {
		return nil, err
	}

	username, err := config.Se("username")
	if err != nil {
		return nil, err
	}

	password, err := config.Se("password")
	if err != nil {
		return nil, err
	}

	return &Database{
		Enabled: enabled,
		Type:    dbType,
		Cluster: cluster,

		context:    context.Background(),
		connection: nil,
		url:        url,
		dbname:     dbname,
		username:   username,
		password:   password,

		logger: loggers.Get("client", "db", cluster),
	}, nil
}
