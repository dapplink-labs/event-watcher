DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
        CREATE DOMAIN UINT256 AS NUMERIC
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
    ELSE
        ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
        ALTER DOMAIN UINT256 ADD
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
    END IF;
END $$;


CREATE TABLE IF NOT EXISTS block_headers (
    hash        VARCHAR PRIMARY KEY,
    parent_hash VARCHAR NOT NULL UNIQUE,
    number      UINT256 NOT NULL UNIQUE,
    timestamp   INTEGER NOT NULL UNIQUE CHECK (timestamp > 0),
    rlp_bytes   VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS block_headers_timestamp ON block_headers(timestamp);
CREATE INDEX IF NOT EXISTS block_headers_number ON block_headers(number);


CREATE TABLE IF NOT EXISTS contract_events (
    guid             VARCHAR PRIMARY KEY,
    block_hash       VARCHAR NOT NULL REFERENCES block_headers(hash) ON DELETE CASCADE,
    contract_address VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    log_index        INTEGER NOT NULL,
    event_signature  VARCHAR NOT NULL,
    timestamp        INTEGER NOT NULL CHECK (timestamp > 0),
    rlp_bytes        VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS contract_events_timestamp ON contract_events(timestamp);
CREATE INDEX IF NOT EXISTS contract_events_block_hash ON contract_events(block_hash);
CREATE INDEX IF NOT EXISTS contract_events_event_signature ON contract_events(event_signature);
CREATE INDEX IF NOT EXISTS contract_events_contract_address ON contract_events(contract_address);


CREATE TABLE IF NOT EXISTS deposit_tokens (
    guid                          VARCHAR PRIMARY KEY,
    block_number                  UINT256 NOT NULL,
    token_address                 VARCHAR NOT NULL,
    sender                        VARCHAR NOT NULL,
    amount                        UINT256,
    timestamp                     INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS deposit_tokens_sender ON deposit_tokens(sender);
CREATE INDEX IF NOT EXISTS deposit_tokens_timestamp ON deposit_tokens(timestamp);


CREATE TABLE IF NOT EXISTS withdraw_tokens (
    guid                          VARCHAR PRIMARY KEY,
    block_number                  UINT256 NOT NULL,
    token_address                 VARCHAR NOT NULL,
    sender                        VARCHAR NOT NULL,
    receiver                      VARCHAR NOT NULL,
    amount                        UINT256,
    timestamp                     INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS withdraw_tokens_token_address ON withdraw_tokens(token_address);
CREATE INDEX IF NOT EXISTS withdraw_tokens_timestamp ON withdraw_tokens(timestamp);


CREATE TABLE IF NOT EXISTS grant_reward_tokens (
    guid                          VARCHAR PRIMARY KEY,
    block_number                  UINT256 NOT NULL,
    token_address                 VARCHAR NOT NULL,
    granter                       VARCHAR NOT NULL,
    amount                        UINT256,
    timestamp                     INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS grant_reward_tokens_token_address ON grant_reward_tokens(token_address);
CREATE INDEX IF NOT EXISTS grant_reward_tokens_timestamp ON grant_reward_tokens(timestamp);


CREATE TABLE IF NOT EXISTS withdraw_manager_update (
    guid                          VARCHAR PRIMARY KEY,
    block_number                  UINT256 NOT NULL,
    withdraw_manager              VARCHAR NOT NULL,
    timestamp                     INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS withdraw_manager_update_withdraw_manager ON withdraw_manager_update(withdraw_manager);
CREATE INDEX IF NOT EXISTS withdraw_manager_update_timestamp ON withdraw_manager_update(timestamp);
