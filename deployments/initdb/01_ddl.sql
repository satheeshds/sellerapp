
CREATE TYPE "products_status" AS ENUM (
  'out_of_stock',
  'in_stock',
  'running_low'
);

CREATE TYPE "payment_method" AS ENUM (
  'cash',
  'credit_card',
  'debit_card'
);

CREATE TYPE "payment_status" AS ENUM (
  'initiated',
  'successful',
  'rejected',
  'refunded'
);

CREATE TYPE "order_status" AS ENUM (
  'created',
  'cancelled',
  'completed'
);

CREATE TABLE "transactions" (
  "id" int PRIMARY KEY,
  "total_amount" decimal,
  "payment_method" payment_method,
  "status" payment_status,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "orders" (
  "id" int PRIMARY KEY,
  "user_id" int,
  "transaction_id" int,
  "status" order_status,
  "created_at" timestamptz DEFAULT 'now()',
  "total_amount" decimal,
  "product_id" int,
  "product_quantity" int
);

CREATE TABLE "products" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "price" decimal,
  "status" products_status,
  "created_at" timestamptz,
  "quantity" int not null check (quantity >= 0)
);

CREATE TABLE "users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "full_name" varchar,
  "created_at" timestamptz
);

COMMENT ON COLUMN "orders"."created_at" IS 'When order created';

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
