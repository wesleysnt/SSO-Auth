CREATE TABLE forgot_password_tokens (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id bigint NOT NULL,
  token varchar(255) NOT NULL, 
  expiry_time timestamp NOT NULL,
  is_used boolean NULL default false,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp NULL
);
