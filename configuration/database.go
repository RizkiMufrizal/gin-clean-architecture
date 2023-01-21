package configuration

import (
	"database/sql"
	"github.com/RizkiMufrizal/gin-clean-architecture/common"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
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

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	exception.PanicLogging(err)

	db = sqldblogger.OpenDriver(dsn, db.Driver(), common.NewLogrusAdapter(common.NewLogger()))

	sqlxDB := sqlx.NewDb(db, "mysql")

	sqlxDB.SetMaxOpenConns(maxPoolOpen)
	sqlxDB.SetMaxIdleConns(maxPoolIdle)
	sqlxDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	return sqlxDB
}
