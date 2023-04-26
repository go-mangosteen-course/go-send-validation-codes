CREATE TABLE validation_codes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    used_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

