-- Remove the comments for the `account` table and `user_id` column
COMMENT ON COLUMN account.user_id IS NULL;

COMMENT ON TABLE account IS NULL;

-- Drop the foreign key constraint linking `user_id` to `user.id`
ALTER TABLE account
DROP CONSTRAINT fk_user_id;

-- Remove the `user_id` column from the `account` table
ALTER TABLE account
DROP COLUMN user_id;

-- Remove the comments for the `user` table and its columns
COMMENT ON COLUMN "user".updated_at IS NULL;

COMMENT ON COLUMN "user".created_at IS NULL;

COMMENT ON COLUMN "user".username IS NULL;

COMMENT ON COLUMN "user".id IS NULL;

COMMENT ON TABLE "user" IS NULL;

-- Drop the `user` table
DROP TABLE IF EXISTS "user";

-- Drop the `pgcrypto` extension
DROP EXTENSION IF EXISTS pgcrypto;