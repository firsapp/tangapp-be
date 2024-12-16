CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar,
  "email" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- Events table
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
  "updated_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "is_active" bool NOT NULL DEFAULT true,
  CONSTRAINT "fk_events_created_by" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON DELETE CASCADE
);

-- Event Purchase Details table
CREATE TABLE "event_purchase_details" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "event_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "qty" int NOT NULL,
  "each_price" int NOT NULL,
  "total_price" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp,
  CONSTRAINT "fk_event_purchase_details_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE
);

-- Event Member Details table
CREATE TABLE "event_member_details" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "event_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "bill" int,
  "paid" int,
  "compensation" int,
  "notes" text,
  "done" bool NOT NULL DEFAULT false,
  CONSTRAINT "fk_event_member_details_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_event_member_details_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

-- Payment History table
CREATE TABLE "payment_history" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "event_member_details_id" uuid NOT NULL,
  "from_user_id" uuid NOT NULL,
  "to_user_id" uuid NOT NULL,
  "nominal" int NOT NULL,
  "description" text,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CONSTRAINT "fk_payment_history_event_member_details_id" FOREIGN KEY ("event_member_details_id") REFERENCES "event_member_details" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_payment_history_from_user_id" FOREIGN KEY ("from_user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_payment_history_to_user_id" FOREIGN KEY ("to_user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

-- Indexes for foreign keys (optional, improves performance in large datasets)
CREATE INDEX idx_event_id ON event_member_details (event_id);
CREATE INDEX idx_user_id ON event_member_details (user_id);
CREATE INDEX idx_event_purchase_details_event_id ON event_purchase_details (event_id);
CREATE INDEX idx_event_member_details_id ON payment_history (event_member_details_id);
CREATE INDEX idx_from_user_id ON payment_history (from_user_id);
CREATE INDEX idx_to_user_id ON payment_history (to_user_id);
