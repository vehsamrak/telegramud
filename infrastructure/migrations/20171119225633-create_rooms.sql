-- +migrate Up
CREATE TABLE rooms
(
  id          varchar(100) NOT NULL
    CONSTRAINT rooms_pkey
    PRIMARY KEY,
  name        varchar(100),
  description text,
  coordinates text
);

CREATE UNIQUE INDEX uix_rooms_coordinates
  ON rooms (coordinates);

-- +migrate Down
DROP TABLE public.rooms;
