package data

import (
	"context"
	"database/sql"

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

func InitDB() (*sql.DB, error) {
	dbURL := util.SetDBURL()
	db, err := sql.Open("libsql", dbURL)
	return db, err
}

func InitTurgo(d *sql.DB) *Turgo {
	return &Turgo{
		db:  d,
		ctx: context.Background(),
		Tb:  *tabby.New(),
	}
}

func (t *Turgo) InitTables() error {
	_, err := t.db.ExecContext(t.ctx,
		`CREATE TABLE IF NOT EXISTS collection (
			ID INT PRIMARY KEY,
			Topic VARCHAR(255),
			URL VARCHAR(255),
			Note VARCHAR(255)
		)`)
	return err
}

func (t *Turgo) CreateCollection(c Collection) error {
	_, err := t.db.ExecContext(
		t.ctx,
		"INSERT INTO collection (ID, Topic, URL, Note) VALUES (?, ?, ?, ?)",
		c.ID, c.Topic, c.URL, c.Note)
	return err
}

func (t *Turgo) ListTopics() ([]string, error) {
	var topic string
	var res []string
	rows, err := t.db.Query("SELECT Topic FROM collection")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&topic)
		res = append(res, topic)
		if err != nil {
			return res, err
		}
	}
	return res, err
}

func (t *Turgo) List() ([]string, []string, []string, error) {
	var url string
	var topic string
	var note string
	var URLs []string
	var topics []string
	var notes []string
	rows, err := t.db.Query("SELECT Topic, URL, Note FROM collection")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&topic, &url, &note)
		if err != nil {
			return topics, URLs, notes, err
		}
		URLs = append(URLs, url)
		notes = append(notes, note)
		topics = append(topics, topic)

	}

	return topics, URLs, notes, err

}

func (t *Turgo) FilterByTopic(topic string) ([]string, []string, error) {
	var url string
	var note string
	var URLs []string
	var notes []string
	rows, err := t.db.Query("SELECT URL, Note FROM collection WHERE Topic LIKE ?", topic)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&url, &note)
		if err != nil {
			return URLs, notes, err
		}
		URLs = append(URLs, url)
		notes = append(notes, note)
	}
	return URLs, notes, err
}

func (t *Turgo) PurgeCollection(id string) error {
	_, err := t.db.ExecContext(t.ctx,
		`DELETE FROM collection WHERE ID = ?`,
		id)
	return err
}
