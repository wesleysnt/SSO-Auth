CREATE TABLE auth_codes(
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    expiry_time TIMESTAMP NOT NULL,
    client_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL,


    CONSTRAINT fk_client_id
    FOREIGN KEY(client_id)
        REFERENCES clients(id) ON UPDATE CASCADE ON DELETE SET NULL,

    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
        REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
)