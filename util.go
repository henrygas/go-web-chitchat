package main

import (
	"errors"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	sess := Session{
		Uuid: cookie.Value,
	}
	if ok, _ := sess.Check(); !ok {
		err = errors.New("Invalid session")
		return nil, err
	}

	return &sess, nil
}
