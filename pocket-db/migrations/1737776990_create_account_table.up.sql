CREATE TABLE
  account (
    id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier
    account_number VARCHAR(20) NOT NULL UNIQUE, -- Unique account identifier
    account_name VARCHAR(255) NOT NULL, -- Account holder's name
    balance BIGINT NOT NULL DEFAULT 0, -- Current account balance
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE', -- Account status (e.g., ACTIVE, INACTIVE)
    created_at TIMESTAMP NOT NULL DEFAULT NOW (), -- Record creation timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT NOW () -- Last update timestamp
  );

CREATE TABLE
  transaction_history (
    id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier
    account_id BIGINT NOT NULL, -- Foreign key reference to the account table
    transaction_type VARCHAR(50) NOT NULL, -- Type of transaction (e.g., INTRABANK, PAYMENT)
    transaction_amount BIGINT NOT NULL DEFAULT 0, -- Transaction amount
    ending_balance BIGINT NOT NULL DEFAULT 0, -- Ending balance after the transaction
    status VARCHAR(20) NOT NULL DEFAULT 'PROCESSING', -- Transaction status (e.g., SUCCESS, PROCESSING, FAILED)
    created_at TIMESTAMP NOT NULL DEFAULT NOW (), -- Record creation timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT NOW () -- Last update timestamp
  );