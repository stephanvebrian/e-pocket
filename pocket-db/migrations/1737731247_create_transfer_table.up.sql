CREATE TABLE
  transfer (
    id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier
    transaction_id VARCHAR(36) NOT NULL UNIQUE, -- Unique transaction ID
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