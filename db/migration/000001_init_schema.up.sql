CREATE TABLE "users" (
  "id" SERIAL  PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE ,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "blog" (
  "id" SERIAL  PRIMARY KEY,
  "title" varchar NOT NULL,
  "content" text NOT NULL,
  "author_id" int NOT NULL
);

CREATE TABLE "comments" (
  "id" SERIAL  PRIMARY KEY,
  "blog_id" int NOT NULL,
  "user_id" int NOT NULL,
  "comment" text NOT NULL
);

CREATE TABLE "tags" (
  "title" varchar PRIMARY KEY,
  "blog_id" int NOT NULL
);

CREATE TABLE "category" (
  "title" varchar PRIMARY KEY,
  "blog_id" int NOT NULL
);

ALTER TABLE "blog" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "tags" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id") ON DELETE CASCADE;

ALTER TABLE "category" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id") ON DELETE CASCADE;

