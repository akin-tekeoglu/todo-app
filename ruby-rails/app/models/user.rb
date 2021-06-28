class User < ApplicationRecord
  def self.from_omniauth(auth)
    where(user_name: auth.info.email).first_or_initialize do |user|
      user.user_name = auth.info.email
      user.password = SecureRandom.hex
    end
  end

  def todos
    Todo.where(user_id: id)
  end
end
