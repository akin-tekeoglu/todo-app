class TodoController < ApplicationController
  def show_all
    @completed = params[:completed]
    @todos = @current_user.todos
    @todos = @todos.where(completed: @completed) unless @completed.nil?
    @page = params[:page].to_i
    @count = @todos.count
    @todos = @todos.order('event_date ASC').limit(6).offset((@page - 1)*6)
  end

  def show_one
    @todo = params.key?('id') ? @current_user.todos.find(params[:id]) : Todo.new
  end

  def create_or_update
    @todo = params.key?('id') ? @current_user.todos.find(params[:id]) : Todo.new
    @todo.title = params[:title]
    @todo.description = params[:description]
    @todo.event_date = params[:event_date]
    @todo.user_id = @current_user.id
    @todo.save
    render 'success'
  end

  def delete
    @current_user.todos.destroy(params[:id])
    render 'success'
  end

  def toggle_complete
    @current_user.todos.find(params[:id]).toggle_completed
    render 'success'
  end
end
