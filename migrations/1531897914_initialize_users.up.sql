CREATE TABLE users(
	id serial PRIMARY KEY,
	username varchar(255),
	password varchar(255),
	role varchar(255),
	token varchar(255)
);
