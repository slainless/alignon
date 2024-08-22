BEGIN;
ALTER TABLE "transaction_records" DROP COLUMN "catalog_id";
DROP TABLE IF EXISTS "purchase_catalog";
COMMIT;