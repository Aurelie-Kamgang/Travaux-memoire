### Requirements
* Docker and Go
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate) 

## Our API will expose the following endpoints:

- POST /items to add a new item to the list
- GET /items to fetch all existing items in the list
- GET /items/{itemId} to fetch a single item from the list using its ID
- PUT /items/{itemId} to update an existing item
- DELETE /items/{itemId} to delete an item from the list

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

- Create the password of user postgres: ``sudo -u postgres psql`` and ``ALTER USER postgres WITH PASSWORD 'monmotdepasse';``
- Install the migrate binary by following the installation guide and generate the database migrations by running in super user mode (``sudo su``):
``- $ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
- $ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
- $ apt-get update
- $ apt-get install -y migrate``

In my case i have:

```bash
psql -h localhost -U postgres -W -c "create database example;"  // Create database example with user postgres in localhost
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

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/8f4379fd-c6f9-4e31-b863-7eb03bcae194)


- View a specific item:

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/b29a244a-d70e-4496-806c-47bbc41ac237)

- Add a new item:
 
 ![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/6f031cec-e739-4802-bc4b-9cbfaa1a935a)

- Supprimer un élément dans la base de donnée:

![image](https://github.com/Aurelie-Kamgang/Travaux-memoire/assets/74322332/8d07eff0-d581-4099-874b-51335eb7c4ea)



Référence: https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/
