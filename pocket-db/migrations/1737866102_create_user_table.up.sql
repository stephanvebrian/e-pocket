-- Enable the `pgcrypto` extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create the `user` table with a `UUID` primary key
CREATE TABLE
  "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (), -- UUID v4 as the primary key
    username VARCHAR(255) NOT NULL UNIQUE, -- Unique username for the user
    created_at TIMESTAMP NOT NULL DEFAULT NOW (), -- Record creation timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT NOW () -- Last update timestamp
  );

-- Add comments for the `user` table and its columns
COMMENT ON TABLE "user" IS 'Table to store user details with unique usernames.';

COMMENT ON COLUMN "user".id IS 'Primary key for the user table, generated as UUID v4.';

COMMENT ON COLUMN "user".username IS 'Unique username for the user, must be provided.';

COMMENT ON COLUMN "user".created_at IS 'Timestamp indicating when the user record was created.';

COMMENT ON COLUMN "user".updated_at IS 'Timestamp indicating the last time the user record was updated.';

-- Add `user_id` column to `account` table
ALTER TABLE account
ADD COLUMN user_id UUID NOT NULL;

-- Add foreign key constraint linking `user_id` to `user.id`
ALTER TABLE account ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES "user" (id);

-- Add comments for the `user_id` column and updated table
COMMENT ON COLUMN account.user_id IS 'References the user owning the account.';

COMMENT ON TABLE account IS 'Table to store account details linked to users.';