package data

func (t *Turgo) CreateCollection(c Collection) error {
	_, err := t.db.ExecContext(
		t.ctx,
		"INSERT INTO collections (ID, Topic, URL, Note) VALUES (?, ?, ?, ?)",
		c.ID, c.Topic, c.URL, c.Note)
	return err
}

func (t *Turgo) ListTopics() ([]string, error) {
	var topic string
	var res []string
	rows, err := t.db.Query("SELECT Topic FROM collections")
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
	rows, err := t.db.Query("SELECT Topic, URL, Note FROM collections")
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
	rows, err := t.db.Query("SELECT URL, Note FROM collections WHERE Topic LIKE ?", topic)
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
		`DELETE FROM collections WHERE ID = ?`,
		id)
	return err
}
