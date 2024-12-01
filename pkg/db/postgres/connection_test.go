package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Connection.ToPostgresConnectionString when all parameters are set.
func TestToPostgresConnectionString_AllFields(t *testing.T) {
	conn := Connection{
		Host:        "localhost",
		Port:        5432,
		Database:    "test_db",
		User:        "user",
		Password:    "password",
		TimeZone:    "UTC",
		SSLMode:     "require",
		SSLCert:     "/path/to/cert",
		SSLKey:      "/path/to/key",
		SSLRootCert: "/path/to/rootcert",
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=UTC&sslmode=require&sslcert=/path/to/cert&sslkey=/path/to/key&sslrootcert=/path/to/rootcert"
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}

// Test Connection.ToPostgresConnectionString when SSL mode is disabled.
func TestToPostgresConnectionString_SSLOff(t *testing.T) {
	conn := Connection{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		User:     "user",
		Password: "password",
		SSLMode:  Disable, // SSL is disabled
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=Asia/Ho_Chi_Minh&sslmode=disable"
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}

// Test Connection.ToPostgresConnectionString when TimeZone is not set.
func TestToPostgresConnectionString_DefaultTimeZone(t *testing.T) {
	conn := Connection{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		User:     "user",
		Password: "password",
		SSLMode:  "require",
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=Asia/Ho_Chi_Minh&sslmode=require"
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}

// Test Connection.ToPostgresConnectionString when no SSL certificates are provided.
func TestToPostgresConnectionString_NoSSLCerts(t *testing.T) {
	conn := Connection{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		User:     "user",
		Password: "password",
		SSLMode:  "require",
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=Asia/Ho_Chi_Minh&sslmode=require"
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}

// Test Connection.ToPostgresConnectionString when only SSL root certificate is provided.
func TestToPostgresConnectionString_SSLRootCert(t *testing.T) {
	conn := Connection{
		Host:        "localhost",
		Port:        5432,
		Database:    "test_db",
		User:        "user",
		Password:    "password",
		SSLMode:     "verify-full",
		SSLRootCert: "/path/to/rootcert",
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=Asia/Ho_Chi_Minh&sslmode=verify-full&sslrootcert=/path/to/rootcert"
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}

// Test Connection.ToPostgresConnectionString when both SSL cert and key are missing.
func TestToPostgresConnectionString_MissingSSLCertKey(t *testing.T) {
	conn := Connection{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		User:     "user",
		Password: "password",
		SSLMode:  "require",
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=Asia/Ho_Chi_Minh&sslmode=require"
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}

// Test Connection.ToPostgresConnectionString when there is no SSL mode.
func TestToPostgresConnectionString_NoSSLMode(t *testing.T) {
	conn := Connection{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		User:     "user",
		Password: "password",
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=Asia/Ho_Chi_Minh&sslmode="
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}

// Test Connection.ToPostgresConnectionString when SSLMode is empty.
func TestToPostgresConnectionString_EmptySSLMode(t *testing.T) {
	conn := Connection{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		User:     "user",
		Password: "password",
		SSLMode:  "",
	}

	expected := "postgresql://user:password@localhost/test_db?TimeZone=Asia/Ho_Chi_Minh&sslmode="
	assert.Equal(t, expected, conn.ToPostgresConnectionString())
}
