CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE clients (
  id SERIAL PRIMARY KEY,
  client_id uuid NULL default uuid_generate_v4(),
  name varchar(255) not null,
  secret varchar(255) NOT NULL,
  redirect_uri TEXT NOT NULL, 
  created_at timestamp NULL,
  updated_at timestamp NULL,
  deleted_at timestamp NULL
);