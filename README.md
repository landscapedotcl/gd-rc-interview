
# gd-rc-interview

Technical test for Golang Developer job.

The project consists on an application wich has a CRUD for users.

## Run Locally

Clone the project.

```bash
  git clone https://github.com/landscapedotcl/gd-rc-interview.git
```

Go to the project directory.

```bash
  cd gd-rc-interview
```

Launch docker container with Postgres image.

```bash
  make docker-container-create
```

Launch the database inside the container.

```bash
  make docker-postgres-createdb
```

Run migrations over the database. (Migrate library must be installed on the computer. Brief tutorial for every OS: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```bash
  make run-migrations-up
```

Run the app.

```bash
  go run ./... 
```

Run the unit tests.

```bash
  go test ./... 
```

  
## API Reference

#### Create a new user

```http
  POST localhost:8000/api/users/create
```

| Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required** - Between 4 & 20 chars.|
| `email` | `string` | **Required** - Must include an '@'.|

Return the created User as a json object.

#### Read all users

```http
  GET localhost:8000/api/users/readall
```

Return all fetched Users as a json object.

#### Filter by name and/or email

```http
  GET localhost:8000/api/users/filter
```

| Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Optional** - Substring of the name.|
| `email` | `string` | **Optional** - Substring of the email.|

Return all Users that match with the provided filter as a json object.
- If no parameters are sent return all Users.
- If only name is send, it return filtering only by name.
- If only email is send, it return filtering only by email.
- If both name and email are send, it return filtering by both of them.


#### Update an existing user

```http
  PUT localhost:8000/api/users/update?{id}
```

| Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required** - Between 4 & 20 chars.|
| `email` | `string` | **Required** - Must include an '@'.|

Return the updated User as a json object. 'id' must be sent through the URL.


#### Delete an existing user

```http
  DELETE localhost:8000/api/users/delete?{id}
```
Return the deleted User as a json object. 'id' must be sent through the URL.

## Tech Stack

**Language:**
- Go

**Packages:** 
- github.com/lib/pq
- github.com/go-chi/chi/v5
- github.com/google/uuid


**Database:**
- PostgreSQL

**Others:**
- Docker
- Makefile

  
## Author

- [@RamiroCuenca](https://www.linkedin.com/in/ramiro-cuenca-salinas-749a2020a/)

  