CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  username varchar(50) NOT NULL,
  created_at timestamp NULL,
  updated_at timestamp NULL,
  deleted_at timestamp NULL

);
create index idx_users_email on users(email);
create index idx_users_username on users(username);