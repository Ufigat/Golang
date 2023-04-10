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
    color_id int NOT NULL,
    brand_id int NOT NULL,
    engine_id int NOT NULL,
    UNIQUE (color_id, brand_id, engine_id),
    FOREIGN KEY (color_id) REFERENCES colors(id),
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);

CREATE INDEX IF NOT EXISTS idx_cars_brands
ON public.cars (brand_id);

CREATE INDEX IF NOT EXISTS idx_cars_colors
ON public.cars (color_id);

INSERT INTO public.brands(name) VALUES ('BMW'), ('Audi');

INSERT INTO public.colors(name) VALUES ('Blue'), ('Red');

INSERT INTO public.cars (color_id, brand_id, engine_id)
	VALUES (1,1,1), (1,1,2),(1,2,1),(1,2,2),(2,1,1),(2,1,2),(2,2,1),(2,2,2);
