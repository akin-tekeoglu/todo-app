class UserController < ApplicationController
  skip_before_action :authorized
  
  def login; end

  def logout
    session[:user_id] = nil
    redirect_to '/'
  end

  def omniauth
    @user = User.from_omniauth(auth)
    @user.save
    session[:user_id] = @user.id
    redirect_to '/'
  end

  private

  def auth
    request.env['omniauth.auth']
  end
end
