CREATE TABLE configuration_change(
	history_id integer REFERENCES history (id),
	key varchar(255),
	old_value varchar(255),
	new_value varchar(255),
	PRIMARY KEY (history_id, key)
);
