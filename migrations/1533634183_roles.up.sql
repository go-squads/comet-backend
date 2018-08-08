CREATE ROLE admin;
CREATE ROLE client;

GRANT INSERT,SELECT,UPDATE ON configuration,namespace,application TO admin;
GRANT USAGE, SELECT ON SEQUENCE application_id_seq,history_id_seq, namespace_id_seq, users_id_seq TO admin;

GRANT SELECT ON configuration,application,namespace,users TO client;
GRANT SELECT ON users TO admin;