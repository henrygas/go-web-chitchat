package main

import "fmt"

var data = &Dao{}

type Dao struct {

}

func (dao *Dao) UserByEmail(email string) (*User, error) {
	user := User{
		Email: "873565064@qq.com",
	}
	return &user, nil
}

func (dao *Dao) Threads() ([]*Thread, error) {
	var threads = make([]*Thread, 3)
	for i := 0; i < len(threads); i++ {
		threads[i] = &Thread{
			ThreadId: i + 1,
			Title: fmt.Sprintf("这是第%d封帖子", i + 1),
		}
	}
	return threads, nil
}

func (dao *Dao) Encrypt(in string) string {
	return in
}