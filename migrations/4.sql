CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL DEFAULT ,
    email text NOT NULL,
    password text NOT NULL,
    name text  NOT NULL,
    gender int,
    height int,
    weight int,
    birth timestamp ,
    activity int,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
);