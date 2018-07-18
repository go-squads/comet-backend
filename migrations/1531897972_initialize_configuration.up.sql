CREATE TABLE configuration(
  namespace_id integer FOREIGN KEY REFERENCES namespace (id),
  version integer,
  key varchar(255),
  value varchar(255),
  PRIMARY KEY (key, version, namespace_id)
);
