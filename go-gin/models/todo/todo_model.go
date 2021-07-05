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
	if title == "" {
		return nil, InvalidTitle
	}
	if description == "" {
		return nil, InvalidDescription
	}
	return &Todo{Title: title, Description: description, EventDate: todoEventDate, Completed: false}, Ok
}

func (t *Todo) Save() {
	if t.Id == 0 {
		utils.Db().NamedExec("INSERT INTO todo(title, description, completed, event_date, user_id) VALUES (:title, :description, :completed, :event_date , :user_id)", map[string]interface{}{
			"title":       t.Title,
			"description": t.Description,
			"completed":   t.Completed,
			"event_date":  t.EventDate.Format(time.RFC3339),
			"user_id":     t.UserId,
		})
	} else {
		utils.Db().NamedExec("UPDATE todo SET title=:title, description=:description, event_date=:event_date WHERE id=:id AND user_id=:user_id", map[string]interface{}{
			"title":       t.Title,
			"description": t.Description,
			"event_date":  t.EventDate.Format(time.RFC3339),
			"user_id":     t.UserId,
			"id":          t.Id,
		})
	}
}

func (t *Todo) ToggleCompleted() {
	t.Completed = !t.Completed
	utils.Db().NamedExec("UPDATE todo SET completed=:completed WHERE id=:id", map[string]interface{}{
		"completed": t.Completed,
		"id":        t.Id,
	})
}

func GetAll(userId int) []Todo {
	todos := []Todo{}
	utils.Db().Select(&todos, "SELECT id, title, description, completed, event_date, user_id FROM todo where user_id=?", userId)
	return todos
}

func GetById(id int, userId int) Todo {
	todoItem := Todo{}
	utils.Db().Get(&todoItem, "SELECT id, title, description, completed, event_date, user_id FROM todo WHERE id=? and user_id=?", id)
	return todoItem
}
