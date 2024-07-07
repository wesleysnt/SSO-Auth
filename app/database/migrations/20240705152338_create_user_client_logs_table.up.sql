CREATE TABLE user_client_logs (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id integer NOT NULL,
  client_id integer NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp NULL,

  CONSTRAINT fk_client_id
    FOREIGN KEY(client_id)
        REFERENCES clients(id) ON UPDATE CASCADE ON DELETE SET NULL,

  CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
        REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);
