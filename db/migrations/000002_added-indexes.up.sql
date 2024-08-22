BEGIN;
CREATE INDEX "transaction_records_loan_id_idx" ON "transaction_records" ("loan_id");
CREATE INDEX "installment_records_transaction_id_idx" ON "installment_records" ("transaction_id");
CREATE INDEX "loans_consumer_id_idx" ON "loans" ("consumer_id");
COMMIT;