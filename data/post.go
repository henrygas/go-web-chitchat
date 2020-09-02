package data

import "time"

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (post *Post) Create() (err error) {
	statement := "INSERT INTO post(uuid, body, user_id, thread_id, created_at) " +
		"VALUES($1, $2, $3, $4, $5) returning id, uuid, created_at;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), post.Body, post.UserId, post.ThreadId, time.Now()).
		Scan(&post.Id, &post.Uuid, &post.CreatedAt)
	if err != nil {
		return
	}
	return
}
