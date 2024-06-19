CREATE TABLE refresh_tokens(
    id SERIAL PRIMARY KEY,
    token TEXT, 
    user_id BIGINT,
    client_id BIGINT,
    expiry_time TIMESTAMP,

    CONSTRAINT fk_client_id
    FOREIGN KEY(client_id)
        REFERENCES clients(id) ON UPDATE CASCADE ON DELETE SET NULL,

    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
        REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
)
