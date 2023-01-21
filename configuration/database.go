package configuration

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strconv"
	"time"
)

func NewDatabase(config Config) *sqlx.DB {
	username := config.Get("DATASOURCE_USERNAME")
	password := config.Get("DATASOURCE_PASSWORD")
	host := config.Get("DATASOURCE_HOST")
	port := config.Get("DATASOURCE_PORT")
	dbName := config.Get("DATASOURCE_DB_NAME")
	maxPoolOpen, err := strconv.Atoi(config.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPoolIdle, err := strconv.Atoi(config.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPollLifeTime, err := strconv.Atoi(config.Get("DATASOURCE_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	db, err := sqlx.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?parseTime=true")
	exception.PanicLogging(err)

	db.SetMaxOpenConns(maxPoolOpen)
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	return db
}
