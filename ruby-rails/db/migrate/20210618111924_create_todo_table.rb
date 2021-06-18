class CreateTodoTable < ActiveRecord::Migration[6.1]
  def change
    create_table :todos do |t|
      t.string :title
      t.string :description
      t.date :date

      t.timestamps
    end
  end
end
