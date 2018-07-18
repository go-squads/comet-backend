CREATE TABLE namespace(
  id serial PRIMARY KEY,
  name varchar(255),
  app_id FOREIGN KEY REFERENCES application (id),
  active_version serial
);
