class TodoController < ApplicationController
    def show_all
      @todos=Todo.all
    end

    def show_one
      
    end

    def create_or_update
      @todo=Todo.new title: params["title"], description: params["description"], event_date: params["event_date"]
      @todo.save
      redirect_to "/todo"
    end

    def delete
      
    end
  end