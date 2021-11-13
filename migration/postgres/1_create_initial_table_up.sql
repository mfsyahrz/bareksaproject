CREATE TYPE "news_status" AS ENUM (
  'draft',
  'deleted',
  'publish'
);

CREATE TABLE IF NOT EXISTS "topic" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(255)
);

CREATE TABLE IF NOT EXISTS "news" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" varchar(255),
  "topic_id" int8 NOT NULL,
  "status" news_status NOT NULL default 'draft',
  "created_date" timestamptz default CURRENT_TIMESTAMP,
  "updated_date" timestamptz default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "tag" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(255),
  "created_date" timestamptz default CURRENT_TIMESTAMP,
  "updated_date" timestamptz default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "news_tag" (
  "id" BIGSERIAL PRIMARY KEY,
  "tag_id" int8 NOT NULL,
  "news_id" int8 NOT NULL
);

CREATE UNIQUE INDEX unique_topic_name ON "topic" ("name");
CREATE UNIQUE INDEX unique_news_title ON "news" ("title");
CREATE UNIQUE INDEX unique_tag_name ON "tag" ("name");
CREATE UNIQUE INDEX unique_news_tags ON "news_tag" ("tag_id", "news_id");

ALTER TABLE "news" ADD FOREIGN KEY ("topic_id") REFERENCES "topic" ("id");
ALTER TABLE "news_tag" ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id") ON DELETE CASCADE;
ALTER TABLE "news_tag" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id") ON DELETE CASCADE;

INSERT INTO "topic"
("name")
VALUES
('investment'),
('economy');