class CreateTodos < ActiveRecord::Migration[6.1]
  def change
    create_table :todos do |t|
      t.string :title
      t.text :description
      t.datetime :event_date

      t.timestamps
    end
  end
end
