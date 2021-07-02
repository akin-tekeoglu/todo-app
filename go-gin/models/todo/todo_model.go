package todo

import (
	"time"
	"todo-app/utils"
)

type Todo struct {
	Id          int
	Title       string
	Description string
	Completed   bool
	EventDate   time.Time
	UserId      int
}

func (t *Todo) Save() {
	t.Completed = false
	utils.Execute("INSERT INTO todo(id, title, description, completed, event_date, user_id) VALUES ( ?, ?, ?, ?, ? ,? )", t.Id, t.Title, t.Description, t.Completed, t.EventDate, t.UserId)
}

func (t *Todo) ToggleCompleted() {
	t.Completed = !t.Completed
	utils.Execute("UPDATE todo SET completed=? WHERE id=?", t.Completed, t.Id)
}

func GetById(id int) *Todo {
	return nil
}

func GetAll(id int) *[]Todo {
	return nil
}
