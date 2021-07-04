package todo

import (
	"database/sql"
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
	utils.Execute("INSERT INTO todo(id, title, description, completed, event_date, user_id) VALUES ( ?, ?, ?, ?, ? ,? )", t.Id, t.Title, t.Description, t.Completed, t.EventDate.Format(time.RFC3339), t.UserId)
}

func (t *Todo) ToggleCompleted() {
	t.Completed = !t.Completed
	utils.Execute("UPDATE todo SET completed=? WHERE id=?", t.Completed, t.Id)
}

func GetAll() []Todo {
	return utils.Query("SELECT id, title, description, completed, event_date, user_id FROM todo", Todo{}, processRow).([]Todo)
}

func GetById(id int) Todo {
	return utils.Single("SELECT id, title, description, completed, event_date, user_id FROM todo WHERE id=?", Todo{}, processRow, id).(Todo)
}

func processRow(rows *sql.Rows) interface{} {
	todoItem := Todo{}
	var eventDate string
	rows.Scan(&todoItem.Id, &todoItem.Title, &todoItem.Description, &todoItem.Completed, &eventDate, &todoItem.UserId)
	todoItem.EventDate, _ = time.Parse(time.RFC3339, eventDate)
	return todoItem
}
