# nfl-picks-server
Go Backend for NFL Picks web app

## Local developmemt
### Environment file
Copy env.template file and name it .env.local
Fill in missing values

### Database
Spin up local DB
```docker-compose up```

### Auto Migration
In the main.go file it is commented out. Uncomment if you need the migrations
```db.AutoMigrate(&models.Team{}, &models.User{}, &models.Pick{}, &models.Week{}, &models.Matchup{})```

### API
Run app with live reloading
```air```
