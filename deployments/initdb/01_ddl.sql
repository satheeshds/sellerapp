
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

CREATE TABLE "payments" (
  "id" serial PRIMARY KEY,
  "total_amount" decimal,
  "payment_method" payment_method,
  "status" payment_status,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "user_id" int,
  "transaction_id" int,
  "status" order_status,
  "created_at" timestamptz DEFAULT 'now()',
  "total_amount" decimal,
  "product_id" int,
  "product_quantity" int
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "price" decimal,
  "status" products_status,
  "created_at" timestamptz,
  "quantity" int not null check (quantity >= 0)
);

COMMENT ON COLUMN "orders"."created_at" IS 'When order created';