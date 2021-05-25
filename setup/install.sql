
-- DROP DATABASE testing_api;

CREATE DATABASE testing_api
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
    
DROP TABLE IF EXISTS public.expiring_courses;
DROP TABLE IF EXISTS public.course;

CREATE TABLE public.course
(
    id integer NOT NULL,
    name character varying(100) COLLATE pg_catalog."default",
    CONSTRAINT course_pkey PRIMARY KEY (id)
);

CREATE TABLE public.expiring_courses
(
    id integer NOT NULL,
    expiring timestamp with time zone NOT NULL,
    link character(32)[] COLLATE pg_catalog."default" NOT NULL,
    empl_name character varying(50) COLLATE pg_catalog."default",
    empl_dep character varying(50) COLLATE pg_catalog."default",
    passing smallint,
    course_id integer,
    CONSTRAINT expiring_courses_pkey PRIMARY KEY (id),
    CONSTRAINT course_fk FOREIGN KEY (course_id)
        REFERENCES public.course (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);
