package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/AnibalDBXD/go-crud-api/utils"
	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

var instance *sql.DB

func getPostgressInstance() *sql.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			fmt.Println("Creating postgres instance")
			envs, error := utils.GetEnvs()
			if error != nil {
				panic(error)
			}
			conn, err := sql.Open("postgres", envs.DB_URL)
			if err != nil {
				log.Fatal("Error opening connection to database: ", err)
				panic(err)
			}
			instance = conn
		} else {
			fmt.Println("Postgres instance already created")
		}
		fmt.Println("Postgres instance already created")
	}
	fmt.Println("Returning postgres instance")
	return instance
}
