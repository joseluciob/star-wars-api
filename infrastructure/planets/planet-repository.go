CREATE TABLE planets (
	id SERIAL PRIMARY KEY,
	name VARCHAR(45) NULL,
	climate VARCHAR(45) NULL,
	terrain VARCHAR(45) NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY(id)
);

CREATE TABLE films (
	id SERIAL PRIMARY KEY,
	title VARCHAR(45) NULL,
	director VARCHAR(45) NULL,
	release_date VARCHAR(45) NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY(id)
 );

CREATE TABLE planets_films (
	film_id    int REFERENCES films (id) ON UPDATE CASCADE ON DELETE CASCADE,
	planet_id int REFERENCES planets (id) ON UPDATE CASCADE,
	CONSTRAINT planets_films_key PRIMARY KEY (film_id, planet_id)
);
