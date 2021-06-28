# Requirements

- Ruby 3
- Rails 6
- Sqllite 3

# Installation

- `bundle install`
- `EDITOR=vi rails credentials:edit` and place your credentials like below
```yaml
google:
  client_id: your-client-id
  client_secret: your-client-secret
```
- `rails db:create && rails db:migrate`
- `guard`
- `rails s` (run this command inside a different terminal)