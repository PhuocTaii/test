package postgres

import (
	"fmt"
	"time"
)

type Connection struct {
	Host                  string
	Port                  int64
	Database              string
	User                  string
	Password              string
	TimeZone              string
	SSLCert               string
	SSLKey                string
	SSLRootCert           string
	SSLMode               SSLMode
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxIdleTime time.Duration
	ConnectionMaxLifeTime time.Duration
	ConnectionTimeout     time.Duration
}

// ToPostgresConnectionString returns the Postgres connection string based on the Connection struct.
func (conn Connection) ToPostgresConnectionString() string {
	// Use a default timezone if none is provided.
	if conn.TimeZone == "" {
		conn.TimeZone = "Asia/Ho_Chi_Minh" // Default timezone
	}

	// Start with the base connection string.
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?TimeZone=%s&sslmode=%s",
		conn.User,
		conn.Password,
		conn.Host,
		conn.Database,
		conn.TimeZone,
		conn.SSLMode,
	)

	// Append SSL parameters if SSL mode is not "disable".
	if conn.SSLMode != Disable {
		if conn.SSLCert != "" {
			dsn += fmt.Sprintf("&sslcert=%s", conn.SSLCert)
		}
		if conn.SSLKey != "" {
			dsn += fmt.Sprintf("&sslkey=%s", conn.SSLKey)
		}
		if conn.SSLRootCert != "" {
			dsn += fmt.Sprintf("&sslrootcert=%s", conn.SSLRootCert)
		}
	}

	return dsn
}
