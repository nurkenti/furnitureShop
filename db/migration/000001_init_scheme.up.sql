CREATE TYPE "user_role" AS ENUM (
  'customer',
  'admin'
);

CREATE TYPE "chair_model" AS ENUM (
  'sonyx',
  'kurumi'
);

CREATE TYPE "wardrobe_model" AS ENUM (
  'unibi',
  'facito'
);

CREATE TYPE "chair_material" AS ENUM (
  'wood',
  'metal',
  'fabric'
);

CREATE TYPE "wardrobe_material" AS ENUM (
  'mdf',
  'dsp'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "age" int NOT NULL,
  "role" user_role DEFAULT 'customer',
  "created_at" timestamp DEFAULT now(),
  "update_at" timestamp DEFAULT now()
);

CREATE TABLE "chair" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "model" chair_model NOT NULL DEFAULT 'sonyx',
  "material" chair_material DEFAULT 'wood',
  "price" float,
  "created_at" timestamp DEFAULT now()
);

CREATE TABLE "wardrobe" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "model" wardrobe_model NOT NULL DEFAULT 'unibi',
  "material" wardrobe_material DEFAULT 'mdf',
  "price" float NOT NULL,
  "created_at" timestamp DEFAULT now()
);

CREATE TABLE "warehouse" (
  "product_model" varchar PRIMARY KEY,  -- модель как первичный ключ
  "product_type" varchar NOT NULL,     
  "quantity" int NOT NULL DEFAULT 0,
  "updated_at" timestamp DEFAULT now()
);
-- добавим один раз и все
INSERT INTO "warehouse" ("product_model", "product_type") VALUES
('sonyx', 'chair'),
('kurumi', 'chair'),
('unibi', 'wardrobe'),
('facito', 'wardrobe');

CREATE UNIQUE INDEX ON "users" ("email");
