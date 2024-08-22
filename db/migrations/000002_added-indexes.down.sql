BEGIN;
DROP INDEX IF EXISTS "transaction_records_loan_id_idx";
DROP INDEX IF EXISTS "installment_records_transaction_id_idx";
DROP INDEX IF EXISTS "loans_consumer_id_idx";
COMMIT;