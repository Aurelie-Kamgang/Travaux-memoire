### Requirements
* Docker and Go
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate) 

## Our API will expose the following endpoints:

POST /items to add a new item to the list
GET /items to fetch all existing items in the list
GET /items/{itemId} to fetch a single item from the list using its ID
PUT /items/{itemId} to update an existing item
DELETE /items/{itemId} to delete an item from the list

### Usage
Clone the repository with:
```bash
git clone gitlab.com/idoko/bucketeer
```

Copy the `env.example` file to a `.env` file.
```bash
$ cp .env.example .env
```
Update the postgres variables declared in the new `.env` to match your preference. 
There's a handy guide on the [Postgres' DockerHub](https://hub.docker.com/_/postgres).

Build and start the services with:
```bash
$ docker-compose up --build
```
The database migration files are in `db/migrations` so feel free to simply source them directly. Alternatively, you can apply them using `migrate` by running:
```bash
$ export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASS@localhost:5432/$PG_DB?sslmode=disable"
$ migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

_**NOTE:** Remember to replace the `$PG*` variables with their actual values_
In my case i have:
```bash
psql -h localhost -U postgres -w -c "create database example;"  // Create database example with user postgres in localhost
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

### Development
After making your changes, you can rebuild the `server` service by running the commands below
```bash
$ docker-compose stop server
$ docker-compose build server
$ docker-compose up --no-start server
$ docker-compose start server
```


 Results:
 
 - Homepage:

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/8f4dbe5e-b34d-4a86-9afb-31ff9a9acadf)

- Fetch all items:

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/1cb7c6a6-ffab-401d-b913-501c7115f5d1)

- View a specific item:

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/b29a244a-d70e-4496-806c-47bbc41ac237)

- Add a new item:
 
 ![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/6f031cec-e739-4802-bc4b-9cbfaa1a935a)

- Supprimer un élément dans la base de donnée:

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/8d07eff0-d581-4099-874b-51335eb7c4ea)



