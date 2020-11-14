# Treat your plants
Trello failed my stingy ass, so I decided to build myself a service that can
help me keep track of when and how to take care of my plants. 

## Development

### Migrations
We use the migrations lib [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

#### Create new ones
Naming style YYYYMMDDHHII_short_description.[up|down].sql

The content of themigrations should be written as just postgres flavour SQL. 

#### Applying migrations
To find out how to apply the migrations, have a look at this [tutorial](https://github.com/golang-migrate/migrate#cli-usage)
