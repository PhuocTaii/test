package models

type Config struct {
	Service  ServiceConfig  `yaml:"service" json:"service"`
	Postgres PostgresConfig `yaml:"postgres" json:"postgres"`
}
