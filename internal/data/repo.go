package data

import (
	"context"
	"database/sql"
	"log"

	"github.com/cheynewallace/tabby"
	_ "github.com/libsql/libsql-client-go/libsql"
	"github.com/musaubrian/turgo/internal/util"
)

type Turgo struct {
	db  *sql.DB
	ctx context.Context
	Tb  tabby.Tabby
}
type Collection struct {
	ID    string
	Topic string
	URL   string
	Note  string
}

func initDB() (*sql.DB, error) {
	dbURL := util.SetDBURL()
	db, err := sql.Open("libsql", dbURL)
	return db, err
}

func InitTurgo() *Turgo {
	d, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	return &Turgo{
		db:  d,
		ctx: context.Background(),
		Tb:  *tabby.New(),
	}
}

func (t *Turgo) InitTables() error {
	_, err := t.db.ExecContext(t.ctx,
		`CREATE TABLE IF NOT EXISTS collections (
			ID INT PRIMARY KEY,
			Topic VARCHAR(255),
			URL VARCHAR(1000),
			Note VARCHAR(500)
		)`)
	return err
}
