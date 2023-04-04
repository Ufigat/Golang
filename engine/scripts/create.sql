CREATE TABLE IF NOT EXISTS public.engines (
    id  SERIAL PRIMARY KEY,
    designation varchar(200)
);

INSERT INTO public.engines(designation)
	VALUES ('V1'),('V2'),('V3'),('V4');