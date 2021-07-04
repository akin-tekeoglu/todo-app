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

const (
	Ok                 = iota
	InvalidTitle       = iota
	InvalidDescription = iota
	InvalidEventDate   = iota
)

func New(title string, description string, eventDate string) (*Todo, uint) {
	todoEventDate, err := time.Parse(time.RFC3339, eventDate)
	if err != nil {
		return nil, InvalidEventDate
	}
	if title != "" {
		return nil, InvalidTitle
	}
	if description != "" {
		return nil, InvalidDescription
	}
	return &Todo{Title: title, Description: description, EventDate: todoEventDate, Completed: false}, Ok
}

func (t *Todo) Save() {
	if t.Id == 0 {
		utils.Execute("INSERT INTO todo(id, title, description, completed, event_date, user_id) VALUES ( ?, ?, ?, ?, ? ,? )", t.Id, t.Title, t.Description, t.Completed, t.EventDate.Format(time.RFC3339), t.UserId)
	} else {
		utils.Execute("UPDATE todo SET title=?, description=?, event_date=? WHERE id=?", t.Title, t.Description, t.EventDate.Format(time.RFC3339), t.Id)
	}
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
	err := rows.Scan(&todoItem.Id, &todoItem.Title, &todoItem.Description, &todoItem.Completed, &eventDate, &todoItem.UserId)
	if err != nil {
		panic(err)
	}
	todoItem.EventDate, err = time.Parse(time.RFC3339, eventDate)
	if err != nil {
		panic(err)
	}
	return todoItem
}
