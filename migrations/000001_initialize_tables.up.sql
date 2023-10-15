CREATE TYPE RoleEnum AS ENUM (
  'admin',
  'contributor',
  'attendee'
);


CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "event" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "location" varchar NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "event_membership" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "event_id" bigint NOT NULL,
  "role"  RoleEnum,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "expense" (
  "id" bigserial PRIMARY KEY,
  "event_id" bigint NOT NULL,
  "item_name" varchar NOT NULL,
  "cost" bigint NOT NULL,
  "description" varchar,
  "category_id" bigint NOT NULL
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "category_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "event_membership" ADD CONSTRAINT "fk_user_id_event_membership" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "event_membership" ADD CONSTRAINT "fk_event_id_event_membership" FOREIGN KEY ("event_id") REFERENCES "event" ("id") ON DELETE CASCADE;

ALTER TABLE "expense" ADD CONSTRAINT "fk_event_id_expense" FOREIGN KEY ("event_id") REFERENCES "event" ("id") ON DELETE CASCADE;

ALTER TABLE "expense" ADD CONSTRAINT "fk_category_id_expense" FOREIGN KEY ("category_id") REFERENCES "category" ("id") ON DELETE CASCADE;