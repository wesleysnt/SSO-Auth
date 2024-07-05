CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  created_at timestamp NULL,
  updated_at timestamp NULL,
  deleted_at timestamp NULL

);
create index idx_users_email on users(email);