BEGIN;
DROP INDEX IF EXISTS "transaction_records_loan_id_idx";
DROP INDEX IF EXISTS "installment_records_contract_id_idx";
DROP INDEX IF EXISTS "loans_consumer_id_idx";
COMMIT;