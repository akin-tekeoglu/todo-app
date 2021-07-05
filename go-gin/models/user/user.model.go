package user

import "todo-app/utils/db"

type User struct {
	Id       int
	Username string
	Password string
}

func GetById(id int) User {
	user := User{}
	db.Db().Get(&user, "SELECT id, username FROM user WHERE id=?", id)
	return user
}
