CREATE TABLE user(
	id int PRIMARY KEY AUTO_INCREMENT, 
  first_name VARCHAR(50) NOT NULL DEFAULT '',
  last_name VARCHAR(50) NOT NULL DEFAULT '',
  user_level INT DEFAULT 1, 
  exp int DEFAULT 0,
  age int NOT NULL DEFAULT 0,
  gender text NOT NULL,
  email text NOT NULL,
  updated_at timestamp,
  created_at timestamp 
);

CREATE TABLE task (
  id int PRIMARY KEY AUTO_INCREMENT,
  title  text,
   amount  int,
   amount_unit  text,
   duration  text,
   user_id  int
);

ALTER TABLE  task  ADD FOREIGN KEY ( user_id ) REFERENCES  user  ( id );