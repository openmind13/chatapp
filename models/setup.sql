

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    username VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(255) UNIQUE,
    encrypted_password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- CREATE TABLE user_sessions (
--     id SERIAL PRIMARY KEY,
--     uuid VARCHAR(64) NOT NULL UNIQUE,
--     user_uuid INTEGER REFERENCES users(uuid),
--     created_at TIMESTAMP NOT NULL
-- );