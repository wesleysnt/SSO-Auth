CREATE TABLE clients (
  id SERIAL PRIMARY KEY,
  client_id varchar(50) NOT NULL,
  secret varchar(255) NOT NULL,
  redirect_uri TEXT NOT NULL, 
  created_at timestamp NULL,
  updated_at timestamp NULL,
  deleted_at timestamp NULL
);