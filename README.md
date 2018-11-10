# School Notes

## Setup

- run db
```
  pg_ctl -D /usr/local/var/postgres start
```

- create test db
```
  createdb school_note_test
  psql school_note_test
  CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(256),
    text VARCHAR,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
  );
```

- create main db
```
  createdb school_note
  psql school_note
  CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(256),
    text VARCHAR,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
  );
```

## Running
```
  go run main.go note_handlers.go store.go
```

### Source
```
  https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
  https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/
  https://getbootstrap.com/
```
