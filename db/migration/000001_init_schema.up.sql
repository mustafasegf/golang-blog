CREATE TABLE "users" (
  "id" SERIAL  PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE ,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "blogs" (
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


ALTER TABLE "blogs" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
