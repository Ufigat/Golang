CREATE TABLE IF NOT EXISTS public.colors (
    id  SERIAL PRIMARY KEY,
    name varchar(200)
);

CREATE TABLE IF NOT EXISTS public.brands (
    id  SERIAL PRIMARY KEY,
    name varchar(200)
);

CREATE TABLE IF NOT EXISTS public.cars (
    id  SERIAL PRIMARY KEY,
    color_id int,
    brand_id int,
    engine_id int,
    UNIQUE (color_id, brand_id, engine_id),
    FOREIGN KEY (color_id) REFERENCES colors(id),
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);

CREATE INDEX IF NOT EXISTS idx_cars_brands
ON public.cars (brand_id);

CREATE INDEX IF NOT EXISTS idx_cars_colors
ON public.cars (color_id);

CREATE TABLE IF NOT EXISTS public.user_cars (
    user_id int,
    car_id int,
    FOREIGN KEY (car_id) REFERENCES cars(id)
);

CREATE INDEX IF NOT EXISTS idx1_user_cars
ON public.user_cars (car_id);


INSERT INTO public.brands(name) VALUES ('BMW'), ('Audi');

INSERT INTO public.colors(name) VALUES ('Blue'), ('Red');

INSERT INTO public.cars (color_id, brand_id, engine_id)
	VALUES (1,1,1), (1,1,2),(1,2,1),(1,2,2),(2,1,1),(2,1,2),(2,2,1),(2,2,2);

INSERT INTO public.user_cars(
    user_id, car_id)
	VALUES (1,2), (1,3), (2,1), (2,2), (1,4), (2,3), (2,4);