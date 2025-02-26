-- Seed the `user` table
INSERT INTO
  "user" (username, created_at, updated_at)
VALUES
  ('john_doe', NOW (), NOW ()),
  ('jane_smith', NOW (), NOW ()),
  ('alice_wonderland', NOW (), NOW ());

-- Seed the `account` table
INSERT INTO
  account (
    account_number,
    prefix,
    suffix,
    pocket_number,
    account_name,
    balance,
    status,
    created_at,
    updated_at,
    user_id
  )
VALUES
  (
    '230915134530456701',
    '230915134530',
    '4567',
    1,
    'Tabungan Utama',
    100000,
    'ACTIVE',
    NOW (),
    NOW (),
    (
      SELECT
        id
      FROM
        "user"
      WHERE
        username = 'john_doe'
    )
  ),
  (
    '230915134530456702',
    '230915134530',
    '4567',
    2,
    'Dompet Recehan',
    0,
    'ACTIVE',
    NOW (),
    NOW (),
    (
      SELECT
        id
      FROM
        "user"
      WHERE
        username = 'john_doe'
    )
  ),
  (
    '250215134530111101',
    '250215134530',
    '1111',
    1,
    'Personal Wallet',
    50000,
    'ACTIVE',
    NOW (),
    NOW (),
    (
      SELECT
        id
      FROM
        "user"
      WHERE
        username = 'jane_smith'
    )
  );

-- Seed the `transfer` table
INSERT INTO
  transfer (
    user_id,
    reference_id,
    sender_account,
    receiver_account,
    amount,
    sender,
    receiver,
    status,
    created_at,
    updated_at
  )
VALUES
  (
    (
      SELECT
        id
      FROM
        "user"
      WHERE
        username = 'john_doe'
    ), -- user_id for John Doe
    'ff389d7e-af4c-4bb8-af8b-3179b2839f90', -- reference_id
    (
      SELECT
        account_number
      FROM
        account
      WHERE
        user_id = (
          SELECT
            id
          FROM
            "user"
          WHERE
            username = 'john_doe'
          LIMIT
            1
        )
        AND pocket_number = 1
    ), -- sender_account
    (
      SELECT
        account_number
      FROM
        account
      WHERE
        user_id = (
          SELECT
            id
          FROM
            "user"
          WHERE
            username = 'jane_smith'
          LIMIT
            1
        )
        AND pocket_number = 1
    ), -- receiver_account
    1000000, -- amount
    '{"Name": "John Doe"}', -- sender (JSONB)
    '{"Name": "Jane Smith"}', -- receiver (JSONB)
    'COMPLETED', -- status
    NOW (), -- created_at
    NOW () -- updated_at
  );

-- Seed the `transaction_history` table
INSERT INTO
  transaction_history (
    user_id,
    account_id,
    transaction_type,
    transaction_amount,
    ending_balance,
    status,
    created_at,
    updated_at
  )
VALUES
  (
    (
      SELECT
        id
      FROM
        "user"
      WHERE
        username = 'john_doe'
    ),
    (
      SELECT
        id
      FROM
        account
      WHERE
        user_id = (
          SELECT
            id
          FROM
            "user"
          WHERE
            username = 'john_doe'
          LIMIT
            1
        )
        AND pocket_number = 1
    ),
    'OUTGOING',
    -1000000,
    1100000,
    'SUCCESS',
    NOW (),
    NOW ()
  ),
  (
    (
      SELECT
        id
      FROM
        "user"
      WHERE
        username = 'jane_smith'
    ),
    (
      SELECT
        id
      FROM
        account
      WHERE
        user_id = (
          SELECT
            id
          FROM
            "user"
          WHERE
            username = 'jane_smith'
          LIMIT
            1
        )
        AND pocket_number = 1
    ),
    'INCOMING',
    1000000,
    50000,
    'SUCCESS',
    NOW (),
    NOW ()
  );