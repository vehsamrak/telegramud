CREATE USER telegramud WITH password 'telegramud';
CREATE DATABASE telegramud OWNER telegramud;

-- changing active database
\connect telegramud telegramud

-- process initial SQL here

-- creating test database
\connect telegramud postgres
DROP DATABASE IF EXISTS telegramud_test;
CREATE DATABASE telegramud_test WITH TEMPLATE telegramud OWNER telegramud;
