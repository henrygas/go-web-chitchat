package main

import (
	"fmt"
	"go-web-chitchat/data"
	"html/template"
	"net/http"
)

// GET /
// 作用：导航栏根据是否有session做不同显示，内容栏展示帖子列表和创建帖子按钮
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		return
	}

	_, err = session(w, r)
	if err != nil {
		generateHTML(w, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, threads, "layout", "private.navbar", "index")
	}
}

// GET /err?msg=
// 作用：展示导航栏，错误信息页面
func err(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, msg, "layout", "public.navbar", "error")
	} else {
		generateHTML(w, msg, "layout", "private.navbar", "error")
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(w, "layout", data)
}
