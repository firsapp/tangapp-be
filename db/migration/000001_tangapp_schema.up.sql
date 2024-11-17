  CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "title" varchar,
  "created_at" timestamp
);

CREATE TABLE "events" (
  "id" int PRIMARY KEY,
  "title" varchar,
  "description" text,
  "status" varchar,
  "total_amount" int,
  "date_event" timestamp,
  "created_by" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "event_purchase_details" (
  "id" int PRIMARY KEY,
  "event_id" int,
  "name" varchar,
  "qty" int,
  "each_price" int,
  "total_price" int,
  "updated_at" timestamp
);

CREATE TABLE "event_member_details" (
  "id" int PRIMARY KEY,
  "event_id" int,
  "user_id" int,
  "bill" int,
  "paid" int,
  "compensation" int,
  "notes" text,
  "done" bool
);

CREATE TABLE "payment_history" (
  "id" int PRIMARY KEY,
  "event_member_details_id" int,
  "to_user_id" int,
  "nominal" int,
  "description" text,
  "created_at" timestamp
);

ALTER TABLE "event_member_details" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "event_purchase_details" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "payment_history" ADD FOREIGN KEY ("event_member_details_id") REFERENCES "event_member_details" ("id");

ALTER TABLE "payment_history" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");

ALTER TABLE "event_member_details" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");
