CREATE TABLE if not exist users (
    created_at       TIMESTAMP NOT NULL DEFAULT now(),
    updated_at       TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at       TIMESTAMP,
    id               uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    credit_card_token VARCHAR(255) UNIQUE NOT NULL,
    user_document    VARCHAR(255) UNIQUE NOT NULL,
    value             VARCHAR(255) NOT NULL
);