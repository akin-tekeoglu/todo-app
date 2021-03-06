class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception
  before_action :authorized

  def authorized
    redirect_to '/login' if session[:user_id].nil?
  end

  def current_user
    return unless session[:user_id]

    @current_user ||= User.find(session[:user_id])
  end
end
