# commands

- to get into pgsql shell
  - psql 
- run as a different user
   psql -U postgres -d docker_test

- exit pgsql
  - \q
- to list all the databases
  - \l 
- Create a new user (if needed) and set a password:
  - CREATE USER your_user WITH PASSWORD 'your_password';
- create database
  - CREATE DATABASE docker_test;
- Grant privileges to the user
  - GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_user;
- to check tables
  - \dt
- postgres url
   ```bash
  postgres://<username>:<password>@<host>:<port>/<database>?sslmode=<mode>
    ```
  PG_USER=postgres
  PG_PASSWORD=postgres
  PG_HOST=localhost
  PG_PORT=5432
  PG_DB=void
  PG_SSLMODE=disable
  - PG_SSLMODE=disable is an environment variable (or query parameter) that tells PostgreSQL not to use SSL/TLS encryption when connecting to the database.
