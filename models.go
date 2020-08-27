package main

import "time"

type User struct {
	UserId int
	Email string
	Password string
}
func (u *User) CreateSession() Session {
	return Session{}
}


type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

func (s *Session) Check() (bool, error) {
	return true, nil
}

type Thread struct {
	ThreadId int
	Title string
}

type Post struct {
	ThreadId int
	Content string
}
