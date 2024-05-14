package env

import (
	"os"

	"github.com/go-test-naka/go-place/log"
)

var DB_USER string
var DB_PASS string
var DB_HOST string
var DB_PORT string
var DB_SCHEMA string

func init() {

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_SCHEMA = os.Getenv("DB_SCHEMA")
	log.Info("DB_USER:" + DB_USER)
}
