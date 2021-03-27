# Blog Created by Golang

### Blog
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

## Usage

the server will run on port 3000. To change this, change port variable on app.env
to acces the api, use postman. For sending data use raw body, choose json type.

List of apis:

```POST /api/users``` will register user. Need data {Username, Password, Name}   
```POST /api/users/login``` will login user. Give token for CUD operation. Need data {Username, password}  

```GET /api/blogs``` will get al blogs  
```GET /api/blogs/:id``` get spesific blog using id 
```POST /api/blogs``` create blog. Need data {Title, Content, AuthorID, Token}  


```POST /api/blogs/:id/update``` update blog need data {title, content, token}  
```POST /api/blogs/:id/delete``` delete blog, need data {token}  

```GET /api/blogs/:id/comments``` get all coments in one blog  
```POST /api/blogs/:id/comments``` create comment in one blog. Need data {comment,  token}  

```POST /comments/:id/update``` update comment, need data {comment, data}  
```POST /comments/:id/delete``` delete comment, need data {token}  