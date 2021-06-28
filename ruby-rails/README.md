# Requirements

- Ruby 3
- Rails 6
- Sqllite 3

# Installation

- `bundle install`
- Create an oauth 2.0 web app in google developer console with `http://localhost:3000/auth/google_oauth2/callback` redirect URI
- `EDITOR=vi rails credentials:edit` and place your credentials like below
```yaml
google:
  client_id: your-client-id
  client_secret: your-client-secret
```
- `rails db:create && rails db:migrate`
- `guard`
- `rails s` (run this command inside a different terminal)