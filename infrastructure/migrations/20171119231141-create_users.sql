-- +migrate Up
create table users
(
  name text not null
    constraint users_pkey
    primary key,
  race text,
  room varchar(100)
)
;

-- +migrate Down
