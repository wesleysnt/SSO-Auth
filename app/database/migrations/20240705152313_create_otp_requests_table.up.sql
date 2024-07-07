CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE otp_requests (
  id SERIAL PRIMARY KEY NOT NULL,
  otp_code varchar(10) not null,
  unique_code uuid not null DEFAULT uuid_generate_v4(),
  is_used boolean not null default false,
  user_id bigint NOT NULL,
  expired_at timestamp not null,
  created_at timestamp NOT NULL,
  updated_at timestamp NULL,
  deleted_at timestamp NULL,

  CONSTRAINT fk_user
    FOREIGN KEY(user_id)
      REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

create index idx__otp_request_unique_code on otp_requests(unique_code);
create index idx__otp_request_is_used on otp_requests(is_used);