CREATE ROLE admin;
CREATE ROLE client;

GRANT INSERT,SELECT,UPDATE ON configuration,namespace,application TO admin;
GRANT SELECT ON configuration,application,namespace TO client;
