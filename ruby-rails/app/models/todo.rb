class Todo < ApplicationRecord
  scope :for_current_user, -> { where(user_id: Thread.current[:user_id]) }
  before_save -> { self.user_id = Thread.current[:user_id] }

  def toggle_completed
    self.completed = !completed
    save
  end
end
