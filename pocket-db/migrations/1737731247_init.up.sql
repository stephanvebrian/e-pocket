-- UP MIGRATION
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

-- Create the `account` table
CREATE TABLE
  account (
    id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier
    account_number VARCHAR(20) NOT NULL UNIQUE, -- Unique account identifier
    prefix VARCHAR(14) NOT NULL, -- Account number prefix
    suffix VARCHAR(4) NOT NULL, -- Account number suffix
    pocket_number INT NOT NULL, -- Pocket number
    account_name VARCHAR(255) NOT NULL, -- Account holder's name
    balance BIGINT NOT NULL DEFAULT 0, -- Current account balance
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE', -- Account status (e.g., ACTIVE, INACTIVE)
    created_at TIMESTAMP NOT NULL DEFAULT NOW (), -- Record creation timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT NOW (), -- Last update timestamp
    user_id UUID NOT NULL REFERENCES "user" (id) -- Foreign key reference to the user table
  );

-- Create the `transfer` table
CREATE TABLE
  transfer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (), -- UUID v4 as the primary key
    user_id UUID NOT NULL REFERENCES "user" (id), -- Foreign key reference to the user table
    reference_id VARCHAR(36) NOT NULL UNIQUE, -- Unique reference ID for the transfer
    sender_account VARCHAR(20) NOT NULL, -- Sender's account
    receiver_account VARCHAR(20) NOT NULL, -- Receiver's account
    amount BIGINT NOT NULL DEFAULT 0, -- Transfer amount
    sender JSONB DEFAULT NULL, -- Sender's information in JSON format
    receiver JSONB DEFAULT NULL, -- Receiver's information in JSON format
    status VARCHAR(20) NOT NULL DEFAULT 'PROCESSING', -- Current transfer status
    created_at TIMESTAMP NOT NULL DEFAULT NOW (), -- Record creation timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT NOW () -- Last update timestamp
  );

-- Create the `transaction_history` table
CREATE TABLE
  transaction_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (), -- UUID v4 as the primary key
    user_id UUID NOT NULL REFERENCES "user" (id), -- Foreign key reference to the user table
    account_id BIGINT NOT NULL REFERENCES account (id), -- Foreign key reference to the account table
    transaction_type VARCHAR(50) NOT NULL, -- Type of transaction (e.g., INTRABANK, PAYMENT)
    transaction_amount BIGINT NOT NULL DEFAULT 0, -- Transaction amount
    ending_balance BIGINT NOT NULL DEFAULT 0, -- Ending balance after the transaction
    status VARCHAR(20) NOT NULL DEFAULT 'PROCESSING', -- Transaction status (e.g., SUCCESS, PROCESSING, FAILED)
    created_at TIMESTAMP NOT NULL DEFAULT NOW (), -- Record creation timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT NOW () -- Last update timestamp
  );