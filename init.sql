CREATE TABLE public."quotes"
(
    id serial NOT NULL,
    name character varying(45) NOT NULL,
    service character varying(45) NOT NULL,
    deadline integer NOT NULL,
    price numeric(12,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."quotes"
    OWNER to admin;