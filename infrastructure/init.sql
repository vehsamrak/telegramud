CREATE DATABASE telegramud;
CREATE USER telegramud WITH password 'telegramud';
GRANT ALL ON DATABASE telegramud TO telegramud;

-- User --
CREATE TABLE public.users
(
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(25),
  race VARCHAR(25) DEFAULT NULL
);
CREATE UNIQUE INDEX users_name_uindex ON public.users (name);
