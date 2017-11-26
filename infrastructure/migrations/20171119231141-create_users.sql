-- +migrate Up
CREATE TABLE users
(
  name text NOT NULL
    CONSTRAINT users_pkey
    PRIMARY KEY,
  race text,
  room varchar(100)
);

-- +migrate Down
DROP TABLE public.users;
