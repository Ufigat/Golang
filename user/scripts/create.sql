CREATE TABLE IF NOT EXISTS public.users (
    id  SERIAL PRIMARY KEY,
    name varchar(200)
);

CREATE TABLE IF NOT EXISTS public.user_cars (
    user_id int,
    car_id int,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx1_user_cars
ON public.user_cars (user_id);

INSERT INTO public.users(name) VALUES ('First'), ('Second');

INSERT INTO public.user_cars(
    user_id, car_id)
	VALUES (1,2), (1,3), (2,1), (2,2), (1,4), (2,3), (2,4);