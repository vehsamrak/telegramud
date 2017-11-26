CREATE USER telegramud WITH password '5d768075a1aaa8a64c02c1d4afe8e252';
CREATE DATABASE telegramud OWNER telegramud;

-- changing active database
\connect telegramud telegramud
