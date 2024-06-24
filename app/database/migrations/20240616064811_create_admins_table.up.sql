CREATE TABLE admins(
    id SERIAL PRIMARY KEY,
    email varchar(255) NOT NULL, 
    password varchar(255) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL
)
