package postgresql

import (
	"io/ioutil"
	"sync"

	"github.com/andresxlp/go-echo-postgres-hexa/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var (
	instance *sqlx.DB
	once     sync.Once
)

func ConnectPGInstance() *sqlx.DB {
	once.Do(func() {
		instance = getConnection()
		err := makeMigration(instance)
		if err != nil {
			log.Panic(err)
		}
	})
	return instance

}

func getConnection() *sqlx.DB {
	uri := configs.Environments().DatabaseUri
	connection, err := sqlx.Connect("postgres", uri)

	if err != nil {
		panic(err)
	}
	log.Infof("Postgresql successfully connected")
	return connection
}

func makeMigration(db *sqlx.DB) error {
	b, err := ioutil.ReadFile("./internal/utils/tables.sql")
	if err != nil {
		return err
	}
	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}
	return rows.Close()
}
