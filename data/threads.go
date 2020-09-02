package data

type mytest struct {
	id   int
	uuid string
	name string
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id,uuid,topic,user_id,created_at from threads" +
		" order by created_at desc;")
	if err != nil {
		return
	}
	threads = make([]Thread, 0)
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	_ = rows.Close()
	return
}
