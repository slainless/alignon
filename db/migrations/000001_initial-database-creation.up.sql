BEGIN;
CREATE TABLE "consumers" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "NIK" varchar(16) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "full_name" varchar(255) NOT NULL,
  "legal_name" varchar(255) NOT NULL,
  "birth_place" varchar(255) NOT NULL,
  "birth_date" date NOT NULL,
  "salary" bigint NOT NULL,
  "ktp_photo" varchar(255) NOT NULL,
  "selfie_photo" varchar(255) NOT NULL
);

CREATE TABLE "limits" (
  "consumer_id" uuid PRIMARY KEY,
  "tenor_1" bigint NOT NULL,
  "tenor_2" bigint NOT NULL,
  "tenor_3" bigint NOT NULL,
  "tenor_4" bigint NOT NULL
);

CREATE TABLE "transaction_records" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "contract_id" varchar(255),
  "loan_id" uuid NOT NULL,
  "otr" bigint,
  "admin_fee" bigint,
  "installment" bigint,
  "interest" bigint,
  "asset_name" varchar(255) NOT NULL,
  "amount" bigint NOT NULL,
  "status" smallint NOT NULL DEFAULT 0
);

CREATE TABLE "installment_records" (
  "installment_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "transaction_id" uuid NOT NULL,
  "paid_at" timestamp NOT NULL
);

CREATE TABLE "loans" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "consumer_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "tenor" smallint NOT NULL,
  "installment_length" smallint NOT NULL,
  "consumer_limit" bigint NOT NULL,
  "consumer_salary" bigint NOT NULL,
  "loaned_at" timestamp,
  "status" smallint NOT NULL DEFAULT 0
);

ALTER TABLE "limits" ADD FOREIGN KEY ("consumer_id") REFERENCES "consumers" ("id");
ALTER TABLE "transaction_records" ADD FOREIGN KEY ("loan_id") REFERENCES "loans" ("id");
ALTER TABLE "installment_records" ADD FOREIGN KEY ("transaction_id") REFERENCES "transaction_records" ("id");
ALTER TABLE "loans" ADD FOREIGN KEY ("consumer_id") REFERENCES "consumers" ("id");
COMMIT;