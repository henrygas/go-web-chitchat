package data

import (
	"database/sql"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func (thread *Thread) User() (user *User, err error) {
	user = &User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at"+
		" FROM users WHERE id = $1;", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func (thread *Thread) CreatedAtDate() (createdAtDate string) {
	createdAtDate = formatTime(thread.CreatedAt)
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1;", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	_ = rows.Close()
	return
}

func (thread *Thread) Create() (err error) {
	statement := "INSERT INTO threads(uuid, topic, user_id, created_at) VALUES ($1, $2, $3, $4)" +
		"returning id, uuid, created_at;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.QueryRow(createUUID(), thread.Topic, thread.UserId, time.Now()).
		Scan(&thread.Id, &thread.Uuid, &thread.CreatedAt)
	return
}

func GetThreadByUuid(uuid string) (*Thread, error) {
	thread := Thread{}
	query := "SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid=$1;"
	stmt, err := Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &thread, nil
}

func (thread *Thread) Posts() (posts []*Post) {
	posts = make([]*Post, 0)
	stmt, err := Db.Prepare("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts" +
		" WHERE thread_id=$1")
	if err != nil {
		return
	}
	rows, err := stmt.Query(thread.Id)
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.Id, &post.Uuid, &post.Body,
			&post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			continue
		} else {
			posts = append(posts, &post)
		}
	}
	return
}
