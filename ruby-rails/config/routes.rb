Rails.application.routes.draw do
  get '/todo', to: 'todo#show_all'
  get '/todo/form(/:id)', to: 'todo#show_one'
  post '/todo/form(/:id)', to: 'todo#create_or_update' 
  post '/todo/:id/delete', to: 'todo#delete' 
end
