package main

import (
	"go-web-chitchat/data"
	"net/http"
)

// POST /authenticate
// 作用: 检验账密, 创建session和cookie, 并302到/;
//      若解析POST表单失败, 则跳转到/;
//      若密码验证失败, 则跳转到/login;
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/", 302)
		return
	}
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, _ := user.CreateSession()
		cookie := http.Cookie{
			Name:     cookieName,
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// GET /login
// 作用: 展示登陆输入框
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "login")
}

// GET /logout
// 作用: 删除session, 302到/
func logout(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err == nil {
		_ = session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}

// GET /signup
// 作用: 展示注册输入框
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "signup")
}

// POST /signup_account
// 作用: 创建user记录, 302到/login
func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/signup", 302)
		return
	}

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	if name == "" || email == "" || password == "" {
		http.Redirect(w, r, "/signup", 302)
		return
	}

	user := &data.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err = user.Create()
	if err != nil {
		http.Redirect(w, r, "/signup", 302)
		return
	}

	http.Redirect(w, r, "/login", 302)
}
