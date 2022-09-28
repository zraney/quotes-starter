DROP TABLE IF EXISTS quotes;

CREATE TABLE quotes (id varchar(50) PRIMARY KEY, phrase varchar(200), author varchar(50));

INSERT INTO quotes (id, phrase, author) 
VALUES 
('b37c9ded-d176-4fe5-a9b9-1427ebf92ed1', 'Errors are values.', 'Rob Pike'),
('0d95d2d8-28b0-4278-960d-cbdd16beab02', 'Clear is better than clever.', 'Rob Pike'),
('0329b963-004d-4add-bb5e-cfe7defd0c6d', 'Don''t panic.', 'Go Code Review Comments'),
('2e774b8c-672e-46bf-8b6f-38d6889edee7', 'A little copying is better than a little dependency.', 'Rob Pike'),
('a2ad7811-22ea-4ba4-8691-a88b5f89a475', 'Concurrency is not parallelism.', 'Rob Pike'),
('ba9b4b54-3070-4665-bd29-de3e99c991d2', 'interface{} says nothing.', 'Rob Pike'),
('ca17bd05-4c0b-41ae-9496-518371e245f2', 'Make the zero value useful.', 'Rob Pike');
