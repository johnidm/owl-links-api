package database

import (
	"gopkg.in/pg.v3"
	//"fmt"
)

const (
	DB_USER     = "jtbrooyyxtkpsz"
	DB_PASSWORD = "kS_f5apof0j0MiHVUC4WP74Rrm"
	DB_NAME     = "d9o4h18p276q2s"
	DB_DATABASE = "ec2-107-21-102-69.compute-1.amazonaws.com"
	DB_PORT     = "5432"
)

func OpenDB() *pg.DB {
	return pg.Connect(&pg.Options{
		Host:     DB_DATABASE,
		Database: DB_NAME,
		User:     DB_USER,
		Port:     DB_PORT,
		Password: DB_PASSWORD,
		SSL:      true,
	})
}
