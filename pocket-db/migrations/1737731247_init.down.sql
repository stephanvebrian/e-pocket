-- DOWN MIGRATION
-- Drop tables in reverse order of creation to respect foreign key dependencies
DROP TABLE IF EXISTS transaction_history;

DROP TABLE IF EXISTS transfer;

DROP TABLE IF EXISTS account;

DROP TABLE IF EXISTS "user";

-- Disable the `pgcrypto` extension (optional, as it is a global extension)
DROP EXTENSION IF EXISTS pgcrypto;