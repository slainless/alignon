BEGIN;
CREATE TABLE "purchase_catalog" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "item_id" varchar(255) NOT NULL,
  "provider_id" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "price" bigint NOT NULL
);

ALTER TABLE "purchase_catalog" ADD CONSTRAINT "unique_item_id_provider_id" UNIQUE ("item_id", "provider_id");

ALTER TABLE "transaction_records" ADD COLUMN "catalog_id" uuid NOT NULL;
ALTER TABLE "transaction_records" ADD FOREIGN KEY ("catalog_id") REFERENCES "purchase_catalog" ("id");

CREATE INDEX "idx_transaction_records_catalog_id" ON "transaction_records" ("catalog_id");
COMMIT;