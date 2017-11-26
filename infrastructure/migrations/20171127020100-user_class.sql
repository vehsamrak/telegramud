-- +migrate Up
ALTER TABLE public.users ADD class varchar(100) DEFAULT NULL NOT NULL;

-- +migrate Down
ALTER TABLE public.users DROP class;
