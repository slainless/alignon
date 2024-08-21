BEGIN;
CREATE TABLE "consumers" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "NIK" varchar(16) UNIQUE NOT NULL,
  "full_name" varchar(256) NOT NULL,
  "legal_name" varchar(256) NOT NULL,
  "birth_place" varchar(256) NOT NULL,
  "birth_date" date NOT NULL,
  "salary" money NOT NULL,
  "ktp_photo" char(36) NOT NULL,
  "selfie_photo" char(36) NOT NULL
);

CREATE TABLE "limits" (
  "consumer_id" uuid PRIMARY KEY,
  "tenor_1" money NOT NULL,
  "tenor_2" money NOT NULL,
  "tenor_3" money NOT NULL,
  "tenor_4" money NOT NULL
);

CREATE TABLE "transaction_records" (
  "contract_id" varchar(255) PRIMARY KEY,
  "loan_id" uuid NOT NULL,
  "otr" money NOT NULL,
  "admin_fee" money NOT NULL,
  "installment" money NOT NULL,
  "interest" money NOT NULL,
  "asset_name" varchar(255) NOT NULL,
  "total" money NOT NULL,
  "status" smallint NOT NULL DEFAULT 0
);

CREATE TABLE "installment_records" (
  "installment_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "contract_id" varchar(255) NOT NULL,
  "paid_at" timestamp NOT NULL
);

CREATE TABLE "loans" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "consumer_id" uuid NOT NULL,
  "amount" money NOT NULL,
  "tenor" smallint NOT NULL,
  "installment_length" smallint NOT NULL,
  "consumer_limit" money NOT NULL,
  "consumer_salary" money NOT NULL,
  "loaned_at" timestamp NOT NULL,
  "status" smallint NOT NULL DEFAULT 0
);

ALTER TABLE "limits" ADD FOREIGN KEY ("consumer_id") REFERENCES "consumers" ("id");
ALTER TABLE "transaction_records" ADD FOREIGN KEY ("loan_id") REFERENCES "loans" ("id");
ALTER TABLE "installment_records" ADD FOREIGN KEY ("contract_id") REFERENCES "transaction_records" ("contract_id");
ALTER TABLE "loans" ADD FOREIGN KEY ("consumer_id") REFERENCES "consumers" ("id");
COMMIT;