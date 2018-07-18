CREATE TABLE configuration(
  key varchar(255),
  version integer,
  namespace_id varchar(255) FOREIGN KEY REFERENCES namespace (id),
  value varchar(255),
  PRIMARY KEY (key, version, namespace_id)
);
