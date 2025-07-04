# GoToDo

Simple ToDo API. 

A Todo has a title, description, priority and a completed flag.

## Development setup

- Run `make start-develop` to start a MySQL server in a docker container.
- Edit `configs/dev.yaml` if required.
- Create the DB tables (`database/mysql/schema.sql`).
- Optionally use `database/mysql/dev-data.sql` for some example data/users.
- Run `make develop` to run the server.

Tests can be run with `make test`.

Run `make stop-develop` to shut the MySQL Docker instance down.

## Users

The API requires a valid user to login. Run `make hashpassword` to create the `bin/hashpassword` 
console command for hashing users passwords. Then create an entry in the `users` table.

## TODO

- Better JSON serialisation/group serialisation a la Symfony.
- Better user management.

## Structure

### Model package

Defines an interface for the data we want to persist, doing it this way means we can
implement this interface for each database type we want to support.

### Code

- cmd/gogotod/main: main entry point, set's up signals and creates and starts the Application. 
- cmd/gogotod/application.go: the "application", reads the config, set's up the DB connection/services and creates and starts the HTTP server.
- http/server.go: Implements an HTTP REST server.