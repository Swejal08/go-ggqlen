CREATE TYPE RoleEnum AS ENUM (
  'admin',
  'contributor',
  'attendee'
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "event" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "location" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "event_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "start_date" timestamp NOT NULL,
  "end_date" timestamp NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "event_membership" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "event_id" uuid NOT NULL,
  "role" RoleEnum,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "expense" (
  "id" uuid PRIMARY KEY,
  "event_id" uuid NOT NULL,
  "item_name" varchar NOT NULL,
  "cost" bigint NOT NULL,
  "description" varchar,
  "category_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "category" (
  "id" uuid PRIMARY KEY,
  "category_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "event_membership" ADD CONSTRAINT "fk_user_id_event_membership" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "event_membership" ADD CONSTRAINT "fk_event_id_event_membership" FOREIGN KEY ("event_id") REFERENCES "event" ("id") ON DELETE CASCADE;

ALTER TABLE "sessions" ADD CONSTRAINT "fk_event_id_sessions" FOREIGN KEY ("event_id") REFERENCES "event" ("id") ON DELETE CASCADE;

ALTER TABLE "expense" ADD CONSTRAINT "fk_event_id_expense" FOREIGN KEY ("event_id") REFERENCES "event" ("id") ON DELETE CASCADE;

ALTER TABLE "expense" ADD CONSTRAINT "fk_category_id_expense" FOREIGN KEY ("category_id") REFERENCES "category" ("id") ON DELETE CASCADE;