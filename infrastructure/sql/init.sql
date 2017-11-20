CREATE USER telegramud WITH password 'telegramud';
CREATE DATABASE telegramud OWNER telegramud;

-- changing active database
\connect telegramud telegramud
