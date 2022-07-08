# nfl-picks-server
Go Backend for NFL Picks web app

## Local developmemt
### Database
```docker-compose up```

### API
go run main.go


## Auto Migration
In the database/main.go file 
```db.AutoMigrate(&models.Team{}, &models.User{}, &models.Pick{}, &models.Week{}, &models.Matchup{})```
