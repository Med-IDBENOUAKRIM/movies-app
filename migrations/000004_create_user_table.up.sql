CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  name text NOT NULL,
  email citext UNIQUE NOT NULL,
  password_hash bytea UNIQUE NOT NULL,
  activated bool UNIQUE NOT NULL,
  version integer UNIQUE NOT NULL DEFAULT 1
);