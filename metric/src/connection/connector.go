package connection

import "github.com/kamontat/fthelper/shared/errors"

// Connector is a high level aggregated data from connection
type Connector interface {
	// Connector name (cluster name)
	Cluster() string

	// Do some startup process/job
	Initial() error

	// Do some cleanup process/job
	Cleanup() error

	// This return parent connector, use when current connector has parent
	// Otherwise might return nil value
	Parent() Connector

	// This function will allow child to use parent connector logic instead
	WithParent(parent Connector) Connector

	// Save data of connect name for use later in connect
	Save(name string, data interface{}) Connector

	// Connect to repository
	Connect(name string) (interface{}, error)

	// Call connect on every possible connect 'name'
	// This can be useful if you use with WithCache()
	ConnectAll() *errors.Handler

	// String represent current connector
	String() string
}
