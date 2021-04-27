


# Golang-blog project
Simple golang blog for backend practice. This project created using gin and sqlc to generate type save code from sql queries.
  

## Informations

### Version

1.0

### Contact

  

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  blog

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /api/api/blogs/{id} | [delete API blogs ID](#delete-api-blogs-id) | Delete a blog |
| GET | /api/api/blogs | [get API blogs](#get-api-blogs) | Show all blog |
| GET | /api/api/blogs/{id} | [get API blogs ID](#get-api-blogs-id) | Show a blog |
| PATCH | /api/api/blogs/{id} | [patch API blogs ID](#patch-api-blogs-id) | Update a blog |
| POST | /api/api/blogs/{id} | [post API blogs ID](#post-api-blogs-id) | Create a blog |
  


###  comment

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /api/api/comments/{id} | [delete API comments ID](#delete-api-comments-id) | delete comment |
| GET | /api/api/comments/{id} | [get API comments ID](#get-api-comments-id) | get all comment in blog |
| PATCH | /api/api/comments/{id} | [patch API comments ID](#patch-api-comments-id) | update comment |
| POST | /api/api/blogs/{id}/comments | [post API blogs ID comments](#post-api-blogs-id-comments) | Create a comment |
  


###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/api/users | [post API users](#post-api-users) | Create a user |
| POST | /api/api/users/login | [post API users login](#post-api-users-login) | login a user |
  


## Paths

### <span id="delete-api-blogs-id"></span> Delete a blog (*DeleteAPIBlogsID*)

```
DELETE /api/api/blogs/{id}
```

Delete blog by id

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-api-blogs-id-200) | OK | Blog deleted |  | [schema](#delete-api-blogs-id-200-schema) |

#### Responses


##### <span id="delete-api-blogs-id-200"></span> 200 - Blog deleted
Status: OK

###### <span id="delete-api-blogs-id-200-schema"></span> Schema
   
  



### <span id="delete-api-comments-id"></span> delete comment (*DeleteAPICommentsID*)

```
DELETE /api/api/comments/{id}
```

delete comment by comment id

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer Token |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-api-comments-id-200) | OK | OK |  | [schema](#delete-api-comments-id-200-schema) |

#### Responses


##### <span id="delete-api-comments-id-200"></span> 200 - OK
Status: OK

###### <span id="delete-api-comments-id-200-schema"></span> Schema
   
  

[DbCreateCommentRow](#db-create-comment-row)

### <span id="get-api-blogs"></span> Show all blog (*GetAPIBlogs*)

```
GET /api/api/blogs
```

get all blog in json

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-api-blogs-200) | OK | OK |  | [schema](#get-api-blogs-200-schema) |

#### Responses


##### <span id="get-api-blogs-200"></span> 200 - OK
Status: OK

###### <span id="get-api-blogs-200-schema"></span> Schema
   
  

[][DbListBlogRow](#db-list-blog-row)

### <span id="get-api-blogs-id"></span> Show a blog (*GetAPIBlogsID*)

```
GET /api/api/blogs/{id}
```

get blog by id

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-api-blogs-id-200) | OK | OK |  | [schema](#get-api-blogs-id-200-schema) |

#### Responses


##### <span id="get-api-blogs-id-200"></span> 200 - OK
Status: OK

###### <span id="get-api-blogs-id-200-schema"></span> Schema
   
  

[DbGetBlogRow](#db-get-blog-row)

### <span id="get-api-comments-id"></span> get all comment in blog (*GetAPICommentsID*)

```
GET /api/api/comments/{id}
```

get all comment by blog id

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer Token |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-api-comments-id-200) | OK | OK |  | [schema](#get-api-comments-id-200-schema) |

#### Responses


##### <span id="get-api-comments-id-200"></span> 200 - OK
Status: OK

###### <span id="get-api-comments-id-200-schema"></span> Schema
   
  

[DbCreateCommentRow](#db-create-comment-row)

### <span id="patch-api-blogs-id"></span> Update a blog (*PatchAPIBlogsID*)

```
PATCH /api/api/blogs/{id}
```

Update blog title and content

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |
| titleContent | `body` | [APICreateBlogRequest](#api-create-blog-request) | `models.APICreateBlogRequest` | | ✓ | | Blog request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-api-blogs-id-200) | OK | Blog Updated |  | [schema](#patch-api-blogs-id-200-schema) |

#### Responses


##### <span id="patch-api-blogs-id-200"></span> 200 - Blog Updated
Status: OK

###### <span id="patch-api-blogs-id-200-schema"></span> Schema
   
  



### <span id="patch-api-comments-id"></span> update comment (*PatchAPICommentsID*)

```
PATCH /api/api/comments/{id}
```

update comment by comment id

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer Token |
| content | `body` | string | `string` | | ✓ | | Comment content |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-api-comments-id-200) | OK | OK |  | [schema](#patch-api-comments-id-200-schema) |

#### Responses


##### <span id="patch-api-comments-id-200"></span> 200 - OK
Status: OK

###### <span id="patch-api-comments-id-200-schema"></span> Schema
   
  

[DbCreateCommentRow](#db-create-comment-row)

### <span id="post-api-blogs-id"></span> Create a blog (*PostAPIBlogsID*)

```
POST /api/api/blogs/{id}
```

Create blog by title and content

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |
| titleContent | `body` | [APICreateBlogRequest](#api-create-blog-request) | `models.APICreateBlogRequest` | | ✓ | | Blog request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-api-blogs-id-200) | OK | OK |  | [schema](#post-api-blogs-id-200-schema) |

#### Responses


##### <span id="post-api-blogs-id-200"></span> 200 - OK
Status: OK

###### <span id="post-api-blogs-id-200-schema"></span> Schema
   
  

[DbCreateBlogRow](#db-create-blog-row)

### <span id="post-api-blogs-id-comments"></span> Create a comment (*PostAPIBlogsIDComments*)

```
POST /api/api/blogs/{id}/comments
```

Create comment by id and comment

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | Blog ID |
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer Token |
| content | `body` | string | `string` | | ✓ | | Comment content |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-api-blogs-id-comments-200) | OK | OK |  | [schema](#post-api-blogs-id-comments-200-schema) |

#### Responses


##### <span id="post-api-blogs-id-comments-200"></span> 200 - OK
Status: OK

###### <span id="post-api-blogs-id-comments-200-schema"></span> Schema
   
  

[DbCreateCommentRow](#db-create-comment-row)

### <span id="post-api-users"></span> Create a user (*PostAPIUsers*)

```
POST /api/api/users
```

Create user by username, password and name

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| register | `body` | [APICreateUserRequest](#api-create-user-request) | `models.APICreateUserRequest` | | ✓ | | register request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-api-users-200) | OK | OK |  | [schema](#post-api-users-200-schema) |

#### Responses


##### <span id="post-api-users-200"></span> 200 - OK
Status: OK

###### <span id="post-api-users-200-schema"></span> Schema
   
  

[APIUserResponse](#api-user-response)

### <span id="post-api-users-login"></span> login a user (*PostAPIUsersLogin*)

```
POST /api/api/users/login
```

login user by username and password

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| login | `body` | [APILoginUserRequest](#api-login-user-request) | `models.APILoginUserRequest` | | ✓ | | login request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-api-users-login-200) | OK | OK |  | [schema](#post-api-users-login-200-schema) |

#### Responses


##### <span id="post-api-users-login-200"></span> 200 - OK
Status: OK

###### <span id="post-api-users-login-200-schema"></span> Schema
   
  

[APILoginUserResponse](#api-login-user-response)

## Models

### <span id="api-create-blog-request"></span> api.createBlogRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| content | string| `string` | ✓ | |  |  |
| title | string| `string` | ✓ | |  |  |



### <span id="api-create-user-request"></span> api.createUserRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



### <span id="api-login-user-request"></span> api.loginUserRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| password | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



### <span id="api-login-user-response"></span> api.loginUserResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access_token | string| `string` |  | |  |  |
| user | [APIUserResponse](#api-user-response)| `APIUserResponse` |  | |  |  |



### <span id="api-user-response"></span> api.userResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  |  |
| username | string| `string` |  | |  |  |



### <span id="db-create-blog-row"></span> db.CreateBlogRow


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| author_id | integer| `int64` |  | |  |  |
| content | string| `string` |  | |  |  |
| id | integer| `int64` |  | |  |  |
| name | [interface{}](#interface)| `interface{}` |  | |  |  |
| title | string| `string` |  | |  |  |



### <span id="db-create-comment-row"></span> db.CreateCommentRow


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| blog_id | integer| `int64` |  | |  |  |
| comment | string| `string` |  | |  |  |
| id | integer| `int64` |  | |  |  |
| name | [interface{}](#interface)| `interface{}` |  | |  |  |
| user_id | integer| `int64` |  | |  |  |



### <span id="db-get-blog-row"></span> db.GetBlogRow


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| content | string| `string` |  | |  |  |
| id | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| title | string| `string` |  | |  |  |
| userid | integer| `int64` |  | |  |  |



### <span id="db-list-blog-row"></span> db.ListBlogRow


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| content | string| `string` |  | |  |  |
| id | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| title | string| `string` |  | |  |  |
| userid | integer| `int64` |  | |  |  |


