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

	// strict mode will force application to panic stop if connection to database is failed
	strictMode bool
	logger     *loggers.Logger
}

// If error occurred
// 1. throw error if strict mode
// 2. log error message and disabled database silently
func (c *Database) handleError(err error) error {
	if c.strictMode {
		return err
	}

	c.Enabled = false
	c.logger.Error(err.Error())
	return nil
}

func (c *Database) Initial() error {
	if !c.Enabled {
		c.logger.Debug("not initial database because it disabled")
		return nil
	}

	if c.Type != "postgres" {
		var err = fmt.Errorf("current database only support 'postgres' type")
		return c.handleError(err)
	}

	conn, err := pgxpool.Connect(
		c.context,
		// postgres://username:password@localhost:5432/database_name
		fmt.Sprintf("postgres://%s:%s@%s/%s", c.username, c.password, c.url, c.dbname),
	)
	if err != nil {
		return c.handleError(err)
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

// TODO: result empty result if database is disabled
func (c *Database) Query(query string, args ...interface{}) pgx.Row {
	return c.connection.QueryRow(c.context, query, args...)
}

// TODO: result empty result if database is disabled
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

		strictMode: config.Bi("strict-mode"),
		logger:     loggers.Get("client", "db", cluster),
	}, nil
}
