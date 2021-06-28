class Todo < ApplicationRecord
  def toggle_completed
    self.completed = !completed
    save
  end
end
