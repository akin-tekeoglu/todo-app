class TodoController < ApplicationController
  def show_all
    @completed = params['completed']
    @todos = @completed.nil? ? Todo.all : Todo.where(completed: @completed)
  end

  def show_one
    @todo = params.key?('id') ? Todo.find(params['id']) : Todo.new
  end

  def create_or_update
    @todo = params.key?('id') ? Todo.find(params['id'].to_i) : Todo.new
    @todo.title = params['title']
    @todo.description = params['description']
    @todo.event_date = params['event_date']
    @todo.save
    render 'success'
  end

  def delete
    Todo.destroy(params['id'].to_i)
    render 'success'
  end

  def toggle_complete
    Todo.find(params['id'].to_i).toggle_completed
    render 'success'
  end
end
