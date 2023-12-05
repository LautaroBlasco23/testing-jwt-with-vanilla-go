# API to handle authorization requests.

In this project I want to develop an API that can handle authorization requests using go's http standard library
and jwt-go library. I've also use Docker and Postgres.

# Tech especifications

## Models

User:

* Email
* Password
* Admin

User Input (for register):

* Email
* Password
* AdminSecret

Token: 

* Token

## Routes

User API Routes:

* /login -> POST
* /register -> POST (Unique Email is activated)
* /user?id=# -> GET (You've to put your user's ID instead of the "#")

Admin API Routes: (Only if you have admin's token)

* /data -> GET
* /user?id=# -> GET (As admin you can get any user's data)

## Env variables

Postgres: 

* POSTGRES_HOST
* POSTGRES_USER
* POSTGRES_PASSWORD
* POSTGRES_DBNAME
* POSTGRES_PORT

JWT:

* JWT_SECRET
* JWT_ADMIN

For admin creation:

* ADMIN_PASS

## Scripts

If you want to run this script, you've to set your env variables, Have go and docker installed also.

* go mod tidy -> Download libraries 
* docker compose up -> To run Database
* go run main.go -> To run sv
