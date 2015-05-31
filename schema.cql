// Holds every iteration of all the schemas messed with so far
CREATE KEYSPACE alexandria WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor' : 1 };

use alexandria;

CREATE TABLE origins (name text, originid uuid, primary key(name));
CREATE INDEX originsIdIndex ON origins (originid);


CREATE TABLE rr (id uuid, origin uuid, type text, name text, value text, ttl int, primary key(id));
CREATE TABLE originrr (id uuid, rr uuid, primary key(id, rr));
CREATE TABLE qanswer (question text, answer text, primary key(question, answer));
