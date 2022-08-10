# NFL Picks Server
Go API and MySQL DB for NFL Picks web app  

## Local developmemt
### Environment file
Copy env.template file and name it .env.local
Fill in missing values  

### Setup Go
TODO: Update this with steps to get go setup with its dependencies

### Database
Start local DB
```
docker-compose up
```

### Auto Migration
In the main.go file it is commented out. Uncomment if you need the migrations
```
db.AutoMigrate(&models.Team{}, &models.User{}, &models.Pick{}, &models.Week{}, &models.Matchup{})
```

### Create weekMatchup View  
TODO: Use gorm or some migration to automatically handle this  
Run SQL from [weekMatchupView](src/database/weekMatchupView.sql)

### Seed database
In the main.go file it is commented out. Uncomment if you need the data seeding
```
seed.LoadAll(database.Connector)
```

### API
Run app with live reloading
```
air
```
