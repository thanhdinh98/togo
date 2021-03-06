package client

import (
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
)

type DB struct {
	PConn *gorm.DB
}

var (
	instance *DB
	once     sync.Once
)

func NewDB() *DB {
	once.Do(func() {
		conn, err := gorm.Open(os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		instance = &DB{
			PConn: conn,
		}
	})
	return instance
}
