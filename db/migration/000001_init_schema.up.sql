CREATE TABLE "users" (
  "id" bigint PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE ,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "blog" (
  "id" bigint PRIMARY KEY,
  "title" varchar NOT NULL,
  "content" text NOT NULL,
  "author_id" int NOT NULL
);

CREATE TABLE "user_comment" (
  "id" bigint PRIMARY KEY,
  "blog_id" int NOT NULL,
  "user_id" int NOT NULL,
  "coment" text NOT NULL
);

CREATE TABLE "tags" (
  "title" varchar PRIMARY KEY,
  "blog_id" int NOT NULL
);

CREATE TABLE "category" (
  "title" varchar PRIMARY KEY,
  "blog_id" int NOT NULL
);

ALTER TABLE "blog" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "user_comment" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "user_comment" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

