CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL,
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

-- ON DELETE CASCADE

CREATE TABLE IF NOT EXISTS public.days
(
    id     SERIAL primary key,
    userid bigint BIGINT NOT NULL unique REFERENCES users(id) ON DELETE CASCADE,
    mon    boolean,
    tue    boolean,
    wed    boolean,
    thu    boolean,
    fri    boolean,
    sat    boolean,
    sun    boolean,
    t      text
    -- CONSTRAINT fk_users
    --   FOREIGN KEY(userid) 
	--   REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS public.products
(
    id          SERIAL primary key,
	Name        text,
	Title       text,
	Description text,
	ImageSrc    text,
	VideoSrc    text,
	Price       real,
	Weight      real,
	Length      text,
	Height      text,
	Depth       text,
	Icon        text,
	Status      int
);

CREATE TABLE IF NOT EXISTS public.lessons
(
    id          SERIAL primary key,
	Name        text,
	Title       text,
	Description text,
	ImageSrc    text,
	VideoSrc    text,
	Duration    real,
    Created_at  timestamp,
    Updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS public.favorites
(
    id       SERIAL PRIMARY KEY,
    userid   BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lessonid BIGINT NOT NULL UNIQUE REFERENCES lessons(id) ON DELETE CASCADE
);