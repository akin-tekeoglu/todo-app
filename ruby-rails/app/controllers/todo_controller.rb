class TodoController < ApplicationController
  before_action :set_user_id
  def set_user_id
    Thread.current[:user_id] = session[:user_id]
  end

  def show_all
    @completed = params['completed']
    @todos = Todo.for_current_user
    @todos = @todos.where(completed: @completed) unless @completed.nil?
  end

  def show_one
    @todo = params.key?('id') ? Todo.for_current_user.find_by(user_id: session[:user_id], id: params[:id]) : Todo.new
  end

  def create_or_update
    @todo = params.key?('id') ? Todo.for_current_user.find_by(id: params[:id], user_id: session[:user_id]) : Todo.new
    @todo.title = params[:title]
    @todo.description = params[:description]
    @todo.event_date = params[:event_date]
    @todo.save
    render 'success'
  end

  def delete
    Todo.for_current_user.destroy(params[:id])
    render 'success'
  end

  def toggle_complete
    Todo.for_current_user.find(params[:id]).toggle_completed
    render 'success'
  end
end
