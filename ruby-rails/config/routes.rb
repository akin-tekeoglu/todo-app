Rails.application.routes.draw do
  root to: redirect('/todo?completed=false&page=1')
  get '/todo', to: 'todo#show_all'
  post '/todo/toggle/:id', to: 'todo#toggle_complete'
  get '/todo/form(/:id)', to: 'todo#show_one'
  post '/todo/form(/:id)', to: 'todo#create_or_update'
  post '/todo/delete/:id', to: 'todo#delete'
  get '/auth/:provider/callback' => 'user#omniauth'
  get '/login' => 'user#login'
  post '/logout' => 'user#logout'
end
