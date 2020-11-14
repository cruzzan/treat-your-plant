# Treat your plants
Trello failed my stingy ass, so I decided to build myself a service that can
help me keep track of when and how to take care of my plants. 

## Development

### Local dependencies
When we are developing, we will sometimes want to have a local version of the
dependencies for the project for example when testing our migrations and that
kind of stuff. 

For that purpose we have a docker compose with prepared deployments.
```bash
docker-compose -f development-docker-compose.yaml up -d
```

### Migrations
We use the migrations lib [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

#### Create new ones
Naming style YYYYMMDDHHIISS_short_description.[up|down].sql

The content of the migrations should be written as just postgres flavour SQL.

Example for creating a new set of migration files:
```bash
docker run -v $(pwd)/migrations:/migrations migrate/migrate create -ext .sql -dir /migrations name_of_the_migration
```
**Don't forget to fix the permissions on the new migrations files**

#### Applying migrations
To find out how to apply the migrations, have a look at this [tutorial](https://github.com/golang-migrate/migrate#cli-usage)

Example for development use (Provided you are using the development docker
compose)
```bash
docker run -v $(pwd)/migrations:/migrations --network treat-your-plant_dbn migrate/migrate -path=/migrations/ -database postgresql://postgres:postgres@db:5432/treat-your-plant?sslmode=disable up
```
