CREATE TABLE person (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255)
);

INSERT INTO person (id, name) VALUES (1, 'Ava');
INSERT INTO person (id, name) VALUES (2, 'John');
INSERT INTO person (id, name) VALUES (3, 'June');
INSERT INTO person (id, name) VALUES (4, 'Tim');

SELECT * FROM person;
UPDATE person SET name='Jack' WHERE id='2';
DELETE FROM person WHERE id='2';