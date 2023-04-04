CREATE TABLE IF NOT EXISTS public.users (
    id  SERIAL PRIMARY KEY,
    name varchar(200)
);

INSERT INTO public.users(name) VALUES ('First'), ('Second');