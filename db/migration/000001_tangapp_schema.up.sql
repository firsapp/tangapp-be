CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar,
  "email" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "events" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "title" varchar,
  "description" text,
  "status" varchar NOT NULL,
  "total_amount" int NOT NULL,
  "date_event" timestamp,
  "created_by" uuid NOT NULL,
  "can_edit" bool NOT NULL DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "event_purchase_details" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "event_id" uuid NOT NULL,
  "name" varchar,
  "qty" int,
  "each_price" int,
  "total_price" int,
  "updated_at" timestamp
);

CREATE TABLE "event_member_details" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "event_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "bill" int,
  "paid" int,
  "compensation" int,
  "notes" text,
  "done" bool NOT NULL DEFAULT false
);

CREATE TABLE "payment_history" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "event_member_details_id" uuid NOT NULL,
  "to_user_id" uuid NOT NULL,
  "nominal" int NOT NULL,
  "description" text,
  "created_at" timestamp NOT NULL
);

ALTER TABLE "event_member_details" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "event_purchase_details" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "payment_history" ADD FOREIGN KEY ("event_member_details_id") REFERENCES "event_member_details" ("id");

ALTER TABLE "payment_history" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");

ALTER TABLE "event_member_details" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");
