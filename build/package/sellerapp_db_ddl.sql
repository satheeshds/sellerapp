CREATE SCHEMA "ecommerce";

CREATE TYPE "ecommerce"."products_status" AS ENUM (
  'out_of_stock',
  'in_stock',
  'running_low'
);

CREATE TYPE "ecommerce"."payment_method" AS ENUM (
  'cash',
  'credit_card',
  'debit_card'
);

CREATE TYPE "ecommerce"."payment_status" AS ENUM (
  'initiated',
  'successful',
  'rejected',
  'refunded'
);

CREATE TYPE "ecommerce"."order_status" AS ENUM (
  'created',
  'cancelled',
  'completed'
);

CREATE TABLE "ecommerce"."transactions" (
  "id" int PRIMARY KEY,
  "payment_method" ecommerce.payment_method,
  "status" ecommerce.payment_status,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "ecommerce"."orders" (
  "id" int PRIMARY KEY,
  "user_id" int,
  "transaction_id" int,
  "status" ecommerce.order_status,
  "created_at" varchar
);

CREATE TABLE "ecommerce"."order_items" (
  "order_id" int,
  "product_id" int,
  "quantity" int DEFAULT 1
);

CREATE TABLE "ecommerce"."products" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "price" int,
  "status" ecommerce.products_status,
  "created_at" datetime DEFAULT (now())
);

CREATE TABLE "ecommerce"."users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "full_name" varchar,
  "created_at" timestamptz
);

COMMENT ON COLUMN "ecommerce"."orders"."created_at" IS 'When order created';

ALTER TABLE "ecommerce"."orders" ADD FOREIGN KEY ("user_id") REFERENCES "ecommerce"."users" ("id");

ALTER TABLE "ecommerce"."orders" ADD FOREIGN KEY ("transaction_id") REFERENCES "ecommerce"."transactions" ("id");

ALTER TABLE "ecommerce"."order_items" ADD FOREIGN KEY ("order_id") REFERENCES "ecommerce"."orders" ("id");

ALTER TABLE "ecommerce"."order_items" ADD FOREIGN KEY ("product_id") REFERENCES "ecommerce"."products" ("id");
