BEGIN;
ALTER TABLE "limits" DROP CONSTRAINT "limits_consumer_id_fkey";
ALTER TABLE "transaction_records" DROP CONSTRAINT "transaction_records_loan_id_fkey";
ALTER TABLE "installment_records" DROP CONSTRAINT "installment_records_contract_id_fkey";
ALTER TABLE "loans" DROP CONSTRAINT "loans_consumer_id_fkey";

DROP TABLE IF EXISTS "installment_records";
DROP TABLE IF EXISTS "transaction_records";
DROP TABLE IF EXISTS "limits";
DROP TABLE IF EXISTS "loans";
DROP TABLE IF EXISTS "consumers";
COMMIT;