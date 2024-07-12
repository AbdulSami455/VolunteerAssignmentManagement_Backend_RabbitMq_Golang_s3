CREATE KEYSPACE IF NOT EXISTS managers_keyspace
WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': 1
};

CREATE KEYSPACE IF NOT EXISTS volunteers_keyspace
WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': 1
};

CREATE TABLE IF NOT EXISTS managers_keyspace.users (
  username text PRIMARY KEY,
  password text
);

CREATE TABLE IF NOT EXISTS volunteers_keyspace.users (
  username text PRIMARY KEY,
  password text
);
INSERT INTO managers_keyspace.users (username, password) VALUES ('manager1', 'hashed_password_here');
INSERT INTO volunteers_keyspace.users (username, password) VALUES ('volunteer1', 'hashed_password_here');
