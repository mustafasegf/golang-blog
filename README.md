# Blog Created by Golang

### Tech
This project created using gin and sqlc to generate type save code from sql queries

## How to use
on make file there's couple command that you can use

this project use prosgres on port 5432.
You can use makefile to install docker

```make pullpostgres```

to make docker image 
```make postgres```

to remove container
```make removecontainer```

to create the db
```make createdb```

to fill the db with table
```make migrateup```

to run the code in windows
```make buiid-win```

to run in linux
```make build-nix```