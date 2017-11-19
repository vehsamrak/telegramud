-- +migrate Up
create table rooms
(
  id varchar(100) not null
    constraint rooms_pkey
    primary key,
  name varchar(100),
  description text,
  coordinates text
)
;

create unique index uix_rooms_coordinates
  on rooms (coordinates)
;

-- +migrate Down
