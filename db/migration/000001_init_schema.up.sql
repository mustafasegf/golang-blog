CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar,
  "password" varchar,
  "role" varchar,
  "created" timestamp,
  "updated" timestamp
);

CREATE TABLE "blog" (
  "id" SERIAL,
  "title" varchar,
  "content" text,
  "created" timestamp,
  "updated" timestamp,
  "author_id" int,
  PRIMARY KEY ("id", "author_id")
);

CREATE TABLE "user_comment" (
  "id" SERIAL PRIMARY KEY,
  "blog_id" int,
  "user_id" int,
  "coment" text
);

CREATE TABLE "tags" (
  "title" varchar PRIMARY KEY,
  "blog_id" int
);

CREATE TABLE "category" (
  "title" varchar PRIMARY KEY,
  "blog_id" int
);

ALTER TABLE "blog" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "user_comment" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "user_comment" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

