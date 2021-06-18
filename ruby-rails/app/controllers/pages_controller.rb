class PagesController < ApplicationController
    def index
      @todo=Todo.new
      @todo.title="my title"
      @todo.description="my description"
      @todo.date=Time.now.utc
    end
  end