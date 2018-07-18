CREATE TABLE history(
	id serial PRIMARY KEY,
	user_id integer,
	namespace_id integer,
	predecessor_version integer,
	successor_version integer,
	created_at timestamp
);
