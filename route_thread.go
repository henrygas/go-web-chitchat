package main

import (
	"fmt"
	"go-web-chitchat/data"
	"net/http"
)

// GET /thread/new
// 作用: 若session不存在, 302到/login
//      若session存在, 展示导航栏和编辑新帖子页面
func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
// 作用: 若session不存在, 302到/login
//      若session存在, 创建帖子记录并302到/
func createThread(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/thread/new", 302)
		return
	}

	userId := session.UserId
	topic := r.PostFormValue("topic")
	thread := &data.Thread{
		Topic:  topic,
		UserId: userId,
	}
	err = thread.Create()
	if err != nil {
		return
	}
	http.Redirect(w, r, "/", 302)
	return
}

// POST /thread/post
// 作用: 若session不存在, 则302到/login
//      若session存在, 则从body中取出帖子编号uuid和回复内容, 并创建post记录, 然后302到/thread/read?id=
func postThread(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil || session.Uuid == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/", 302)
		return
	}

	uuid := r.PostFormValue("uuid")
	body := r.PostFormValue("body")
	if uuid == "" || body == "" {
		http.Redirect(w, r, "/", 302)
		return
	}

	thread, err := data.GetThreadByUuid(uuid)
	if err != nil || thread.Uuid == "" {
		http.Redirect(w, r, "/", 302)
		return
	}

	post := data.Post{
		Body:     body,
		UserId:   session.UserId,
		ThreadId: thread.Id,
	}
	err = post.Create()
	if err != nil {
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/thread/read?id=%s", thread.Uuid), 302)
	return
}

// GET /thread/read?id=
// 作用: 若指定帖子ID不存在, 则显示错误信息页面
//      否则, 若sesion不存在, 则显示帖子详情信息页
//           若session存在, 则显示帖子详情信息页
func readThread(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("id")
	thread, err := data.GetThreadByUuid(uuid)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/err?msg=%s", err), 302)
		return
	}

	if thread.Uuid == "" {
		http.Redirect(w, r, fmt.Sprintf("/err?msg=%s", errNoThread), 302)
		return
	}

	_, err = session(w, r)
	if err != nil {
		generateHTML(w, thread, "layout", "public.navbar", "public.thread")
	} else {
		generateHTML(w, thread, "layout", "private.navbar", "private.thread")
	}
}
