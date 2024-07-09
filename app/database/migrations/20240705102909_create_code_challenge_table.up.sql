CREATE TABLE code_challenges (
  id SERIAL PRIMARY KEY NOT NULL,
  code varchar(255) not null,
  method varchar(10) not null,
  client_id bigint not null,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp NULL,


  CONSTRAINT fk_client_id
    FOREIGN KEY(client_id)
        REFERENCES clients(id) ON UPDATE CASCADE ON DELETE SET NULL
);
