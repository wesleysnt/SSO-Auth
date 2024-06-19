CREATE TABLE access_tokens(
    id SERIAL PRIMARY KEY,
    token TEXT,
    user_id BIGINT NOT NULL,
    client_id BIGINT NOT NULL,
    expiry_time TIMESTAMP NOT NULL,
    scope TEXT NULL,
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