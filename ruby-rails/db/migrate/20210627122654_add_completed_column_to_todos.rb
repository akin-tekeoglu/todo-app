class AddCompletedColumnToTodos < ActiveRecord::Migration[6.1]
  def change
    add_column :todos, :completed, :boolean, :default => 0
  end
end
