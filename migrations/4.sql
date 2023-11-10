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
    userid bigint NOT NULL unique  ON DELETE CASCADE ,
    mon    boolean,
    tue    boolean,
    wed    boolean,
    thu    boolean,
    fri    boolean,
    sat    boolean,
    sun    boolean,
    t      text,
    CONSTRAINT fk_users
      FOREIGN KEY(userid) 
	  REFERENCES users(id)
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