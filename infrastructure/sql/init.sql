CREATE USER telegramud WITH password 'telegramud';
CREATE DATABASE telegramud OWNER telegramud;

-- changing active database
\connect telegramud telegramud

-- User --
CREATE TABLE public.users
(
  name VARCHAR(25) PRIMARY KEY,
  race VARCHAR(25) DEFAULT NULL
);

\connect telegramud postgres

DROP DATABASE IF EXISTS telegramud_test;
CREATE DATABASE telegramud_test WITH TEMPLATE telegramud OWNER telegramud;
