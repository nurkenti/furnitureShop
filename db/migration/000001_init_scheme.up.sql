CREATE TYPE "user_role" AS ENUM (
  'customer',
  'admin'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid),
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "age" int NOT NULL,
  "role" user_role DEFAULT (castomer),
  "created_at" timestamp DEFAULT (now()),
  "update_at" timestamp DEFAULT (now())
);

CREATE UNIQUE INDEX ON "users" ("email");
